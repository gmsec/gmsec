// Command gmsecgen 是 gmsec 仓库的脚手架/清理工具。
//
// 它把根 Makefile 里两段平台相关的 shell 逻辑收进一处:
//
//  1. `make new service=foo` 原来的实现是 cp -r + 一堆 `sed -i 's#example#foo#g'`,
//     Windows 分支用 `sed -i`,macOS 分支用 `sed -i ”`,Linux 用 GNU sed;
//     漏掉一个文件就会留下半改不改的服务,而 sed 对二进制文件还会损坏内容。
//     这里改成一次遍历目录树:文本文件按内容判断二进制,统一替换;
//  2. `make clear` 原来是 `rm -rf ./prc/**/*.go`——目录名拼错了(应为 rpc),
//     且 bash 默认不展开 `**`,实际什么都不会删。这里按后缀精确删除生成物;
//  3. `proto` / `clean-proto` 取代已不存在的 example/tools/protogen:
//     调用系统 protoc 生成 pb.go 并同步到两个 rpc 目录,工具链缺失时明确报错。
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// serviceName 限制成小写字母开头的标识符:它会被直接拼进 go module 路径、
// grpc service 名和目录名,允许 `-`、`/` 会生成无法编译的代码。
var serviceName = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

func main() {
	mode := "scaffold"
	args := os.Args[1:]
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		mode, args = args[0], args[1:]
	}

	var err error
	switch mode {
	case "scaffold", "new":
		err = runScaffold(args)
	case "clean":
		err = runClean(args)
	case "proto":
		err = runProto(args)
	case "clean-proto":
		err = runCleanProto(args)
	default:
		err = fmt.Errorf("未知子命令 %q,可用:scaffold / clean / proto / clean-proto", mode)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "gmsecgen: %v\n", err)
		os.Exit(1)
	}
}

func runScaffold(argv []string) error {
	fs := flag.NewFlagSet("scaffold", flag.ExitOnError)
	root := fs.String("root", "..", "仓库根目录")
	tmplName := fs.String("template", "example", "模板服务名")
	name := fs.String("service", "", "新服务名(必填)")
	dryRun := fs.Bool("dry-run", false, "只打印将要做的改动")
	_ = fs.Parse(argv)

	if *name == "" {
		return errors.New("缺少 -service")
	}
	if !serviceName.MatchString(*name) {
		return fmt.Errorf("服务名 %q 非法:只允许小写字母开头的小写字母/数字/下划线", *name)
	}
	if *name == *tmplName {
		return fmt.Errorf("服务名不能和模板相同(%s)", *tmplName)
	}

	rootDir, err := filepath.Abs(*root)
	if err != nil {
		return err
	}
	srcSvc := filepath.Join(rootDir, *tmplName)
	dstSvc := filepath.Join(rootDir, *name)
	if _, err := os.Stat(srcSvc); err != nil {
		return fmt.Errorf("模板目录不存在: %s", srcSvc)
	}
	if _, err := os.Stat(dstSvc); err == nil {
		// 历史实现是 `test -d ... || cp -r`,存在就跳过复制但仍然 sed,
		// 会把已经改好的服务再改一遍。这里直接拒绝,避免误伤。
		return fmt.Errorf("服务目录已存在: %s(如需重建请先自行删除)", dstSvc)
	}

	fmt.Printf("gmsecgen: %s -> %s\n", *tmplName, *name)
	if err := copyTree(srcSvc, dstSvc, *tmplName, *name, *dryRun); err != nil {
		return err
	}

	srcProto := filepath.Join(rootDir, "apidoc", "proto", *tmplName)
	dstProto := filepath.Join(rootDir, "apidoc", "proto", *name)
	if fi, err := os.Stat(srcProto); err == nil && fi.IsDir() {
		if err := copyTree(srcProto, dstProto, *tmplName, *name, *dryRun); err != nil {
			return err
		}
	}
	fmt.Printf("gmsecgen: 完成。接下来执行: cd %s && make proto\n", *name)
	return nil
}

// copyTree 递归复制 src 到 dst,并把文本文件里的 oldName 替换成 newName。
func copyTree(src, dst, oldName, newName string, dryRun bool) error {
	err := filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if rel == "." {
			if dryRun {
				return nil
			}
			return os.MkdirAll(dst, 0755)
		}
		// 版本库、依赖目录和 IDE 日志都不该被复制。
		if fi.IsDir() {
			if skipDir(fi.Name()) {
				return filepath.SkipDir
			}
			if dryRun {
				return nil
			}
			return os.MkdirAll(filepath.Join(dst, rel), 0755)
		}
		if skipFile(fi.Name()) {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		out := data
		if !isBinary(data) {
			// 服务名会出现在 module 名、import 路径(example/internal/...)、
			// grpc 服务名(gmsec.srv.example)、配置和注释里,所以按全局替换处理,
			// 与旧的 sed 's#example#new#g' 行为保持一致。
			//
			// 额外要处理首字母大写的形式:protoc 会把 `service example` 的 Go 名字
			// 转成 `Example`,生成代码里就是 RegisterExampleServer / GetExampleName。
			// 旧实现只 sed 小写形式,新建服务编译会直接报 undefined。
			for _, pair := range replacements(oldName, newName) {
				out = bytes.ReplaceAll(out, []byte(pair[0]), []byte(pair[1]))
			}
		}
		target := filepath.Join(dst, rel)
		if dryRun {
			fmt.Println("  +", target)
			return nil
		}
		mode := fi.Mode().Perm()
		if mode&0111 != 0 {
			mode = 0755 // 脚本保留可执行位
		} else {
			mode = 0644
		}
		return os.WriteFile(target, out, mode)
	})
	if err != nil {
		return err
	}
	if dryRun {
		return nil
	}
	// 跳过生成物后可能留下空目录(rpc/<old> 这种),清一下,免得多出一个空壳。
	return pruneEmptyDirs(dst)
}

// pruneEmptyDirs 自底向上删除空目录(dst 本身保留)。
func pruneEmptyDirs(root string) error {
	var dirs []string
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil || !fi.IsDir() {
			return err
		}
		dirs = append(dirs, path)
		return nil
	})
	if err != nil {
		return err
	}
	// Walk 是先父后子,倒着走才能把父目录也腾出来。
	for i := len(dirs) - 1; i >= 0; i-- {
		if dirs[i] == root {
			continue
		}
		entries, err := os.ReadDir(dirs[i])
		if err != nil || len(entries) > 0 {
			continue
		}
		_ = os.Remove(dirs[i])
	}
	return nil
}

// replacements 返回需要替换的字面量对:首字母大写的 CamelCase、原样、全大写。
// 顺序上先换 CamelCase,避免子串互相干扰。
func replacements(oldName, newName string) [][2]string {
	return [][2]string{
		{goCamel(oldName), goCamel(newName)},
		{strings.ToUpper(oldName), strings.ToUpper(newName)},
		{oldName, newName},
	}
}

// goCamel 按 protoc 的规则把 service 名转成 Go 名字:
// example -> Example,user_service -> UserService。
func goCamel(name string) string {
	parts := strings.Split(name, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		b.WriteString(p[1:])
	}
	return b.String()
}

func skipDir(name string) bool {
	switch name {
	// .git/依赖/IDE 配置/工具元数据不是服务代码,复制过去只会变成噪音。
	case ".git", "node_modules", "vendor", "__pycache__", ".vscode", ".workbuddy", ".idea":
		return true
	}
	return false
}

func skipFile(name string) bool {
	// 生成物不复制:新服务的 go_package 已经改成 rpc/<新服务>,
	// 旧目录里的 rpc/<模板> 只会变成没人引用的死代码。
	if isGenerated(name) {
		return true
	}
	switch filepath.Ext(name) {
	case ".log", ".exe", ".test", ".zip", ".png", ".jpg", ".jpeg", ".gif", ".ico":
		return true
	}
	return false
}

// isBinary 用 NUL 字节判断:文本文件不会有 NUL,ELF/PE 前几字节就有。
func isBinary(data []byte) bool {
	n := len(data)
	if n > 8000 {
		n = 8000
	}
	return bytes.IndexByte(data[:n], 0) >= 0
}

var (
	goPackageRe = regexp.MustCompile(`(?m)^\s*option\s+go_package\s*=\s*"([^"]+)"`)
	packageRe   = regexp.MustCompile(`(?m)^package\s+(\w+)`)
)

// goPackageSubdir 读 proto 的 go_package 推导生成物落到 rpc 下的子目录:
// 例如 "rpc/example" -> "example";读不到时回退到服务名。
func goPackageSubdir(protoFiles []string, fallback string) (string, error) {
	for _, f := range protoFiles {
		data, err := os.ReadFile(f)
		if err != nil {
			return "", err
		}
		m := goPackageRe.FindSubmatch(data)
		if m == nil {
			continue
		}
		pkg := string(m[1])
		if i := strings.LastIndex(pkg, ";"); i >= 0 { // 兼容 "path;pkgname" 形式
			pkg = pkg[i+1:]
		}
		pkg = strings.Trim(strings.TrimPrefix(pkg, "."), "/")
		if pkg == "" {
			continue
		}
		parts := strings.Split(pkg, "/")
		return parts[len(parts)-1], nil
	}
	return fallback, nil
}

// verifyExistingPackage 读现有 pb.go 的 package 声明,确认目标目录约定一致,
// 避免把生成物同步进错误的目录。目录不存在(首次生成)时直接放行。
func verifyExistingPackage(dir, want string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".pb.go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return err
		}
		m := packageRe.FindSubmatch(data)
		if m == nil {
			return fmt.Errorf("现有生成文件缺少 package 声明: %s", e.Name())
		}
		if string(m[1]) != want {
			return fmt.Errorf("现有 pb.go 的 package 为 %q,与 go_package 推导出的 %q 不一致,拒绝写入 %s", m[1], want, dir)
		}
		return nil
	}
	return nil
}

// isGeneratedCode 用文件头判断是不是真正的生成器产物,防止误删同名手写代码。
func isGeneratedCode(data []byte) bool {
	head := data
	if len(head) > 512 {
		head = head[:512]
	}
	return bytes.Contains(head, []byte("Code generated by"))
}

// writeGenerated 只覆盖已经是生成物的文件;若目标存在但不像生成物则拒绝覆盖。
func writeGenerated(target string, data []byte) error {
	if old, err := os.ReadFile(target); err == nil {
		if !isGeneratedCode(old) {
			return fmt.Errorf("拒绝覆盖疑似手写文件: %s", target)
		}
	}
	return os.WriteFile(target, data, 0644)
}

// runProto 用系统 protoc 生成 pb.go 并同步到 apidoc/rpc/<name> 与 example/rpc/<name>。
// 工具链缺失或 protoc 失败时都返回非 0,绝不静默成功。
func runProto(argv []string) error {
	fs := flag.NewFlagSet("proto", flag.ExitOnError)
	root := fs.String("root", "..", "仓库根目录")
	name := fs.String("name", "", "服务名(必填),对应 apidoc/proto/<name>")
	_ = fs.Parse(argv)

	if *name == "" {
		return errors.New("缺少 -name")
	}
	if !serviceName.MatchString(*name) {
		return fmt.Errorf("服务名 %q 非法:只允许小写字母开头的小写字母/数字/下划线", *name)
	}

	rootDir, err := filepath.Abs(*root)
	if err != nil {
		return err
	}
	protoRoot := filepath.Join(rootDir, "apidoc", "proto")
	protoDir := filepath.Join(protoRoot, *name)
	if fi, err := os.Stat(protoDir); err != nil || !fi.IsDir() {
		return fmt.Errorf("proto 目录不存在: %s", protoDir)
	}
	protoFiles, err := filepath.Glob(filepath.Join(protoDir, "*.proto"))
	if err != nil {
		return err
	}
	if len(protoFiles) == 0 {
		return fmt.Errorf("未找到任何 .proto 文件: %s", protoDir)
	}
	sort.Strings(protoFiles)

	// 工具链检查:protoc-gen-gmsec 内部封装了 protoc-gen-go 的消息生成,
	// 但仍要求三个工具都在 PATH 中,保证工具链完整、版本可追溯。
	var missing []string
	for _, tool := range []string{"protoc", "protoc-gen-go", "protoc-gen-gmsec"} {
		if _, err := exec.LookPath(tool); err != nil {
			missing = append(missing, tool)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("缺少 protoc 工具链: %s\n请先执行 make install(或按脚本里的固定版本 go install)把它们装到 PATH 后重试", strings.Join(missing, ", "))
	}

	// 目标目录由 go_package + 现有 pb.go 的 package 约定共同决定。
	relDir, err := goPackageSubdir(protoFiles, *name)
	if err != nil {
		return err
	}
	targets := []string{
		filepath.Join(rootDir, "apidoc", "rpc", relDir),
		filepath.Join(rootDir, "example", "rpc", relDir),
	}
	for _, dst := range targets {
		if err := verifyExistingPackage(dst, relDir); err != nil {
			return err
		}
	}

	out, err := os.MkdirTemp("", "gmsecgen-proto-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(out)

	args := []string{
		"--proto_path=" + protoRoot,
		"--gmsec_out=plugins=gmsec,paths=source_relative:" + out,
	}
	for _, f := range protoFiles {
		rel, err := filepath.Rel(protoRoot, f)
		if err != nil {
			return err
		}
		args = append(args, filepath.ToSlash(rel))
	}
	cmd := exec.Command("protoc", args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("protoc 生成失败: %w", err)
	}

	// 只收集生成器产物,手写代码一律不碰。
	var generated []string
	err = filepath.Walk(out, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() || !isGenerated(fi.Name()) {
			return nil
		}
		generated = append(generated, path)
		return nil
	})
	if err != nil {
		return err
	}
	if len(generated) == 0 {
		return errors.New("protoc 未产出任何 *.pb.go 生成物(检查插件与 proto 文件)")
	}

	for _, dst := range targets {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
		for _, src := range generated {
			data, err := os.ReadFile(src)
			if err != nil {
				return err
			}
			target := filepath.Join(dst, filepath.Base(src))
			if err := writeGenerated(target, data); err != nil {
				return err
			}
			fmt.Println("gmsecgen: 写入", target)
		}
	}
	return nil
}

// runCleanProto 只删除 *.pb.go / *.pb.gmsec.go / *.pb.constants.go 生成物。
func runCleanProto(argv []string) error {
	fs := flag.NewFlagSet("clean-proto", flag.ExitOnError)
	root := fs.String("root", "..", "仓库根目录")
	name := fs.String("name", "", "只清理某个服务(可选)")
	dryRun := fs.Bool("dry-run", false, "只打印将删除的文件")
	_ = fs.Parse(argv)

	rootDir, err := filepath.Abs(*root)
	if err != nil {
		return err
	}
	roots := []string{
		filepath.Join(rootDir, "apidoc", "rpc"),
		filepath.Join(rootDir, "example", "rpc"),
	}
	var victims []string
	for _, r := range roots {
		if _, err := os.Stat(r); err != nil {
			continue // 目录不存在就跳过
		}
		err = filepath.Walk(r, func(path string, fi os.FileInfo, err error) error {
			if err != nil || fi.IsDir() {
				return err
			}
			if !isGenerated(fi.Name()) {
				return nil
			}
			if *name != "" && filepath.Base(filepath.Dir(path)) != *name {
				return nil
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if !isGeneratedCode(data) {
				return nil // 同名手写代码,保留
			}
			victims = append(victims, path)
			return nil
		})
		if err != nil {
			return err
		}
	}
	sort.Strings(victims)
	for _, p := range victims {
		if *dryRun {
			fmt.Println("  -", p)
			continue
		}
		if err := os.Remove(p); err != nil {
			return err
		}
		fmt.Println("gmsecgen: 删除", p)
	}
	return nil
}

func runClean(argv []string) error {
	fs := flag.NewFlagSet("clean", flag.ExitOnError)
	root := fs.String("root", "..", "仓库根目录")
	dryRun := fs.Bool("dry-run", false, "只打印将删除的文件")
	_ = fs.Parse(argv)

	rootDir, err := filepath.Abs(*root)
	if err != nil {
		return err
	}
	rpcDir := filepath.Join(rootDir, "apidoc", "rpc")
	if _, err := os.Stat(rpcDir); err != nil {
		return fmt.Errorf("目录不存在: %s", rpcDir)
	}

	var victims []string
	err = filepath.Walk(rpcDir, func(path string, fi os.FileInfo, err error) error {
		if err != nil || fi.IsDir() {
			return err
		}
		if isGenerated(fi.Name()) {
			victims = append(victims, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	sort.Strings(victims)
	for _, p := range victims {
		if *dryRun {
			fmt.Println("  -", p)
			continue
		}
		if err := os.Remove(p); err != nil {
			return err
		}
		fmt.Println("gmsecgen: 删除", p)
	}
	return nil
}

// isGenerated 只认生成器会产出的后缀,避免误删手写代码。
func isGenerated(name string) bool {
	return strings.HasSuffix(name, ".pb.go") ||
		strings.HasSuffix(name, ".pb.gmsec.go") ||
		strings.HasSuffix(name, ".pb.constants.go")
}

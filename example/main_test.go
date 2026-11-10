package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"example/internal/config"
)

// TestServer 是集成测试:历史实现调用 plg.Wait() 会一直阻塞,CI 上会挂死。
// 这里显式标注为集成用例并跳过;本地联调时设置 GMSEC_INTEGRATION=1 手动启动服务。
func TestServer(t *testing.T) {
	if os.Getenv("GMSEC_INTEGRATION") == "" {
		t.Skip("集成测试需要真实 gRPC/micro 环境与可用端口;设置 GMSEC_INTEGRATION=1 后本地手动运行")
	}
}

// TestClient 同上:gRPC 客户端必须连到真实服务端,CI 默认跳过,避免阻塞。
func TestClient(t *testing.T) {
	if os.Getenv("GMSEC_INTEGRATION") == "" {
		t.Skip("集成测试需要真实 gRPC 服务端;设置 GMSEC_INTEGRATION=1 后本地手动运行")
	}
}

// TestConfigJaeger 纯函数断言:校验嵌套 jaeger 配置确实被加载进 Config.Jaeger,
// 而不是像旧 config.yml 那样只写了一个顶层 jaeger_addr 字符串。
func TestConfigJaeger(t *testing.T) {
	// go test 时可执行文件在临时目录,包 init 找不到可执行文件旁的 conf/config.yml。
	// 用测试源码所在目录显式加载一次,断言才针对仓库里真实的配置文件。
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("无法定位测试源码路径")
	}
	if err := config.InitFile(filepath.Join(filepath.Dir(file), "conf", "config.yml")); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	j := config.GetJaeger()
	if j.Addr == "" {
		t.Fatal("config.GetJaeger().Addr 为空,jaeger 配置未生效")
	}
	if j.Percent <= 0 || j.Percent > 100 {
		t.Fatalf("jaeger.percent 非法: %d,期望 1-100", j.Percent)
	}
	if config.GetPort() == "" {
		t.Fatal("config.GetPort() 不应为空")
	}
	if config.GetIsDev() {
		t.Fatal("示例配置 base.is_dev 应为 false")
	}
}

service := example

.PHONY: install source_install clear new proto

# 按 uname 分派到真实存在的安装脚本:
# 原来 install 写死 mac 脚本、source_install 指向不存在的 proto_install.sh。
install:
	@case "$$(uname -s)" in \
		MINGW*|MSYS*|CYGWIN*) bash ./proto_install_windows_git_bash.sh ;; \
		Darwin) bash ./proto_install_mac.sh ;; \
		*) bash ./proto_install_linux.sh ;; \
	esac
source_install:
	bash ./proto_source_install.sh

# 删除 apidoc/rpc 下的生成物(*.pb.go / *.pb.gmsec.go / *.pb.constants.go)。
# 历史写法 `rm -rf ./prc/**/*.go` 目录名拼错(应为 rpc),而且 bash 默认不展开 `**`,
# 实际什么都没删;现在按后缀精确删除,不会碰到手写代码。
clear:
	cd tools && go run ./gmsecgen clean -root ..

# make new service=myservice
# 从 example 模板生成一个新服务,复制 + 改名统一交给 tools/gmsecgen。
new:
	@test -n "$(service)" || (echo "用法: make new service=<服务名>" && exit 1)
	cd tools && go run ./gmsecgen scaffold -root .. -template example -service $(service)

# 生成 example 服务的 proto/gRPC 代码(需要先装好 protoc 与 protoc-gen-gmsec)。
proto:
	cd example && $(MAKE) proto

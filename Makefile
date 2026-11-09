service := example

.PHONY: install source_install clear new proto

install:
	./proto_install_mac.sh
source_install:
	./proto_install.sh

# 删除 apidoc/rpc 下的生成物(*.pb.go / *.pb.gmsec.go / *.pb.constants.go)。
# 历史写法 `rm -rf ./prc/**/*.go` 目录名拼错(应为 rpc),而且 bash 默认不展开 `**`,
# 实际什么都没删;现在按后缀精确删除,不会碰到手写代码。
clear:
	cd tools && go run ./gmsecgen clean -root ..

# make new service=myservice
# 从 example 模板生成一个新服务。复制 + 改名统一交给 tools/gmsecgen:
# 原来的实现是 cp -r 加十几条 sed,Windows 用 `sed -i`、macOS 用 `sed -i ''`、
# Linux 又是另一套参数,漏一条就留下半改不改的服务。
new:
	@test -n "$(service)" || (echo "用法: make new service=<服务名>" && exit 1)
	cd tools && go run ./gmsecgen scaffold -root .. -template example -service $(service)

# 生成 example 服务的 proto/gRPC 代码(需要先装好 protoc 与 protoc-gen-gmsec)
proto:
	cd example && $(MAKE) proto

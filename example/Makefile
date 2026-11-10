NAME := example

.PHONY: all build run windows mac linux clear clean-proto proto gen orm install source_install master

all: # 构建
	make clear
build:
	make clear
	# gormt 是外部可选二进制:存在才执行,不存在给提示并继续,不影响 go build。
	@if [ -x ./tools/gormt ]; then ./tools/gormt -o internal/model; \
	else echo "跳过 gormt: 未找到 ./tools/gormt(可选外部工具)"; fi
	go build -o $(NAME) main.go
	./$(NAME) debug
run:
	go build -o $(NAME) *.go
	./$(NAME) debug
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(NAME).exe main.go
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(NAME) main.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(NAME) main.go

# 只清理编译产物和 gormt 生成物。
# 注意这里**不能**删生成的 proto 代码:build 依赖 clear,
# 删了 rpc/<name> 之后 go build 会找不到包,历史实现就是这么坏的。
clear:
	test ! -d internal/model/ || rm -rf internal/model/*
	test ! -d err/ || rm -rf err/
	test ! -f $(NAME) || rm $(NAME)
	test ! -f $(NAME).exe || rm $(NAME).exe

# 只删除生成的 proto 代码,不碰手写代码;删完用 make proto 重新生成。
clean-proto:
	go -C ../tools run ./gmsecgen clean-proto -root .. -name $(NAME)

# 生成代码统一走 tools/gmsecgen(取代已不存在的 ./tools/protogen)。
proto:
	go -C ../tools run ./gmsecgen proto -root .. -name $(NAME)
gen:
	go -C ../tools run ./gmsecgen proto -root .. -name $(NAME)
orm: # gormt 生成 orm代码(可选)
	@if [ -x ./tools/gormt ]; then ./tools/gormt -o internal/model; \
	else echo "跳过 gormt: 未找到 ./tools/gormt(可选外部工具)"; fi

# 按 uname 分派到真实存在的安装脚本(历史写法 ../proto_install.sh 并不存在)。
install:
	@case "$$(uname -s)" in \
		MINGW*|MSYS*|CYGWIN*) bash ../proto_install_windows_git_bash.sh ;; \
		Darwin) bash ../proto_install_mac.sh ;; \
		*) bash ../proto_install_linux.sh ;; \
	esac
source_install:
	bash ../proto_source_install.sh

# 依赖升级:@master 不可复现,已移除。需要升级时把下面注释里的版本改成
# go.mod 中记录的固定版本后手动执行,禁止使用 go get -u ...@master。
#   go get github.com/xxjwxc/public@v0.0.0-20260818092048-1d99c237fc20
#   go get github.com/xxjwxc/ginrpc@v0.0.0-20250705164228-1d0f30c90f21
#   go get github.com/gmsec/micro@v0.0.0-20261109063142-72c07e589862
#   go get github.com/gmsec/goplugins@v0.0.0-20261109063246-61926158433f
master:
	@echo "禁止 @master 升级:请按 go.mod 里的固定版本显式指定后手动执行(见上方注释)"

NAME := example

.PHONY: all build run windows mac linux clear clean-proto proto gen orm install source_install master

all: # 构建
	make clear
build:
	make clear
	./tools/gormt -o internal/model
	go build -o $(NAME) main.go
	./$(NAME) debug
run:
	# make clear
	# ./tools/gormt -o internal/model
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

# 显式删除生成的 proto 代码(原 clear 的行为),删完记得 make proto 重新生成
clean-proto:
	go run ./tools/protogen -name $(NAME) -clean

# 代码生成统一走 Go 程序(tools/protogen):
# 旧链路是 make + `ls | cut` + `sed -i ''` + cp -rf,
# BSD sed/GNU sed 参数不同、Windows 没有 sed,且 cp 的目标目录两个平台分支还不一致。
proto:
	go run ./tools/protogen -name $(NAME) -apidoc ../apidoc -v
gen:
	go generate
orm: # gormt 生成 orm代码
	./tools/gormt -o internal/model
install:
	../proto_install.sh
source_install:
	../proto_install.sh
master:
	go get -u github.com/xxjwxc/public@master
	go get -u github.com/xxjwxc/ginrpc@master
	go get -u github.com/gmsec/micro@master
	go get -u github.com/gmsec/goplugins@master

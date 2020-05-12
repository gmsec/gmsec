NAME := gmsec
all: # 构建
	make clear
gen:
	- ln -s ./rpc ./gmsec/
	protoc --proto_path="./apidoc/proto/hello/" --gmsec_out=plugins=gmsec:./rpc/ hello.proto
clear:
	rm -rf ./rpc/*
install:
	./proto_install.sh
source_install:
	./proto_install.sh

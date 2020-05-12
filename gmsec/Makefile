NAME := gmsec
all: # 构建
	make clear
gen:
	ln -s ../rpc ./
	protoc --proto_path="../apidoc/proto/hello/" --gmsec_out=plugins=gmsec:./rpc/ hello.proto
install:
	./proto_install.sh
source_install:
	./proto_install.sh

NAME := gmsec
all: # 构建
gen:
	protoc --proto_path="./proto/hello" --gmsec_out=plugins=gmsec:./proto/hello hello.proto

clear: # 删除proto文件夹下所有go文件（谨慎操作）
	rm -rf ./proto/**/*.go


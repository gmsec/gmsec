SaverName := example
all: # 构建
	make clear
# gen:
# 	- mkdir ../rpc
# 	- mkdir ./rpc/$(SaverName)
# 	test -L rpc || ln -s ../rpc ./
# 	protoc --proto_path="../apidoc/proto/$(SaverName)/" --gmsec_out=plugins=gmsec:./rpc/$(SaverName) hello.proto
gen:
	test -L rpc || ln -s ../rpc ./
	go generate
orm: # gormt 生成 orm代码
	./tools/gormt -o internal/model
install:
	../proto_install.sh
source_install:
	../proto_install.sh
clear: # 删除proto文件夹下所有go文件（谨慎操作）
	rm -rf ./prc/$(SaverName)/*.go
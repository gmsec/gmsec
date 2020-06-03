install:
	./proto_install.sh
source_install:
	./proto_install.sh
clear: # 删除proto文件夹下所有go文件（谨慎操作）
	rm -rf ./prc/**/*.go

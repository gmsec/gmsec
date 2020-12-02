install:
	./proto_install.sh
source_install:
	./proto_install.sh
clear: # 删除proto文件夹下所有go文件（谨慎操作）
	rm -rf ./prc/**/*.go
new: #make new service=example
	echo start install $(service)
	# 开始替换文件
	test -d $(service) || cp -r ./example/ ./$(service)/
	find ./$(service)/ -type f -name "*.go" | xargs sed -i 's#example#$(service)#g' -i
	sed -i 's#example#$(service)#g' ./$(service)/go.mod
	sed -i 's#example#$(service)#g' ./$(service)/Makefile
	sed -i 's#example#$(service)#g' ./$(service)/generate/proto_makefile
	sed -i 's#example#$(service)#g' ./$(service)/conf/config.yml
	# 开始更新apidoc
	test -d ./apidoc/proto/$(service)/ || cp -r ./apidoc/proto/example/ ./apidoc/proto/$(service)/
	sed -i 's#example#$(service)#g' ./apidoc/proto/$(service)/*.proto
#!/bin/bash -x 

version="3.11.4"

export https_proxy=http://git.iizone.com.cn:32222 http_proxy=http://git.iizone.com.cn:32222 all_proxy=socks5://git.iizone.com.cn:32222
PROTOBUF_RELEASES_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-osx-x86_64.zip"

# su - xxj -c "qwer"
# download
curl -fLo protobuf.tar.zip ${PROTOBUF_RELEASES_URL}
mkdir protobuf-${version}
tar -xvf protobuf.tar.gz -C ./protobuf-${version}
cd protobuf-${version}

# install
xattr -c ./bin/protoc # mac 
cp -r ./bin/protoc $GOPATH/bin
cd ../
rm -rf protobuf-${version}/
rm -rf ./protobuf.tar.gz

# install go-grpc
go get -u google.golang.org/grpc
go get -u github.com/gmsec/protoc-gen-gmsec

chmod +x $GOPATH/bin/protoc

echo "SUCCESS"
#end
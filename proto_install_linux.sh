#!/bin/bash -x 

version="3.11.4"

ABROAD_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-linux-x86_64.zip"
INTERNAL_URL="https://github.com.cnpmjs.org/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-linux-x86_64.zip"

# su - xxj -c "qwer"
# download
curl -fLo protobuf.tar.zip ${ABROAD_URL} || curl -fLo protobuf.tar.zip ${INTERNAL_URL}
mkdir protobuf-${version}
unzip -d ./protobuf-${version} protobuf.tar.zip 
cd protobuf-${version}

# install
# xattr -c ./bin/protoc # mac 
cp -r ./bin/protoc $GOPATH/bin
cd ../
rm -rf protobuf-${version}/
rm -rf ./protobuf.tar.zip

# install go-grpc
go get -u google.golang.org/grpc
go get -u github.com/gmsec/protoc-gen-gmsec

chmod +x $GOPATH/bin/protoc

echo "SUCCESS"
#end
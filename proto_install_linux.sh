#!/bin/bash -x 

version="3.11.4"

# su - xxj -c "qwer"
# download
curl -fLo protobuf.tar.zip https://github.com/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-linux-x86_64.zip
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
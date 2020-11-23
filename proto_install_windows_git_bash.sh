#!/bin/bash -x 

version="3.11.4"

# su - xxj -c "qwer"
# download
curl -fLo protobuf.tar.gz https://github.com/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-win64.zip
mkdir protobuf-${version}
unzip -d ./protobuf-${version}/ protobuf.tar.gz
cd protobuf-${version}

# install
# xattr -c ./bin/protoc
cp -r ./bin/protoc.exe $GOPATH/bin
cd ../
rm -rf protobuf-${version}/
rm -rf ./protobuf.tar.gz

# install go-grpc
go get -u google.golang.org/grpc
go get -u github.com/gmsec/protoc-gen-gmsec

# chmod +x $GOPATH/bin/protoc

echo "SUCCESS"
#end
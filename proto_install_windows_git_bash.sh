#!/bin/bash -x 

version="3.11.4"

ABROAD_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-win64.zip"
INTERNAL_URL="https://github.com.cnpmjs.org/protocolbuffers/protobuf/releases/download/v${version}/protoc-${version}-win64.zip"

# su - xxj -c "qwer"
# download
curl -fLo protobuf.tar.gz ${ABROAD_URL} || curl -fLo protobuf.tar.zip ${INTERNAL_URL}
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
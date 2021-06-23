#!/bin/bash -x 

version="3.11.4"

ABROAD_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${version}/protobuf-all-${version}.tar.gz"
INTERNAL_URL="https://github.com.cnpmjs.org/protocolbuffers/protobuf/releases/download/v${version}/protobuf-all-${version}.tar.gz"

# su - xxj -c "qwer"
# download
curl -fLo protobuf.tar.zip ${ABROAD_URL} || curl -fLo protobuf.tar.zip ${INTERNAL_URL}
tar -xvf protobuf.tar.gz
cd protobuf-${version}

#build
curPath=$(pwd)
./configure --prefix=$curPath --enable-shared=no CFLAGS="-fPIC -fvisibility=hidden" CXXFLAGS="-fPIC -fvisibility=hidden" || exit 1
make clean
make -j6 || exit 1
make check
make install
xattr -c ./bin/protoc
cp -r ./bin/protoc $GOPATH/bin
cd ../
rm -rf protobuf-${version}/

# install go-grpc
go get -u google.golang.org/grpc
go get -u github.com/gmsec/protoc-gen-gmsec

echo "SUCCESS"
#end
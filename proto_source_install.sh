#!/usr/bin/env bash
# 从源码构建 protoc 并安装 Go 代码生成插件。版本说明见 proto_install_linux.sh。
# 注意:protobuf v33.0 的源码包已改用 CMake(不再带 autotools 的 ./configure),
# 这里对两种源码包都做了兼容。
set -euo pipefail

PROTOC_VERSION="33.0"
PROTOC_GEN_GO_VERSION="v1.36.11"
PROTOC_GEN_GMSEC_VERSION="v0.0.0-20261109063214-7e430756646d"

# 内网代理已移除:需要代理请由调用方 export HTTPS_PROXY,不设置就直连。
: "${HTTPS_PROXY:=}"
export HTTPS_PROXY https_proxy="${HTTPS_PROXY}"

GOBIN_DIR="$(go env GOBIN)"
if [ -z "${GOBIN_DIR}" ]; then
	GOBIN_DIR="$(go env GOPATH)/bin"
fi
if command -v cygpath >/dev/null 2>&1; then
	GOBIN_DIR="$(cygpath -m "${GOBIN_DIR}")"
fi
mkdir -p "${GOBIN_DIR}"
export GOBIN="${GOBIN_DIR}"

JOBS="$(getconf _NPROCESSORS_ONLN 2>/dev/null || echo 4)"
PROTOBUF_RELEASES_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protobuf-${PROTOC_VERSION}.tar.gz"

# 下载/解压都在临时目录,trap 清理;仓库里已被 git 跟踪的 protobuf.tar.zip 一概不碰。
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT

curl -fLo "${TMP_DIR}/protobuf.tar.gz" "${PROTOBUF_RELEASES_URL}"
mkdir -p "${TMP_DIR}/src"
tar -xzf "${TMP_DIR}/protobuf.tar.gz" -C "${TMP_DIR}/src"

SRC_DIR=""
for d in "${TMP_DIR}"/src/*/; do
	SRC_DIR="${d%/}"
	break
done
test -n "${SRC_DIR}" || { echo "源码包解压失败" >&2; exit 1; }

cd "${SRC_DIR}"
if [ -x ./configure ]; then
	./configure --prefix="${TMP_DIR}/install" --enable-shared=no CFLAGS="-fPIC" CXXFLAGS="-fPIC"
	make -j "${JOBS}"
	make install
	cp "${TMP_DIR}/install/bin/protoc" "${GOBIN}/protoc"
else
	cmake -S . -B build -Dprotobuf_BUILD_TESTS=OFF -Dprotobuf_BUILD_SHARED_LIBS=OFF -DCMAKE_BUILD_TYPE=Release
	cmake --build build --target protoc -j "${JOBS}"
	if [ -f build/protoc ]; then
		cp build/protoc "${GOBIN}/protoc"
	elif [ -f build/protoc.exe ]; then
		cp build/protoc.exe "${GOBIN}/protoc.exe"
	elif [ -f build/Release/protoc.exe ]; then
		cp build/Release/protoc.exe "${GOBIN}/protoc.exe"
	else
		echo "源码构建后未找到 protoc" >&2
		exit 1
	fi
fi

go install "google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}"
go install "github.com/gmsec/protoc-gen-gmsec@${PROTOC_GEN_GMSEC_VERSION}"

echo "SUCCESS"

#!/usr/bin/env bash
# 安装 protoc 与 Go 代码生成插件(linux x86_64)。
# 版本集中由下面三个变量管理,与仓库已提交的生成代码头部保持一致:
#   apidoc/rpc/example/hello.pb.go -> protoc-gen-go v1.36.11 / protoc v6.33.0
#   (protobuf 发布标签 v33.0 对应 `protoc --version` = libprotoc 33.0)
set -euo pipefail

PROTOC_VERSION="33.0"
PROTOC_GEN_GO_VERSION="v1.36.11"
PROTOC_GEN_GMSEC_VERSION="v0.0.0-20261109063214-7e430756646d"

# 内网代理已移除:需要代理请由调用方 export HTTPS_PROXY,不设置就直连。
: "${HTTPS_PROXY:=}"
export HTTPS_PROXY https_proxy="${HTTPS_PROXY}"

# 统一换算成 Git Bash / Windows Go 都能识别的路径(cygpath -m)。
GOBIN_DIR="$(go env GOBIN)"
if [ -z "${GOBIN_DIR}" ]; then
	GOBIN_DIR="$(go env GOPATH)/bin"
fi
if command -v cygpath >/dev/null 2>&1; then
	GOBIN_DIR="$(cygpath -m "${GOBIN_DIR}")"
fi
mkdir -p "${GOBIN_DIR}"
export GOBIN="${GOBIN_DIR}"

PROTOBUF_RELEASES_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"

# 下载到临时目录并 trap 清理;绝不删除仓库里已被 git 跟踪的 protobuf.tar.zip。
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "${TMP_DIR}"' EXIT

curl -fLo "${TMP_DIR}/protoc.zip" "${PROTOBUF_RELEASES_URL}"
unzip -q "${TMP_DIR}/protoc.zip" -d "${TMP_DIR}/protoc"
cp "${TMP_DIR}/protoc/bin/protoc" "${GOBIN}/protoc"
chmod +x "${GOBIN}/protoc"

# 固定版本安装(go install module@version 不会改动仓库的 go.mod/go.sum)。
go install "google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}"
go install "github.com/gmsec/protoc-gen-gmsec@${PROTOC_GEN_GMSEC_VERSION}"

test -x "${GOBIN}/protoc"
echo "SUCCESS"

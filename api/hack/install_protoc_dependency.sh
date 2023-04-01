#!/usr/bin/env bash

#
# install all protoc dependency.
#
set -ex

VERSION=3.20.3
PROTOC_VERSION=''

GOOS=$(go env GOOS)
if [[ "$GOOS" == "darwin" ]]; then
    	PROTOC_VERSION="$VERSION-osx"
else
    	PROTOC_VERSION="$VERSION-linux"
fi

wget https://ghproxy.com/https://github.com/protocolbuffers/protobuf/releases/download/v${VERSION}/protoc-${PROTOC_VERSION}-x86_64.zip
unzip protoc-${PROTOC_VERSION}-x86_64.zip
mv bin/protoc /usr/local/bin
rm -rf protoc-${PROTOC_VERSION}-x86_64.zip protoc-${PROTOC_VERSION}-x86_64/

protoc --version

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.10.0 &&\
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.0 &&\
go install github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts@v1.1.1 &&\
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 &&\
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 &&\
go install github.com/rakyll/statik@v0.1.7 &&\
go install github.com/envoyproxy/protoc-gen-validate@v0.6.7

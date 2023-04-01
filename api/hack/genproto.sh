#!/usr/bin/env bash
#
# Generate all protobuf bindings.
#
set -ex
DIRS=("types" "plugin")
GEN_DIR=gen

mkdir -p ${GEN_DIR}/hub.io/api

for dir in ${DIRS[@]}
do
  rm -rf ${GEN_DIR}/hub.io/api/${dir}
  rsync -a ./${dir} ${GEN_DIR}/hub.io/api
done
pushd "${GEN_DIR}"


for dir in ${DIRS[@]}
do
  for var in `find ./hub.io/api/${dir} -name "*.proto"`
  do
    protoc \
    -I . \
    -I ../third_party/ \
    -I ../vendor/ \
    --go_out . \
    --go-grpc_out . \
    --validate_out="lang=go:." \
    --grpc-gateway_out . --grpc-gateway_opt=logtostderr=true,allow_delete_body=true \
    "${var}";
  done
done

popd

# Copy back generated *.go files to the api root dir, so that `go mod vendor` can find the packages
# with directives replace( amamba.io/api/ => ./api )
rsync -a ${GEN_DIR}/hub.io/api/* ./
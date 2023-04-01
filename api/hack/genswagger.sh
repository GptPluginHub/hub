#!/usr/bin/env bash
#
# Generate all swagger bindings.
#
set -ex

rm -rf assets
mkdir -p assets/swagger

GEN_DIR=gen

for var in `find ${GEN_DIR}/hub.io -name "*.proto" `
do
  protoc  \
    -I ./ \
    -I gen/ \
    -I third_party/ \
    -I vendor/ \
    --openapiv2_out ./assets/swagger \
    --openapiv2_opt logtostderr=true,repeated_path_param_separator=ssv,allow_delete_body=true \
    "${var#gen/}"   # removes the prefixed dir `gen/` in the path
done

cp -r third_party/swagger-ui/* ./assets/swagger/

DEMO_SWAGGER_URL="https://petstore.swagger.io/v2/swagger.json"
INDEX_SWAGGER_URL="./hub.io/api/plugin/v1alpha1/rpc.swagger.json"

if [ "$(uname)" == "Darwin" ]; then
  sed -i '' "s@$DEMO_SWAGGER_URL@$INDEX_SWAGGER_URL@g"  ./assets/swagger/index.html
else
  sed -i  "s@$DEMO_SWAGGER_URL@$INDEX_SWAGGER_URL@g"  ./assets/swagger/index.html
fi


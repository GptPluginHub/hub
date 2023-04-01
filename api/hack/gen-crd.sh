#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo "generating crd yaml files" 
CONTROLLER_GEN=../bin/controller-gen
CRD_API_PATH=./crd/apis
CRD_OUTPUT_DIR=../charts/crds

cmd="${CONTROLLER_GEN} paths=${CRD_API_PATH}/... crd output:crd:dir=${CRD_OUTPUT_DIR}"
echo $cmd
$cmd

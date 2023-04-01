#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

DIFFPROTO="../vendor/hub.io"
TMP_DIFFPROTO="../_tmp"

cleanup() {
  rm -rf "${TMP_DIFFPROTO}"
}
trap "cleanup" EXIT SIGINT

cleanup

mkdir -p "${TMP_DIFFPROTO}"
cp -a "${DIFFPROTO}"/* "${TMP_DIFFPROTO}"
echo ${SCRIPT_ROOT}
bash -c "make genproto"
echo "diffing ${DIFFPROTO} against freshly generated files"

ret=0

echo "${DIFFPROTO}"
echo "${TMP_DIFFPROTO}"

diff -x sdk -x "rpc.pb.gw.go" -Naupr "${DIFFPROTO}" "${TMP_DIFFPROTO}" || ret=$?
cp -a "${TMP_DIFFPROTO}"/* "${DIFFPROTO}"
if [[ $ret -eq 0 ]]
then
  echo "${DIFFPROTO} up to date."
else
  echo "${DIFFPROTO} is out of date. Please run make genproto"
  exit 1
fi

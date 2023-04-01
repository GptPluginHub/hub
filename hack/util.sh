#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

MIN_Go_VERSION=go1.16.0

# This function installs a Go tools by 'go get' command.
# Parameters:
#  - $1: package name, such as "sigs.k8s.io/controller-tools/cmd/controller-gen"
#  - $2: package version, such as "v0.4.1"
# Note:
#   Since 'go get' command will resolve and add dependencies to current module, that may update 'go.mod' and 'go.sum' file.
#   So we use a temporary directory to install the tools.
function util::install_tools() {
	local package="$1"
	local version="$2"

	temp_path=$(mktemp -d)
	pushd "${temp_path}" >/dev/null
	GO111MODULE=on go install "${package}"@"${version}"
	GOPATH=$(go env GOPATH | awk -F ':' '{print $1}')
	export PATH=$PATH:$GOPATH/bin
	popd >/dev/null
	rm -rf "${temp_path}"
}


function util::verify_go_version {
    local go_version
    IFS=" " read -ra go_version <<< "$(GOFLAGS='' go version)"
    if [[ "${MIN_Go_VERSION}" != $(echo -e "${MIN_Go_VERSION}\n${go_version[2]}" | sort -s -t. -k 1,1 -k 2,2n -k 3,3n | head -n1) && "${go_version[2]}" != "devel" ]]; then
      echo "Detected go version: ${go_version[*]}."
      echo "kangaroo requires ${MIN_Go_VERSION} or greater."
      echo "Please install ${MIN_Go_VERSION} or later."
      exit 1
    fi
}

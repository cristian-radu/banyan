#! /bin/bash

# We use k8s.io/code-generator for generating most of the Kubernetes related code

set -e

declare -r projectRoot=$(dirname $(dirname $(readlink -fm $0)))
declare -r projectOwner="github.com/cristian-radu"
declare -r binaryName="banyan"
declare -r apiVersion="v1alpha1"
declare -r apisPath="${projectOwner}/${binaryName}/pkg/apis"

function downloadCodeGenerator() {
    go mod download k8s.io/code-generator
}

function runCodeGenerator() {
    declare -r codeGenDirName="k8s.io/code-generator@$(grep "replace k8s.io/code-generator" go.mod | awk '{print $5}')"
    declare -r codeGenDirPath="${GOPATH}/pkg/mod/${codeGenDirName}"

    # Remove the old generated code
    rm -rf ${projectRoot}/pkg/generated

    # Run the generators
    chmod +x ${codeGenDirPath}/generate-groups.sh
	${codeGenDirPath}/generate-groups.sh all ${projectOwner}/${binaryName}/pkg/generated ${apisPath} ${binaryName}:${apiVersion} -o ${projectRoot}
}

function cleanup() {
    # move the generated code files to their desired locations
	mv ${projectRoot}/${apisPath}/${binaryName}/${apiVersion}/zz_generated.deepcopy.go ${projectRoot}/pkg/apis/${binaryName}/${apiVersion}
	mv ${projectRoot}/${projectOwner}/${binaryName}/pkg/generated ${projectRoot}/pkg
	rm -rf ${projectRoot}/github.com
}

downloadCodeGenerator
runCodeGenerator
cleanup
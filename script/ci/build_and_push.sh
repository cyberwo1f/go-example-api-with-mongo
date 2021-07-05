#!/bin/bash

set -eu
set -o pipefail

DOCKER=docker

main(){
    local prefix tag ref
    ref=$(echo "${GITHUB_REF##refs/tags/}")
    echo "[INFO] Got ref ${ref}"

    case "${1}" in
        "tag")
            tag="${ref}"
            ;;
        "commit")
            sha=$(git rev-parse --short "${GITHUB_SHA}")
            if [[ "${ref}" = "master" ]]; then
                prefix="master"
            else
                prefix="develop"
            fi
            tag="${prefix}-${sha}"
            ;;
        *)
            echo "[ERROR] Not supported type ${1}" >&2
            exit 1
    esac

    echo "[INFO] Build docker image with ${ECR_REGISTRY}/${ECR_REPOSITORY}:${tag}"
    "${DOCKER}" build --build-arg VERSION="${tag}" --build-arg GITHUB_TOKEN="${GITHUB_TOKEN}" . -t "${ECR_REGISTRY}/${ECR_REPOSITORY}:${tag}"
    "${DOCKER}" push "${ECR_REGISTRY}/${ECR_REPOSITORY}:${tag}"
}

main "$@"

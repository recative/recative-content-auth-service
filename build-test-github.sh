#!/bin/bash
set -e
export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

ENV="test"

BASE_DIR=$(cd "$(dirname "$0")"; pwd)
RUN_DIR="${BASE_DIR}/cmd"
DOCKER_DIR="${BASE_DIR}/docker/${ENV}"
DIST_DIR="${DOCKER_DIR}/dist"

REGISTRY_HOST="ghcr.io"
REGISTRY_NAMESPACE="recative"

TAG="$(date +'%y%m%d%H%M')"

APP_NAME="$1"
IMAGE_NAME="${REGISTRY_HOST}/${REGISTRY_NAMESPACE}/${APP_NAME}"
IMAGE_FULL_TAG="${IMAGE_NAME}:${TAG}"
IMAGE_LATEST_TAG="${IMAGE_NAME}:latest"

rm -rf "${DIST_DIR:?}"
mkdir "${DIST_DIR:?}"

prepare(){
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
  go generate ./...
}

serverBuild(){
    echo "------> build ${APP_NAME} ${VERSION}"

	  if [[ -f "${RUN_DIR}/main.go" ]]; then
	    go build -v -o "${DIST_DIR}/${APP_NAME}" "${RUN_DIR}/main.go"
    fi
}

dockerBuild(){
    echo "Docker Build"
    if [[ -f ${DOCKER_DIR}/Dockerfile ]];then
        if [[ -f ${DIST_DIR}/${APP_NAME} ]]; then
            echo docker building version: "${TAG}"
            docker build -t "${IMAGE_FULL_TAG}" \
                         -t "${IMAGE_LATEST_TAG}" \
                         -f "${DOCKER_DIR}/Dockerfile" \
                         --platform=linux/amd64 \
                         "${DOCKER_DIR}"
        fi
    fi
}

pushImage(){
    echo "Push Image"
    echo "${IMAGE_FULL_TAG}"
    if [[ $(docker images | grep -c "${IMAGE_NAME}") -eq 2 ]]; then
        docker push --all-tags "${IMAGE_NAME}"
        docker rmi "${IMAGE_FULL_TAG}"
        docker rmi "${IMAGE_LATEST_TAG}"
    fi
}

prepare
serverBuild
dockerBuild
pushImage

rm -rf "${DIST_DIR:?}"

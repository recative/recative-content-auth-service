#!/bin/bash
set -e
export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

ENV="prod"

BASE_DIR=$(cd "$(dirname "$0")"; pwd)
RUN_DIR="${BASE_DIR}/cmd"
DOCKER_DIR="${BASE_DIR}/docker/${ENV}"
DIST_DIR="${DOCKER_DIR}/dist"

REGISTRY_HOST="711643798364.dkr.ecr.us-east-1.amazonaws.com"
REGISTRY_NAMESPACE="/nicestick/"

TAG="$1"

APP_NAME="recative-content-auth-service"
IMAGE_NAME="${REGISTRY_HOST}${REGISTRY_NAMESPACE}${APP_NAME}"
IMAGE_FULL_TAG="${IMAGE_NAME}:${TAG}"
IMAGE_LATEST_TAG="${IMAGE_NAME}:stable"

rm -rf "${DIST_DIR:?}"
mkdir "${DIST_DIR:?}"

gitCheck(){
    git fetch --tags
    echo "Git Check"
    if ! [[ $(git tag | grep -c "${TAG}") -eq 1 ]]; then
        echo "tag ${TAG} not exists"
        exit
    fi
    git checkout "${TAG}"
}

serverBuild(){
    echo "------> build ${APP_NAME} ${VERSION}"

	  if [[ -f "${RUN_DIR}/main.go" ]]; then
	    go build -v -o "${DIST_DIR}/${APP_NAME}" "${RUN_DIR}/main.go"
    fi
}

baseDataBuild(){
    echo "------> build base data"

    go build -v -o "${DIST_DIR}/base" "${RUN_DIR}/base/base.go"
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
#    if [[ $(docker images | grep -c "${IMAGE_NAME}") -eq 1 ]]; then
        docker push "${IMAGE_FULL_TAG}"
        docker push "${IMAGE_LATEST_TAG}"
#        docker rmi "${IMAGE_FULL_TAG}"
#    fi
}

gitCheck
serverBuild
baseDataBuild
dockerBuild
pushImage

rm -rf "${DIST_DIR:?}"

#!/bin/bash

BASE_DIR=$(cd "$(dirname "$0")";cd .. || exit; pwd)

DOCKER_DIR="${BASE_DIR}/docker/api-test"
RUN_DIR="${BASE_DIR}/cmd"
DIST_DIR="${DOCKER_DIR}/dist"
TEST_DIR="${BASE_DIR}/test"

APP_NAME="recative-backend"

init(){
  rm -rf "${DIST_DIR:?}"
  mkdir "${DIST_DIR:?}"
}
init

serverBuild(){
  export GOOS=linux
  echo "------> build ${APP_NAME}"
  go mod tidy

  if [[ -f "${RUN_DIR}/main.go" ]]; then
    go build -o "${DIST_DIR}/${APP_NAME}" "${RUN_DIR}"
  fi
}
serverBuild

runDockerCompose(){
  if [[ -f ${DIST_DIR}/${APP_NAME} ]]; then
      docker-compose -f "${DOCKER_DIR}/docker-compose.yml" down
      docker-compose -f "${DOCKER_DIR}/docker-compose.yml" up -d --build
  fi
}
runDockerCompose

apiTest(){
  sleep 15
  cd "$TEST_DIR" && npm install && npm run test:update:snapshot && npm run test
}
apiTest



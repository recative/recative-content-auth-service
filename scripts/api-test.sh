#!/bin/bash
set -e

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
  # important without this, binary will brake in alpine
  export CGO_ENABLED=0
  export GOOS=linux
  echo "------> build ${APP_NAME}"

  go generate ./...
  go mod tidy
  go build -o "${DIST_DIR}/${APP_NAME}" "${RUN_DIR}"
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
  i=0
  while ! nc -z localhost 12211; do
    echo "waiting for server to start: $i times"
    if [[ "$i" == 10 ]]; then
      echo "timeout"
      exit 1
    fi
    sleep 0.5 # wait for 1/10 of the second before check again
    ((i++))
  done
  cd "$TEST_DIR" && npm install && npm run test:update:snapshot && npm run test
}
apiTest



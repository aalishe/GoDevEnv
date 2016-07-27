#!/bin/bash

USERNAME='johandry'
IMG_NAME='godevenv'

FULL_IMG_NAME="${USERNAME}/${IMG_NAME}"

build() {
  echo -e "\033[93;1mBuilding the image\033[0m"
  docker build -t ${FULL_IMG_NAME} .
}

if [[ "$1" == "--build" ]]
then
  build
  shift
fi

if [[ "$1" == "--push" ]]
then
  # If the image is not there, build it
  docker images | grep -q ${FULL_IMG_NAME} || build

  # Login to DockerHub if you are not
  cat  ${HOME}/.docker/config.json | tr -d '\n' | grep -q '"https://index.docker.io/.*":.*{.*"auth": ".\{1,\}"' || (
    echo -e "\033[93;1mLogin to DockerHub as ${USERNAME}\033[0m."
    docker login -u ${dkr_username}
  )

  # Push the new image
  echo -e "\033[93;1mPushing the new image to DockerHub\033[0m"
  docker push ${FULL_IMG_NAME}
fi

echo -e "\033[93;1mRunning the container\033[0m"
docker run --rm -it --name "${IMG_NAME}" -v "${PWD}/workspace":/root/workspace -w /root/workspace ${FULL_IMG_NAME}

[[ $? -eq 125 ]] && echo -e "\033[91;1mERROR\033[0m: Build the image and (optional) push it with $0 --build [ --push ]"

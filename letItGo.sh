#!/bin/bash

USERNAME='johandry'
IMG_NAME='godevenv'

FULL_IMG_NAME="${USERNAME}/${IMG_NAME}"

# If the image is not there, build it
if ! docker images | grep -q ${FULL_IMG_NAME}
then
  echo -e "\033[93;1mBuilding the image\033[0m"
  docker build -t ${FULL_IMG_NAME} .

  egrep -q '"auth": ".+"' ~/.docker/config.json || (
    echo -e "\033[93;1mLogin to DockerHub as ${USERNAME}\033[0m"
    docker login -u ${USERNAME}
  )

  # TODO: Should I push the image to DockerHub? Is there a way if it is outdated?
  # docker push ${FULL_IMG_NAME}
fi

echo -e "\033[93;1mRunning the container\033[0m"
docker run --rm -it --name "${IMG_NAME}" -v "${PWD}/workspace":/root/workspace -w /root/workspace ${FULL_IMG_NAME}

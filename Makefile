#=======================================================================================================
# Author: Johandry Amador <johandry@gmail.com>
# Title:  Makefile to build, push to DockerHub, run and clean a container for Go
#
# Usage: make <option>
#
# Options:
#     help      Print this help
#     run       Startup the container using the image from DockerHub or locally builded.
#     exec cmd  Execute the command 'cmd' into the container and exit.
#     build     Build the image
#     push      Login to DockerHub (if you are not) and push the existing image
#     release   Build and Push the recently created image
#     clean     Remove every image stored related to Go
#
# Description:
#     Execute 'make run' to startup the Go container and start developing. Make
#     will download the latest image from DockerHub.
#     If you modify the Dockerfile or anything on the imgage, then execute
#     'make build' or 'make release'.
#     When you finish, you may clean up everything with 'make clean'.
#
# More information, report Issues or create Pull Requests in http://github.com/johandry/GoDevEnv
#=======================================================================================================

USERNAME = johandry
IMG_NAME = godevenv
VERSION := $(shell cat VERSION)

# From the base image (https://hub.docker.com/_/golang/)
GOPATH=/go

VOLUMES = -v "${PWD}/workspace":${GOPATH}
WORKDIR = -w ${GOPATH}
ENV 		=
PORTS   =
#PORTS   = -p 127.0.0.1:6060:6060

.PHONY: build login push release run exec clean help all

# If the first argument is "exec"...
ifeq (exec,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "exec"
  EXEC_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(EXEC_ARGS):;@:)
endif

build:
	@echo "\033[93;1mBuilding the image $(USERNAME)/$(IMG_NAME):$(VERSION)\033[0m"
	git pull
	docker build --rm -t $(USERNAME)/$(IMG_NAME):latest .

tag:
	@echo "\033[93;1mTagging the new image with version $(VERSION)\033[0m"
	git add -A
	git commit -am "Version: $(VERSION)"
	git tag -a "$(VERSION)" -m "Version: $(VERSION)"
	git push
	git push --tags
	docker tag $(USERNAME)/$(IMG_NAME):latest $(USERNAME)/$(IMG_NAME):$(VERSION)

login:
	@if ! docker info | grep -q 'Username: $(USERNAME)'; then \
		echo "\033[93;1mLogin to DockerHub as ${USERNAME}\033[0m."; \
		docker login -u $(USERNAME); \
 	 fi

push: login
	@echo "\033[93;1mPushing the new image to DockerHub\033[0m"
	docker push $(USERNAME)/$(IMG_NAME):latest
	docker push $(USERNAME)/$(IMG_NAME):$(VERSION)

release: build tag push

exec:
	@echo "\033[93;1mExecuting $(EXEC_ARGS) into the existing container $(IMG_NAME)\033[0m"
	docker exec -it $(IMG_NAME) $(EXEC_ARGS)

run:
	@if docker ps | grep -q godevenv; then \
		echo "\033[93;1mLogin to the existing container $(IMG_NAME)\033[0m"; \
		docker exec -it $(IMG_NAME) /bin/bash --login; \
	else \
		echo "\033[93;1mRunning the container $(IMG_NAME)\033[0m"; \
		docker run --rm -it --name $(IMG_NAME) $(VOLUMES) $(WORKDIR) $(ENV) $(PORTS) $(USERNAME)/$(IMG_NAME):$(VERSION); \
	fi

all: release run

clean:
	@echo "\033[93;1mDeleting the created images\033[0m"
	docker rmi $(USERNAME)/$(IMG_NAME):$(VERSION)
	docker rmi `grep ^FROM Dockerfile | cut -d\  -f2`

help:
	@sed -ne '/^# Usage/,/^# More/p' Makefile | sed -e 's/^#\(.*\)/\1/' | sed -e 's/^ \(.*\)/\1/'

default: build

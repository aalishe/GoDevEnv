USERNAME = johandry
IMG_NAME = godevenv
VERSION ?= latest

VOLUMES = -v "${PWD}/workspace":/root/workspace
WORKDIR = -w /root/workspace
ENV 		=
PORTS   =

.PHONY: build login push release run clean

build:
	@echo "\033[93;1mBuilding the image $(USERNAME)/$(IMG_NAME):$(VERSION)\033[0m"
	docker build --rm -t $(USERNAME)/$(IMG_NAME):$(VERSION) .

login:
	@if ! docker info | grep -q 'Username: $(USERNAME)'; then \
		echo "\033[93;1mLogin to DockerHub as ${USERNAME}\033[0m."; \
		docker login -u $(USERNAME); \
 	 fi

push: login
	@echo "\033[93;1mPushing the new image to DockerHub\033[0m"
	docker push $(USERNAME)/$(IMG_NAME):$(VERSION)

release: build push

run:
	@echo "\033[93;1mRunning the container $(IMG_NAME)\033[0m"
	docker run --rm -it --name $(IMG_NAME) $(VOLUMES) $(WORKDIR) $(ENV) $(PORTS) $(USERNAME)/$(IMG_NAME):$(VERSION)

clean:
	@echo "\033[93;1mDeleting the created images\033[0m"
	docker rmi $(USERNAME)/$(IMG_NAME):$(VERSION)
	docker rmi `grep ^FROM Dockerfile | cut -d\  -f2`

default: build
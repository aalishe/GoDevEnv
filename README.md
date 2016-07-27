# Go Language Development and Learning Environment

This is a docker container to learn and develop in Go language without the need to install Go in your computer. It can also be used - eventually - to containerize your Go application.

## Who to use?

To login to the container and start using Go, just execute `make run`. The image will be pulled (if not done before) and run the container. After that you will see the container prompt.

```bash
$ make run
Running the container
root@8dac5ea8da15:~/workspace#

```

## Build and Push

Execute `make` to build the image and keep it in locally.

```bash
make
```

To push the image to DockerHub use the `push` parameter:

```bash
make push
```

But, if you want to build and push it to DockerHub, it is better to use the `release` parameter:

```bash
make release
```

## Cleanup

The container will be deleted after you exit from it. If not, list the running containers and remove it.

```bash
docker ps -a
docker rm <id>
```

To remove the image:

```bash
make clean
```

## Customize

Modify the DockerHub account and image or repository name in the `Makefile` file:

```make
USERNAME = johandry
IMG_NAME = godevenv
```

Also the volumes to share, the working directory, any environment variable you want to set in the container and any port to map between the localhost and the container.

```make
VOLUMES = -v "${PWD}/workspace":/root/workspace
WORKDIR = -w /root/workspace
ENV 		= -e SOME_KEY=SOME_VALUE OTHER_KEY=OTHER_VALUE
PORTS   = -p 8080:80
```

It is possible to modify the user profile editing the `.bashrc` file and re-building the image.

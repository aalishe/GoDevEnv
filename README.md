# Go Language Development and Learning Environment

This is a docker container to learn and develop in Go language without the need to install Go in your computer. It can also be used - eventually - to dockerize your Go application.

## How to use?

To login to the container and start using Go, just execute `make run`. The image will be pulled (if not done before) and run the container. After that you will see the container prompt.

```bash
$ make run
Running the container godevenv
root@8dac5ea8da15:~/workspace#

```

For more options continue reading or execute `make help`.

## Another terminal and execute commands

To have a second, third, or Nth terminal, execute `make run` again. The `make run` is used to start up the container for the first time but if used again in other terminal will login to the container. This is useful to have a console when a process is running in the container.

```bash
make run
Login to the existing container godevenv
root@92f064f5cf3f:~/workspace# tail -f /tmp/output
```

Sometimes there is no need to login into the container to execute a command. That can be done with `make exec` and the command(s) to execute.

```bash
make exec tail /tmp/output
```

However, if the command's parameters use dashes (- or --), redirects (< or >), etc... then quote them.

```bash
# Assuming there is a container running a process:
make exec 'tail -f /tmp/output | grep ERROR'
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
ENV     = -e SOME_KEY=SOME_VALUE OTHER_KEY=OTHER_VALUE
PORTS   = -p 8080:80
```

If you need more packages in the container, modify the `PACKAGES` environment variable in the Dockerfile.

```docker
ENV PACKAGES 'vim less lynx links'
```

It is possible to modify the user profile editing the `.bashrc` file. A modification to this file and the Dockerfile will require re-building the image.

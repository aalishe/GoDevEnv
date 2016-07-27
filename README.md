# Go Language Development and Learning Environment

This is a docker container to learn and develop in Go language without the need to install Go in your computer. It can also be used - eventually - to containerize your Go application.

## Who to use?

To login to the container and start using Go, just execute `letItGo.sh` script. The image will be pulled (if not done before) and run the container. Next think you will see is the container prompt.

```bash
$ ./letItGo.sh
Running the container
root@8dac5ea8da15:~/workspace#

```

## Build and Push

Execute `letItGo.sh` with the parameter `--build` to build the image and keep it in locally.

```bash
./letItGo.sh --build
```

To push the image use the `--push` parameter:

```bash
./letItGo.sh --push
```

The `--push` parameter will also build and push if the image is not found locally.

## Cleanup

This is not implemented in the script.

The container will be deleted after you exit from it. If not, list the running containers and remove it.

```bash
docker ps -a
docker rm <id>
```

To remove the image:

```bash
docker rmi johandry/godevenv
docker images
```

## Customize

Modify the DockerHub account and image name in the `letItGo.sh` script:

```bash
USERNAME='johandry'
IMG_NAME='godevenv'
```

It is also possible to modify the user profile modifying the `.bashrc` file.

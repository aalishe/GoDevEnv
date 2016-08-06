FROM golang:latest

MAINTAINER "Johandry Amador" <johandry@gmail.com>

COPY .bashrc /root/.bashrc

ENV PACKAGES 'vim less lynx links'

RUN apt-get update && \
    apt-get install -y ${PACKAGES} && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /root/workspace

CMD ["/bin/bash","--login"]

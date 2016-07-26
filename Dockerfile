FROM golang:1.6

MAINTAINER "Johandry Amador" <johandry@gmail.com>

COPY .bashrc /root/.bashrc

ENV PACKAGES 'vim less'

RUN apt-get update && \
    apt-get install -y ${PACKAGES} && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /root/workspace

CMD ["/bin/bash","--login"]

FROM golang:latest

MAINTAINER "Johandry Amador" <johandry@gmail.com>

ADD VERSION .

COPY .bashrc /root/.bashrc

ENV PACKAGES 'vim less lynx links'

RUN apt-get update && \
    apt-get install -y ${PACKAGES} && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* \
    godoc -http=:6060 &

EXPOSE 6060

CMD ["/bin/bash","--login"]

# https://hub.docker.com/_/golang
FROM golang:1.10-alpine

LABEL MAINTAINER Riya Dennis


RUN apk add --update --no-cache \
        ca-certificates \
        # https://github.com/Masterminds/glide#supported-version-control-systems
        git mercurial subversion bzr \
        openssh \
 && update-ca-certificates \
    \
 # Install build dependencies
 && apk add --no-cache --virtual .build-deps \
        curl make

###Install dep##################################
WORKDIR /root
RUN wget https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-amd64.zip
RUN unzip dep-linux-amd64.zip -d /usr/bin
ENV PATH="/root/bin:${PATH}"

WORKDIR $GOPATH
RUN mkdir -p "$GOPATH/src/github.com" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN mkdir $GOPATH/src/github.com/prog-image

ADD . $GOPATH/src/github.com/prog-image/

WORKDIR $GOPATH/src/github.com/prog-image

RUN dep ensure

RUN go build -i -o prog-image .

EXPOSE 8080

ENTRYPOINT ./prog-image server -config=config.yaml
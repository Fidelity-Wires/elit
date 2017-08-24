FROM golang:1.8-alpine
LABEL maintainer "UshioShugo<ushio.s@gmail.com>"

ENV APP_PATH=${GOPATH}/src/github.com/ushios/elit

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

RUN apk add --no-cache --virtual .goget \
	git openssh-client && \
	go get -v ./... && \
	apk del .goget

FROM golang:1.17-alpine

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/swayne275/joke-web-server
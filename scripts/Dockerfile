FROM golang:latest

ADD . /go/src/github.com/jhidalgo3/microservice-go
WORKDIR /go/src/github.com/jhidalgo3/microservice-go


RUN apt-get update && \
    apt-get install -y vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN go get ./... && \
    go install github.com/jhidalgo3/microservice-go

ENTRYPOINT scripts/go-app-entrypoint.sh

EXPOSE 8080
# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN apk add build-base
WORKDIR /

COPY go.sum ./
COPY go.mod ./
RUN go mod download

COPY /config /config
COPY main.go /

RUN go build -o helloworld

EXPOSE 8080

CMD [ "/helloworld" ]
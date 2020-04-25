FROM golang:latest

RUN apt-get update && apt-get install -y cron

ADD . /go/src/github.com/kg0r0/ctftime-bot

WORKDIR /go/src/github.com/kg0r0/ctftime-bot

RUN go get github.com/nlopes/slack
RUN go build
RUN ./ctftime-bot


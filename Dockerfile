FROM golang:latest

ADD . /go/src/github.com/kg0r0/ctftime-bot

WORKDIR /go/src/github.com/kg0r0/ctftime-bot

RUN go github.com/nlopes/slack
RUN go build
RUN ./ctftime-bot


FROM golang:latest

RUN apt-get update && apt-get install -y cron

RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

ADD . /go/src/github.com/kg0r0/ctftime-bot

WORKDIR /go/src/github.com/kg0r0/ctftime-bot

RUN go github.com/nlopes/slack
RUN env > /env

CMD echo '*/1 * * * * cd /go/src/github.com/kg0r0/ctftime-bot; env - `cat /env` go run /go/src/github.com/kg0r0/ctftime-bot/main.go >> /var/log/cron.log 2>&1' | crontab - && cron -f

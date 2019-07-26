package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kg0r0/ctftime-bot/ctftime"
	"github.com/nlopes/slack"
	"github.com/robfig/cron"
	"github.com/takama/daemon"
)

type service struct {
	daemon.Daemon
}

const (
	name        = "ctftime-bot"
	description = "A bot for notifying ctf date to slack"
)

var c, t *string

func (servier *service) manage(api *slack.Client, attachment slack.Attachment) (string, error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	cron := cron.New()
	cron.AddFunc("*/5 * * * * *", func() {
		channelID, timestamp, err := api.PostMessage(*c, slack.MsgOptionText("Upcoming events", false), slack.MsgOptionAttachments(attachment))
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	})

	cron.Start()
	killSignal := <-interrupt
	log.Println("Got signal:", killSignal)
	return "Service exited", nil
}

func main() {
	config, err := ctftime.NewConfig("conf/ctftime_conf.json")
	if err != nil {
		log.Fatal("Can not read config file!")
	}
	if config.SlackConfig.APIToken == "" {
		log.Fatal("Can not find AccessToken!")
	}
	t = flag.String("t", config.SlackConfig.APIToken, "api token")
	c = flag.String("c", config.SlackConfig.ChannelID, "channel id")
	flag.Parse()
	if len(*t) == 0 {
		log.Fatal("APIToken is Invalid!")
	}

	if len(*c) == 0 {
		log.Fatal("ChannelID is Invalid!")
	}

	info, err := ctftime.GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	contents := []string{}
	for _, i := range info {
		contents = append(contents, fmt.Sprintf(":trophy: *%s* \n *url* : %s\n *format* : %s\n *start* : %s\n *finish* :%s\n", i.Title, i.URL, i.Format, i.Start.Local(), i.Finish.Local()))
	}
	api := slack.New(*t)
	attachment := slack.Attachment{
		Text: strings.Join(contents[:], "\n"),
	}
	srv, err := daemon.New(name, description)
	if err != nil {
		log.Fatal(err)
	}
	service := &service{srv}
	service.manage(api, attachment)
}

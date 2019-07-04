package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/kg0r0/ctftime-bot/ctftime"
	"github.com/nlopes/slack"
)

func main() {
	config, err := ctftime.NewConfig("conf/ctftime_conf.json")
	if err != nil {
		log.Fatal("Can not read config file!")
	}
	if config.SlackConfig.AccessToken == "" {
		log.Fatal("Can not find AccessToken!")
	}
	t := flag.String("t", config.SlackConfig.AccessToken, "access token")
	c := flag.String("c", config.SlackConfig.ChannelID, "channel id")
	flag.Parse()
	if len(*t) == 0 {
		log.Fatal("AccessToken is Invalid!")
	}

	if len(*c) == 0 {
		log.Fatal("ChannelID is Invalid!")
	}

	info := ctftime.GetInfo()
	contents := []string{}
	for _, i := range info {
		contents = append(contents, fmt.Sprintf("[%s](%s)\n *start* : %s\n *finish* :%s\n", i.Title, i.URL, i.Start, i.Finish))
	}
	api := slack.New(*t)
	attachment := slack.Attachment{
		Text: strings.Join(contents[:], "\n"),
	}
	channelID, timestamp, err := api.PostMessage(*c, slack.MsgOptionText("直近2週間のCTF", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

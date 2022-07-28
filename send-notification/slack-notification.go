package main

import (
	"log"
	"os"

	"github.com/nicholasjackson/env"
	"github.com/slack-go/slack"
)



var token = env.String("SLACK_BOT_AUTH_TOKEN", false, "auth_token", "Slack Bot Auth Token")

func main() {

	l := log.New(os.Stdout, "golang-slack-bot ", log.LstdFlags)

	err := env.Parse()
	if err != nil {
		l.Println("error parsing SLACK_BOT_AUTH_TOKEN")
		return
	}

	args := os.Args[1:]
	log.Println(args)

	client := slack.New(*token)

	preText := "*Hello! Your Jenkins build has finished!*"

	jenkinsURL := "*Build URL:* " + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + ":x:"
	}

	dividerBlock1 := slack.NewDividerBlock()

	preTextField := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	jenkinsBuldDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
	jenkinsBuldDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuldDetails, false, false)
	jenkinsBuldDetailsSection := slack.NewSectionBlock(jenkinsBuldDetailsField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerBlock1,
		jenkinsBuldDetailsSection,
	)

	_, _, _, err = client.SendMessage(
		"channel_id",
		msg,
	)
	if err != nil {
		l.Println(err)
		return
	}

	l.Printf("Message sent successfully")
}

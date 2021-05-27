package main

import (
	"log"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func main() {
	l := log.New(os.Stdout, "golang-slack-bot ", log.LstdFlags)

	client := slack.New(os.Getenv("SLACK_BOT_AUTH_TOKEN"))
	
	// waiting for 10 seconds before sending message
	// c := make(chan string)
	// go wait(c)
	// msg := <-c
	
	channelID, timestamp, err := client.PostMessage(
		"C023ESTCXNU",
		slack.MsgOptionText("Hello from bot!", false),
	)
	if err != nil {
		l.Println(err)	
		return
	}
	l.Printf("Message sent successfully %s at %s", channelID, timestamp)
}

func wait(c chan string) {
	time.Sleep(time.Second*10)
	c <- "completed extraction successfully"
}
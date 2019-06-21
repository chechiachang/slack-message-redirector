package main

import (
	"os"
)

func main() {
	config := &Config{
		SlackToken:            os.Getenv("SLACK_TOKEN"),
		DefaultSlackChannelID: os.Getenv("SLACK_CHANNEL_ID"),
		AuthToken:             os.Getenv("AUTH_TOKEN"),
	}
	app := New(config)
	app.Run()
}

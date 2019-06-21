package main

import (
	"os"
)

func main() {
	config := &Config{
		AuthToken:      os.Getenv("AUTH_TOKEN"),
		SlackUrl:       os.Getenv("SLACK_URL"),
		SlackChannel:   os.Getenv("SLACK_CHANNEL"),
		SlackUsername:  os.Getenv("SLACK_USERNAME"),
		SlackIconEmoji: os.Getenv("SLACK_ICON_EMOJI"),
	}
	app := New(config)
	app.Run()
}

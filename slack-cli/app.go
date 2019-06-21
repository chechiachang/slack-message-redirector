package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nlopes/slack"
)

/*
golang slack API:
	https://github.com/nlopes/slack
*/

type App struct {
	Config *Config
	Api    *slack.Client
}

type Config struct {
	SlackToken            string
	DefaultSlackChannelID string
	AuthToken             string
}

func New(config *Config) *App {
	if config.AuthToken == "" {
		log.Println("Running without AuthToken")
	}
	if config.DefaultSlackChannelID == "" {
		log.Println("Running without DefaultSlackChannelID")
	}
	if config.SlackToken == "" {
		log.Println("Running without SlackToken")
	}
	return &App{
		Config: config,
		Api:    slack.New(config.SlackToken),
	}
}

func (a *App) Run() {
	log.Println("Running app with config: ", a.Config)
	router := mux.NewRouter()

	router.HandleFunc("/slack/sendgrid", SendSendgridSlackMessage(a)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func SendSendgridSlackMessage(a *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		const (
			msgUsername  = "Sendgrid"
			msgIconEmoji = ":sendgrid:"
		)

		if a.Config.AuthToken != "" && r.Header.Get("Authorization") != a.Config.AuthToken {
			log.Fatal(errors.New("Authorization token not match"))
			w.WriteHeader(403)
			return
		}

		groupID := a.Config.DefaultSlackChannelID // channel: sendgrid
		// TODO Get channel(group) information
		//group, err := api.GetGroupInfo("group")
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}
		// groupId = group.ID

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
			return
		}
		bodyString := string(bodyBytes)
		//attachment := slack.Attachment{
		//	Pretext: "Sendgrid",
		//	Text:    bodyString,
		//	// Uncomment the following part to send a field too
		//	/*
		//		Fields: []slack.AttachmentField{
		//			slack.AttachmentField{
		//				Title: "a",
		//				Value: "no",
		//			},
		//		},
		//	*/
		//}

		escape := false
		//channelID, timestamp, err := api.PostMessage(groupID, slack.MsgOptionText(bodyString, escape), slack.MsgOptionAttachments(attachment))
		channelID, timestamp, err := a.Api.PostMessage(groupID, slack.MsgOptionUsername(msgUsername), slack.MsgOptionIconEmoji(msgIconEmoji), slack.MsgOptionText(bodyString, escape))
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
			return
		}
		log.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)

		w.WriteHeader(200)
		return
	}
}

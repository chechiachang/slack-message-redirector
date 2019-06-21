package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
golang slack API:
	https://github.com/nlopes/slack
*/

const (
	payloadTemplate = "{\"channel\": \"%s\", \"username\": \"%s\", \"text\": \"%s\", \"icon_emoji\": \"%s\"}"
)

type App struct {
	Config *Config
	Client *http.Client
}

type Config struct {
	SlackUrl       string
	SlackChannel   string
	SlackUsername  string
	SlackIconEmoji string
	AuthToken      string
}

func New(config *Config) *App {
	if config.AuthToken == "" {
		log.Println("Running without AuthToken")
	}
	if len(config.SlackUrl) == 0 {
		panic("Running without Slack url")
	}
	if len(config.SlackChannel) == 0 {
		log.Println("Running without SLACK_CHANNEL. Using webhook default")
	}
	if len(config.SlackUsername) == 0 {
		log.Println("Running without SLACK_USERNAME. Using webhook default")
	}
	if len(config.SlackIconEmoji) == 0 {
		log.Println("Running without SLACK_ICON_EMOJI. Using webhook default")
	}
	return &App{
		Config: config,
		Client: &http.Client{},
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

		if a.Config.AuthToken != "" && r.Header.Get("Authorization") != a.Config.AuthToken {
			log.Fatal(errors.New("Authorization token not match"))
			w.WriteHeader(403)
			return
		}

		// Read upstream data from sendgrid
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

		payload, err := json.Marshal(Payload{
			Channel:   a.Config.SlackChannel,
			Username:  a.Config.SlackUsername,
			Text:      bodyString,
			IconEmoji: a.Config.SlackIconEmoji,
		})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
			return
		}

		req, err := http.NewRequest("POST", a.Config.SlackUrl, bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
			return
		}
		req.Header.Add("Content-Type", "application/json")

		resp, err := a.Client.Do(req)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
			return
		}

		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		log.Printf("Message successfully sent. Status: %s Response: %s", resp.Status, string(respBody))
		w.WriteHeader(200)
		return
	}
}

type Payload struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

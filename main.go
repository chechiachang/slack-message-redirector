package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nlopes/slack"
)

/*
golang slack API:
	https://github.com/nlopes/slack
*/
var api *slack.Client
var SlackChannelID string

func init() {
	// TODO use token in apps for slack instead of personal token
	token := os.Getenv("SLACK_TOKEN")
	SlackChannelID = os.Getenv("SLACK_CHANNEL_ID")
	api = slack.New(token)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/slack", SendSlackMessage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func SendSlackMessage(w http.ResponseWriter, r *http.Request) {

	groupID := SlackChannelID // channel: sendgrid
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
	channelID, timestamp, err := api.PostMessage(groupID, slack.MsgOptionText(bodyString, escape))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}
	log.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)

	w.WriteHeader(200)
	return
}

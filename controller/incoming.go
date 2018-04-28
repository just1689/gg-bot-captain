package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/just1689/gg-bot-captain/model/messages"
	log "github.com/sirupsen/logrus"
)

func HandleIncoming(b []byte) {

	go func(b []byte) {
		con, err := getConversation(b)
		if err != nil {
			return
		}
		if con == messages.ConversationShareTag {
			handleMyTagMessage(b)
			return
		} else if con == messages.ConversationListOfGames {
			handleListOfGames(b)
			return
		} else if con == messages.ConversationShareDynamicThings {
			handleDynamicThings(b)
			return
		}

		log.Infoln(fmt.Sprintf("Received: %s", string(b)))

	}(b)

}

func getConversation(b []byte) (string, error) {
	r := bytes.NewReader(b)
	decoder := json.NewDecoder(r)
	var message messages.Message
	err := decoder.Decode(&message)
	if err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return message.Conversation, err
}

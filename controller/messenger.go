package controller

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
)

func HandleIncoming(v string) {
	go func(v string) {
		con, err := getConverstation(v)
		if err != nil {
			return
		}
		if con == ConversationShareTag {
			tag, errorBuild := model.BuildTagFromString(v)
			if errorBuild != nil {
				return
			}
			handleMyTag(tag)
		}
	}(v)
}

func getConverstation(v string) (string, error) {
	decoder := json.NewDecoder(v)
	var message model.Message
	err := decoder.Decode(&message)
	if err != nil {
		fmt.Errorf("There was a problem decoding the post message: %s", err.Error())
	}
	return message.Conversation, err
}

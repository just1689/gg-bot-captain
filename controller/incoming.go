package controller

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/just1689/gg-bot-captain/model/messages"
	log "github.com/sirupsen/logrus"
)

func HandleIncoming(b []byte) {

	go func(b []byte) {

		jsonParsed, err := gabs.ParseJSON(b)

		con, _ := jsonParsed.Path("conversation").Data().(string)

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
		} else if con == "S_SHARE_HEALTH" {
			//Ignore
			return
		} else if con == "S_CHANGE_VIEW" {
			//Ignore
			return
		}

		log.Infoln(fmt.Sprintf("Received: %s", con))

	}(b)

}
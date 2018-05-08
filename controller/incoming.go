package controller

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/just1689/gg-bot-captain/model/messages"
	"github.com/just1689/gg-bot-captain/util"
	"github.com/sirupsen/logrus"
)

//HandleIncoming handles incoming messages
func HandleIncoming(b []byte) {

	go func(b []byte) {

		//Handle incoming pings
		if string(b) == "0" {
			util.SendMessage("1")
			return
		}

		jsonParsed, err := gabs.ParseJSON(b)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("Error: %e", err))
			return
		}
		con, _ := jsonParsed.Path("conversation").Data().(string)

		if messages.InConversationsToIgnore(con) {
			logrus.Debugln(fmt.Sprintf("Ignoring message about: %s", con))
			return
		}

		routeMessage(con, b)

	}(b)

}

func routeMessage(con string, b []byte) {
	if con == messages.ConversationShareTag {
		handleMyTagMessage(b)
	} else if con == messages.ConversationListOfGames {
		handleListOfGames(b)
	} else if con == messages.ConversationShareDynamicThings {
		handleDynamicThings(b)
	} else if con == messages.ConversationShareMap {
		handleMap(b)
	} else {
		logrus.Infoln(fmt.Sprintf("Received: %s", con))
	}

}

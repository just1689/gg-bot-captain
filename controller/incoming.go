package controller

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/just1689/gg-bot-captain/model/messages"
	"github.com/just1689/gg-bot-captain/util"
	log "github.com/sirupsen/logrus"
)

//HandleIncoming handles incoming messages
func HandleIncoming(b []byte) {

	go func(b []byte) {

		jsonParsed, err := gabs.ParseJSON(b)
		if err != nil {
			log.Errorln(fmt.Sprintf("Error: %e", err))
			return
		}

		con, _ := jsonParsed.Path("conversation").Data().(string)
		routeMessage(con, b)

	}(b)

}

func routeMessage(con string, b []byte) {
	if con == messages.ConversationShareTag {
		handleMyTagMessage(b)
		return
	} else if con == messages.ConversationListOfGames {
		handleListOfGames(b)
		return
	} else if con == messages.ConversationShareDynamicThings {
		handleDynamicThings(b)
		return
	} else if con == messages.ConversationShareMap {
		handleMap(b)
		return
	} else if con == "S_SHARE_HEALTH" {
		//Ignore
		return
	} else if con == "S_CHANGE_VIEW" {
		//Ignore
		return
	} else if con == "S_SCOREBOARD" {
		//Ignore
		return
	} else if con == "S_P_LEVEL" {
		//Ignore
		return
	} else if con == "S_PLAY_SOUND" {
		//Ignore
		return
	} else if con == "S_SHARE_BULLETS" {
		//Ignore
		return
	} else if con == "S_SHARE_SPRAY" {
		//Ignore
		return
	} else if con == "S_ORB_N" {
		//Ignore
		return
	} else if con == "S_PLAYER_LEFT" {
		//Ignore
		return

	} else if con == "" && string(b) == "0" {
		util.SendMessage("1")
		return
	}

	log.Infoln(fmt.Sprintf("Received: %s", con))
}

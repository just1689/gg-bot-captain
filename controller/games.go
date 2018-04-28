package controller

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
)

func handleListOfGames(b []byte) {
	list, errorBuild := incoming.BuildListOfGamesFromString(b)
	if errorBuild != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
		return
	}
	if len(list.Games) == 1 {
		game := list.Games[0]
		SendJoinGameMessage(game)
		return
	}

	log.Errorln(fmt.Sprintf("Was expecting one game. Found: %v", len(list.Games)))

}

package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	"github.com/sirupsen/logrus"
)

func handleListOfGames(b []byte) {
	list, errorBuild := incoming.BuildListOfGamesFromString(b)
	if errorBuild != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
		return
	}
	if len(list.Games) == 1 {
		game := list.Games[0]
		SendJoinGameMessage(game)
		return
	}
	logrus.Errorln(fmt.Sprintf("Was expecting one game. Found: %v", len(list.Games)))

}

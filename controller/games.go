package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	"github.com/sirupsen/logrus"
)

func handleListOfGames(b []byte) chan model.Game {
	c := make(chan model.Game)

	go func() {
		list, errorBuild := incoming.BuildListOfGamesFromString(b)
		if errorBuild != nil {
			logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
			close(c)
			return
		}
		if len(list.Games) == 1 {
			game := list.Games[0]
			c <- game
			close(c)
			return
		}
		logrus.Errorln(fmt.Sprintf("Was expecting one game. Found: %v", len(list.Games)))
		close(c)
	}()

	return c

}

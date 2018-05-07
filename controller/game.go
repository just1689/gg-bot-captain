package controller

import (
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/outgoing"
	"github.com/just1689/gg-bot-captain/util"
)

//SendJoinGameMessage handles incoming messages
func SendJoinGameMessage(game model.Game) {
	message := outgoing.BuildMessageJoinGame(game.ID)

	util.SendMessage(message)
}

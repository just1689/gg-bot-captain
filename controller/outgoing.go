package controller

import (
	"github.com/just1689/gg-bot-captain/model/messages/outgoing"
	"github.com/just1689/gg-bot-captain/util"
)

//SendMessageJoinServer tells the server about the bot
func SendMessageJoinServer() {
	item := outgoing.BuildMessageRequestJoinServer("Botty McBotface")
	util.SendMessage(item)
}

package controller

import (
	"github.com/just1689/gg-bot-captain/util"
	"github.com/just1689/gg-bot-captain/model/messages/outgoing"
)

func SendMessageJoinServer() {
	item := outgoing.BuildMessageRequestJoinServer("Captain")
	util.SendMessage(item)
}

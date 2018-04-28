package controller

import (
	"github.com/just1689/gg-bot-captain/model/messages/outgoing"
	"github.com/just1689/gg-bot-captain/util"
)

func SendMessageJoinServer() {
	item := outgoing.BuildMessageRequestJoinServer("Captain")
	util.SendMessage(item)
}

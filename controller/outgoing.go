package controller

import (
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/util"
)

func SendMessageJoinServer() {
	item := model.BuildMessageRequestJoinServer("Captain")
	util.SendMessage(item)
}

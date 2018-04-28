package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	log "github.com/sirupsen/logrus"
)

func handleMap(b []byte) {

	//log.Infoln(fmt.Sprintf("Map: %s", string(b)))

	mapMessage, errorBuild := incoming.BuildMessageMapFromString(b)
	if errorBuild != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
		return
	}
	assignTileID(mapMessage.Map)
	pushMap(mapMessage.Map)
}

func assignTileID(list []model.Tile) {
	for _, item := range list {
		item.ID = fmt.Sprintf("%v.%v.%v", item.Point.X, item.Point.Y, item.Point.Z)
	}
}

func pushMap(list []model.Tile) {
	go func(list []model.Tile) {
		for _, item := range list {
			mem.Push(model.TableNameTile, item)
		}
	}(list)
}

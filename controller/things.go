package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	log "github.com/sirupsen/logrus"
)

func handleDynamicThings(b []byte) {
	things, errorBuild := incoming.BuildMessageShareDynamicThingsFromString(b)
	if errorBuild != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
		return
	}
	list := things.Things
	pushThingsAsync(list)
}

func pushThingsAsync(things []model.Thing) {
	go func(things []model.Thing) {
		for _, thing := range things {
			mem.Push(model.TableNameThing, thing)
		}
	}(things)
}

package ai

import (
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
)

func pursueThingAndAttack(thing model.Thing) {
	myThing, err := mem.GetAllByFieldAndUintValue(model.TableNameThing, "id", controller.GetMyTag())
	if err != nil {
		return
	}

}

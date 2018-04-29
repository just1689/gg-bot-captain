package ai

import (
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/controller"
)

//PickTarget for now picks the first thing that is not me
func PickTarget() model.Thing {
	myTag := controller.GetMyTag()
	things := controller.GetAllThings()
	for _, thing := range things {
		if thing.Tag == myTag {
			continue
		}
		return thing
	}
	return model.Thing{Tag: -1}
}

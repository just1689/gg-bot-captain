package ai

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/model"
	log "github.com/sirupsen/logrus"
)

//PickTarget for now picks the first thing that is not me
func PickTarget() (model.Thing, bool) {
	myTag := controller.GetMyTag()
	log.Infoln(fmt.Sprintf("My tag is: %v", myTag))

	things := controller.GetAllThings()
	log.Infoln(fmt.Sprintf("Found things: %v", len(things)))
	for _, thing := range things {
		if thing.Tag == myTag {
			continue
		}
		return thing, true
	}
	log.Infoln(fmt.Sprintf("Could not find target: %v", len(things)))
	return model.Thing{Tag: -1}, false
}

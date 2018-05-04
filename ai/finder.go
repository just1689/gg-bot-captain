package ai

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/model"
	log "github.com/sirupsen/logrus"
)

func pickTarget() (model.Thing, bool) {
	myTag := controller.GetMyTag()
	log.Infoln(fmt.Sprintf("My tag is: %v", myTag))

	things := controller.GetAllThings()
	log.Debugln(fmt.Sprintf("Found things: %v", len(things)))
	for _, thing := range things {
		if thing.Tag == myTag {
			continue
		}
		log.Infoln(fmt.Sprintf("Found a thing to target"))
		return thing, true
	}
	log.Infoln(fmt.Sprintf("Could not find target. Only found %v items", len(things)))
	return model.Thing{}, false
}

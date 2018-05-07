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

	c := controller.GetAllThings()
	count := 0
	for thing := range c {
		count++
		if thing.Tag == myTag {
			continue
		}
		log.Infoln(fmt.Sprintf("Found a thing to target"))
		return thing, true
	}
	log.Infoln(fmt.Sprintf("Could not find target. Only found %v items", count))
	return model.Thing{}, false
}

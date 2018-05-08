package ai

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
)

//PickTarget looks for another tank to shoot
func PickTarget() (model.Thing, bool) {
	myTag := controller.GetMyTag()
	logrus.Infoln(fmt.Sprintf("My tag is: %v", myTag))

	c := controller.GetAllThings()
	count := 0
	for thing := range c {
		count++
		if thing.Tag == myTag {
			continue
		}
		logrus.Infoln(fmt.Sprintf("Found a thing to target"))
		return thing, true
	}
	logrus.Infoln(fmt.Sprintf("Could not find target. Only found %v items", count))
	return model.Thing{}, false
}

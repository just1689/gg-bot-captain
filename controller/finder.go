package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
)

//PickTarget looks for another tank to shoot
func PickTarget() (model.Thing, bool) {
	myTag := GetMyTag()

	c := GetAllThings()
	count := 0
	for thing := range c {
		count++
		if thing.Tag != myTag {
			logrus.Infoln(fmt.Sprintf("Found a thing to target"))
			return thing, true
		}
	}
	return model.Thing{}, false
}

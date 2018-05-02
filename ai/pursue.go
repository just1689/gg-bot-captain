package ai

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
	"time"
)

func pursueThingAndAttack(target model.Thing) {
	iter, err := mem.GetAllByFieldAndUintValue(model.TableNameThing, "id", controller.GetMyTag())
	if err != nil {
		logrus.Error(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	me, ok := model.IteratorToFirstThing(iter, err)
	if !ok {
		logrus.Error(fmt.Sprintf("Could not find me!"))
		return
	}

	start := time.Now()
	for {

		if wouldIHit(me, target) {
			shoot()
		}

		duration := time.Since(start)
		if duration.Seconds() > 10 {
			break
		}
	}

}

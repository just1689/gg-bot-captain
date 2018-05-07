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
	iterator, err := mem.GetAllByFieldAndUintValue(model.TableNameThing, "id", controller.GetMyTag())
	if err != nil {
		logrus.Error(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	me, ok := model.IteratorToFirstThing(iterator, err)
	if !ok {
		logrus.Error(fmt.Sprintf("Could not find me!"))
		return
	}

	start := time.Now()
	for {

		logrus.Infoln(fmt.Sprintf("Checking if I would hit: %v", target.Tag))
		if wouldIHit(me, target) {
			shoot()
		}

		duration := time.Since(start)
		if duration.Seconds() > 10 {
			logrus.Infoln(fmt.Sprintf("Timing out pursuit"))
			break
		}

		time.Sleep(5 * time.Second)
	}

}

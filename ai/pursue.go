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

		if wouldIHit(me, target) {
			shoot()

		} else {
			//Rotate
			direction := GetRotationBetween(me, target)
			if direction == model.DirectionClockwise {
				controller.PressRight()
			} else if direction == model.DirectionAntiClockwise {
				controller.PressLeft()
			}
		}

		duration := time.Since(start)
		if duration.Seconds() > 10 {
			logrus.Infoln(fmt.Sprintf("Timing out pursuit"))
			break
		}

		//Sleep 100 ms
		//time.Sleep(time.Duration(100))
		time.Sleep(100 * time.Millisecond)
	}

}

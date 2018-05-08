package ai

import (
	"github.com/sirupsen/logrus"
	"time"
)

//Schedule handles the running of the AI go routine
func Schedule() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			logrus.Println("ai.Scheduler run...")
			targetThing, ok := PickTarget()
			if !ok {
				logrus.Infoln("Could not find a target. Going to sleep")
				time.Sleep(5 * time.Second)
				continue
			}
			pursueThingAndAttack(targetThing)
		}
	}()
}

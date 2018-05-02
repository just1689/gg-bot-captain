package ai

import (
	log "github.com/sirupsen/logrus"
	"time"
)

//Schedule handles the running of the AI go routine
func Schedule() {
	go func() {
		for {
			log.Println("ai.Schedule run...")
			targetThing, ok := pickTarget()
			if !ok {
				log.Infoln("Could not find a target. Going to sleep")
				time.Sleep(5 * time.Second)
				continue
			}
			pursueThingAndAttack(targetThing)
		}
	}()
}

package ai

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

//Schedule handles the running of the AI go routine
func Schedule() {
	go func() {
		for {
			targetThing, ok := PickTarget()
			if !ok {
				log.Infoln(fmt.Sprintf("Could not find a target. Going to sleep"))
				time.Sleep(5 * time.Second)
				continue
			}
			pursueThingAndAttack(targetThing)
		}
	}()
}

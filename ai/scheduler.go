package ai

import "time"

//Schedule handles the running of the AI go routine
func Schedule() {
	go func() {
		for {
			targetThing := PickTarget()
			if targetThing.Tag == -1 {
				time.Sleep(10 * time.Second)
				continue
			}
			pursueThing(targetThing)
		}
	}()
}

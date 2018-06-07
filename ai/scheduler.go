package ai

import (
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/sirupsen/logrus"
	"time"
)

//Schedule handles the running of the AI go routine
func Schedule() (ticker chan bool) {
	ticker = make(chan bool)
	go func() {
		for {
			logrus.Println("ai.Scheduler run...")
			myPersonality := personality.ChoosePersonality()
			ticker <- true
			myGoal := personality.ChooseGoal(myPersonality)
			ticker <- true
			plan(myGoal, myPersonality)
			ticker <- true
			time.Sleep(1 * time.Second)
		}
	}()
	return ticker
}

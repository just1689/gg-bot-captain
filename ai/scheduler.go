package ai

import (
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/just1689/gg-bot-captain/ai/planner"
	"github.com/just1689/gg-bot-captain/controller"
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
			actions, ok := planner.Plan(myGoal, myPersonality)
			if !ok {
				time.Sleep(1 * time.Second)
				continue
			}
			controller.Act(actions)
			ticker <- true
		}
	}()
	return ticker
}

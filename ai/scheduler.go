package ai

import (
	"github.com/sirupsen/logrus"
	"github.com/just1689/gg-bot-captain/ai/personality"
)

//Schedule handles the running of the AI go routine
func Schedule() {

	go func() {
		for {
			logrus.Println("ai.Scheduler run...")
			myPersonality := personality.ChoosePersonality()
			myGoal := personality.ChooseGoal(myPersonality)
			plan(myGoal, myPersonality)
		}
	}()
	
}

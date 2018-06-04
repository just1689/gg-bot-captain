package ai

import (
	"github.com/sirupsen/logrus"
	"github.com/just1689/gg-bot-captain/ai/goal"
)

//Schedule handles the running of the AI go routine
func Schedule() {
	go func() {
		for {

			logrus.Println("ai.Scheduler run...")

			myPersonality := ChoosePersonality()

			myGoal := ChooseGoal(myPersonality)

			if myGoal.Goal == goal.Kill {
				//Pursue
			} else if myGoal.Goal == goal.Hide {
				//Hide
			}

		}
	}()
}

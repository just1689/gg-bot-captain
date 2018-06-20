package ai

import (
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/just1689/gg-bot-captain/ai/planner"
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
	"time"
)

//Schedule handles the running of the AI go routine
func Schedule() (ticker chan bool) {
	ticker = make(chan bool)
	go func() {

		logrus.Println("ai.Scheduler run...")
		myPersonality := personality.ChoosePersonality()
		ticker <- true

		for {
			scheduleTick(myPersonality, ticker)
		}
	}()
	return ticker
}

func scheduleTick(myPersonality model.Personality, ticker chan bool) {
	myGoal := personality.ChooseGoal(myPersonality)
	ticker <- true
	myActions, ok := planner.Plan(myPersonality, myGoal)
	if !ok {
		time.Sleep(1 * time.Second)
		return
	}
	controller.Act(myPersonality, myGoal, myActions)
	ticker <- true

}

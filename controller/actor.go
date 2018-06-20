package controller

import (
	"github.com/just1689/gg-bot-captain/model"
	"log"
)

func Act(myPersonality model.Personality, goal model.Goal, actions []model.Action) (abortPersonality bool, abortGoal bool) {

	for _, action := range actions {

		if myPersonality.IsNowIrrelevant() {
			log.Println("My personality doesn't fit. Aborting")
			abortPersonality = true
			return
		}

		if goal.IsSatisfied() {
			log.Println("The goal has been satisfied. Aborting")
			abortGoal = true
			return
		}

		if action.Action == model.ActionMoveNear {
			autoMoveNear(action.TagTarget)
		} else if action.Action == model.ActionAim {
			autoAim(action.TagTarget)
		} else if action.Action == model.ActionAttack {
			autoAttack(action.TagTarget)
		} else if action.Action == model.ActionRunAway {
			autoRun()
		}

	}
	return
}

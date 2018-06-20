package controller

import (
	"github.com/just1689/gg-bot-captain/model"
	"log"
)

func Act(myPersonality model.Personality, goal model.Goal, actions []model.Action) {
	for _, action := range actions {
		if goal.IsSatisfied() {
			log.Println("The goal has been satisfied. Aborting")
			return
		}

		if action.Action == model.ActionMoveNear {

		}

	}
}

package controller

import (
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/just1689/gg-bot-captain/ai/goal"
)

/*
	Goal reached -> Abort Goal

 */
func Act(myPersonality personality.Personality, goal goal.Goal, actions []model.Action) {
	for _, action := range actions {

	}
}

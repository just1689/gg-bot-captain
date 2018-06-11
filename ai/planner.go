package ai

import (
	"github.com/just1689/gg-bot-captain/ai/goal"
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/just1689/gg-bot-captain/model"
)

func plan(myGoal goal.Goal, myPersonality personality.Personality) (actions []model.Action, ok bool) {
	ok = false
	if myGoal.Goal == goal.Kill {
		//Pursue
		ok = true
	} else if myGoal.Goal == goal.Hide {
		//Hide
		ok = true
	}
	return actions, ok

}

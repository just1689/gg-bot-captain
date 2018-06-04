package ai

import (
	"github.com/just1689/gg-bot-captain/ai/goal"
	"github.com/just1689/gg-bot-captain/ai/personality"
)

func plan(myGoal goal.Goal, myPersonality personality.Personality) {
	if myGoal.Goal == goal.Kill {
		//Pursue
	} else if myGoal.Goal == goal.Hide {
		//Hide
	}

}

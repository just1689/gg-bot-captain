package planner

import (
	"github.com/just1689/gg-bot-captain/ai/goal"
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/just1689/gg-bot-captain/model"
)

func Plan(myGoal goal.Goal, myPersonality personality.Personality) (actions []model.Action, ok bool) {
	ok = false
	if myGoal.Goal == goal.Kill {

		actionMove := model.Action{Action: model.ActionMoveNear, TagTarget: myGoal.Tag}
		actions = append(actions, actionMove)

		actionAim := model.Action{Action: model.ActionAim, TagTarget: myGoal.Tag}
		actions = append(actions, actionAim)

		actionAttack := model.Action{Action: model.ActionAttack, TagTarget: myGoal.Tag}
		actions = append(actions, actionAttack)

		//Pursue
		ok = true
	} else if myGoal.Goal == goal.Hide {
		//Hide
		ok = true
	}
	return

}

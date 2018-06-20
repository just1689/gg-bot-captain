package planner

import (
	"github.com/just1689/gg-bot-captain/model"
)

func Plan(myPersonality model.Personality, myGoal model.Goal) (actions []model.Action, ok bool) {
	ok = false
	if myGoal.Goal == model.Kill {

		actionMove := model.Action{Action: model.ActionMoveNear, TagTarget: myGoal.Tag}
		actions = append(actions, actionMove)

		actionAim := model.Action{Action: model.ActionAim, TagTarget: myGoal.Tag}
		actions = append(actions, actionAim)

		actionAttack := model.Action{Action: model.ActionAttack, TagTarget: myGoal.Tag}
		actions = append(actions, actionAttack)

		//Pursue
		ok = true
	} else if myGoal.Goal == model.Hide {

		actionAttack := model.Action{Action: model.ActionRunAway}
		actions = append(actions, actionAttack)

		//Hide
		ok = true
	}
	return

}

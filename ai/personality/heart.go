package personality

import (
	"github.com/just1689/gg-bot-captain/ai/goal"
	"github.com/just1689/gg-bot-captain/controller"
)

func ChoosePersonality() Personality {

	//TODO: better impl
	//Consider conditions?

	return RandomPersonality()
}

func ChooseGoal(personality Personality) goal.Goal {
	if personality == Hunter {
		target, found := controller.PickTarget()
		if found {
			return goal.NewKillGoal(target.Tag)
		}
	}
	if personality == Wimp {
		return goal.NewHideGoal()

	}

	//If my personality does not help me choose a goal, then hide is goal
	return goal.NewHideGoal()

}

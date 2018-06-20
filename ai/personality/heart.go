package personality

import (
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/just1689/gg-bot-captain/model"
)

func ChoosePersonality() model.Personality {

	//TODO: better impl
	//Consider conditions?

	return model.RandomPersonality()
}

func ChooseGoal(personality model.Personality) model.Goal {
	if personality == model.Hunter {
		target, found := controller.PickTarget()
		if found {
			return model.NewKillGoal(target.Tag)
		}
	}
	if personality == model.Wimp {
		return model.NewHideGoal()

	}

	//If my personality does not help me choose a goal, then hide is goal
	return model.NewHideGoal()

}

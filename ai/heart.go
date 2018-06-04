package ai

import "github.com/just1689/gg-bot-captain/ai/goal"

func ChoosePersonality() Personality {

	//TODO: better impl
	//Consider conditions?

	return randomPersonality()
}

func ChooseGoal(personality Personality) goal.Goal {
	if personality == Hunter {
		target, found := PickTarget()
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

package ai

func ChoosePersonality() Personality {

	//TODO: better impl
	//Consider conditions?

	return randomPersonality()
}

func ChooseGoal(personality Personality) Goal {
	if personality == Hunter {
		target, found := PickTarget()
		if found {
			return newKillGoal(target)
		}
	}
	if personality == Wimp {
		return newHideGoal()

	}

	//If my personality does not help me choose a goal, then hide is goal
	return newHideGoal()

}

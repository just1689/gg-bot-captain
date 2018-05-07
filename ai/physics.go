package ai

import "github.com/just1689/gg-bot-captain/model"

func wouldIHit(me model.Thing, them model.Thing) bool {
	//Todo: Impl
	return false
}

//GetRotationBetween finds out which way to rotate to meet another tank
func GetRotationBetween(me model.Thing, them model.Thing) int {
	//Todo: find quickest route between two tanks
	return model.DirectionClockwise
}

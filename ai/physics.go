package ai

import (
	"github.com/just1689/gg-bot-captain/model"
	"math/rand"
)

const maxRand = 1000
const midRand = maxRand/2 - 1

func wouldIHit(me model.Thing, them model.Thing) bool {

	r := rand.Intn(maxRand)
	if r > midRand {
		return true
	}

	//Todo: Impl.
	//For now, just be random
	return false
}

//GetRotationBetween finds out which way to rotate to meet another tank
func GetRotationBetween(me model.Thing, them model.Thing) int {
	//Todo: find quickest route between two tanks
	return model.DirectionClockwise
}

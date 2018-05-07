package ai

import (
	"github.com/just1689/gg-bot-captain/model"
	"math/rand"
)

const maxRand = 1000
const midRand = maxRand/2 - 1

func wouldIHit(me model.Thing, them model.Thing) bool {

	rot := GetRotationBetween(me, them)

	r := rand.Intn(maxRand * rot)
	if r > midRand {
		return true
	}

	//Todo: Impl.
	//For now, just be random
	return false
}

//GetRotationBetween finds out which way to rotate to meet another tank
func GetRotationBetween(me model.Thing, them model.Thing) int {

	viewer := me.Point.PointToSpacialPoint()
	target := them.Point.PointToSpacialPoint()
	gradient := GetGradient(viewer, target)

	//Impossible
	if gradient == 2 {
		return 0
	}

	//Todo: find quickest route between two tanks
	return model.DirectionClockwise
}

func IsAimingAt(viewer model.SpacialPoint, target model.SpacialPoint) bool {
	gradient := GetGradient(viewer, target)
	return Aligns(viewer, target, gradient)
}

func Aligns(viewer model.SpacialPoint, target model.SpacialPoint, gradient int) bool {
	return true
}

func GetGradient(viewer model.SpacialPoint, target model.SpacialPoint) int {
	deltaX := viewer.X - target.X
	deltaZ := viewer.Z - target.Z
	gradient := deltaZ / deltaX
	return int(gradient)
}

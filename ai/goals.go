package ai

import "github.com/just1689/gg-bot-captain/model"

type GoalLabel int

const (
	Kill GoalLabel = iota
	Hide
)

type Goal struct {
	Goal GoalLabel
	Tag  uint
}

func newHideGoal() (goal Goal) {
	return Goal{Goal: Hide}
}
func newKillGoal(thing model.Thing) Goal {
	item := Goal{Goal: Kill, Tag: thing.Tag}
	return item
}

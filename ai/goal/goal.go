package goal

type GoalLabel int

const (
	Kill GoalLabel = iota
	Hide
)

type Goal struct {
	Goal GoalLabel
	Tag  uint
}

func NewHideGoal() (goal Goal) {
	return Goal{Goal: Hide}
}
func NewKillGoal(tag uint) Goal {
	item := Goal{Goal: Kill, Tag: tag}
	return item
}

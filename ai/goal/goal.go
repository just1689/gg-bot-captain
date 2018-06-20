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

func (g Goal) IsSatisfied() bool {
	if g.Goal == Kill {
		return g.isSatisfiedKill()
	} else if g.Goal == Hide {
		return g.isSatisfiedHide()
	}

	/*
		Goals with no satisfaction criteria risk implied yet unintended endless execution
		This means it would be better for it to consider it satisfied if does not explicitly
		implement criteria to measure that.

	 */
	return true
}

func (g Goal) isSatisfiedKill() bool {
	//Todo: time-series implementation for capturing state related to deaths
	return false
}

func (g Goal) isSatisfiedHide() bool {
	//Todo: is hiding ever complete?
	return false
}

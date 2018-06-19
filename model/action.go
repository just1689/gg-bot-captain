package model

type ActionLabel int

const (
	ActionAim      ActionLabel = iota
	ActionMoveNear
	ActionAttack
)

type Action struct {
	Action    ActionLabel
	TagTarget uint
}

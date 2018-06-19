package model

type ActionLabel int

const (
	ActionAim ActionLabel = iota
	ActionMoveNear
	ActionAttack
	ActionRunAway
)

type Action struct {
	Action    ActionLabel
	TagTarget uint
}

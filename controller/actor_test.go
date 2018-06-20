package controller

import (
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAct(t *testing.T) {
	setupTwoPlayers()

	myPersonality := model.Hunter
	myGoal := model.NewKillGoal(2)

	var myActions = make([]model.Action, 1)
	myActions = append(myActions, model.Action{Action: model.ActionAttack, TagTarget: 2})
	abortPersonality, abortGoal := Act(myPersonality, myGoal, myActions)
	assert.False(t, abortPersonality, "A two player game should not abort personality Hunter")
	assert.False(t, abortGoal, "A two player game should not abort personality Hunter")

}

func setupTwoPlayers() {
	//Setup the game
	mem.Init()

	//Add my player
	myTag := uint(1)
	me := model.Tag{Tag: myTag}
	mem.Push(model.TableNameTag, me)
	myTank := model.Thing{Tag: me.Tag}
	mem.Push(model.TableNameThing, myTank)

	//Add a player
	theirTag := uint(2)
	theirTank := model.Thing{Tag: theirTag}
	mem.Push(model.TableNameThing, theirTank)

}

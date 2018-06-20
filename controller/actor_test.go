package controller

import (
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/just1689/gg-bot-captain/ai/planner"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAct(t *testing.T) {

	setupTwoPlayers()

	//Test the hunters actions
	myPersonality := model.Hunter
	myGoal := personality.ChooseGoal(myPersonality)
	myActions, ok := planner.Plan(myPersonality, myGoal)
	assert.True(t, ok, "The planner should give the ok for TestAct")
	Act(myPersonality, myGoal, myActions)

	myPersonality = model.Wimp
	myGoal = personality.ChooseGoal(myPersonality)
	myActions, ok = planner.Plan(myPersonality, myGoal)
	assert.True(t, ok, "The planner should give the ok for TestAct")
	Act(myPersonality, myGoal, myActions)

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

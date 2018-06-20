package planner

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/ai/personality"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlan(t *testing.T) {

	name := "ai.personality.TestPlan"

	//Setup the game
	mem.Init()

	setupTwoPlayers()

	myPersonality := model.Hunter
	myGoal := personality.ChooseGoal(myPersonality)
	_, ok := Plan(myPersonality, myGoal)

	assert.True(t, ok, "The planner should return ok is true")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

}

func TestPlanActions(t *testing.T) {

	name := "ai.personality.TestPlanActions"

	setupTwoPlayers()

	myPersonality := model.Hunter
	myGoal := personality.ChooseGoal(myPersonality)
	actions, _ := Plan(myPersonality, myGoal)

	assert.NotEmpty(t, actions, "The list of actions planned should not be empty")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

}

func TestPlanActionsForWimp(t *testing.T) {

	name := "ai.personality.TestPlanActionsForWimp"

	setupTwoPlayers()

	myPersonality := model.Wimp
	myGoal := personality.ChooseGoal(myPersonality)
	actions, _ := Plan(myPersonality, myGoal)

	assert.NotEmpty(t, actions, "The list of actions for a Wimp planned should not be empty")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

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

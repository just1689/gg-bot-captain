package ai

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/just1689/gg-bot-captain/ai/personality"
)

func TestPlan(t *testing.T) {

	name := "ai.personality.TestChoosePersonality"

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

	myPersonality := personality.Hunter
	myGoal := personality.ChooseGoal(myPersonality)
	_, ok := plan(myGoal, myPersonality)

	assert.True(t, ok, "The planner should return ok is true")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

}

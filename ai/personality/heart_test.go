package personality

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/ai/goal"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChoosePersonality(t *testing.T) {
	name := "ai.personality.TestChoosePersonality"

	myPersonality := ChoosePersonality()
	assert.Equal(t, myPersonality, Hunter, "My personality should be hunter")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}
}

func TestChooseGoal(t *testing.T) {
	name := "ai.personality.TestChoosePersonality"

	//Setup the game
	mem.Init()
	myTag := uint(1)
	me := model.Tag{Tag: myTag}
	mem.Push(model.TableNameTag, me)
	myTank := model.Thing{Tag: me.Tag}
	mem.Push(model.TableNameThing, myTank)

	myPersonality := Hunter
	myGoal := ChooseGoal(myPersonality)
	assert.Equal(t, myGoal.Goal, goal.Hide, "A hunter should hide when there is no other player")

	//Add a player
	theirTag := uint(2)
	theirTank := model.Thing{Tag: theirTag}
	mem.Push(model.TableNameThing, theirTank)

	myGoal = ChooseGoal(myPersonality)
	assert.Equal(t, myGoal.Goal, goal.Kill, "A hunter should kill if there is another player")
	assert.Equal(t, myGoal.Tag, theirTag, "A hunter should kill the only other player")

	myPersonality = Wimp
	myGoal = ChooseGoal(myPersonality)
	assert.Equal(t, myGoal.Goal, goal.Hide, "A wimp should hide")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}
}

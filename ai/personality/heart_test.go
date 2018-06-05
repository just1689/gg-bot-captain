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

const NAME = "ai.personality"

func TestChooseGoal(t *testing.T) {

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

	assert.Equal(t, 1, 2, "Well this does not make any sense at all")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", NAME))
	}
}

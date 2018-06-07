package ai

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {

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

	ticker := Schedule()
	timeouts := 0
	ticks := 0
	cont := true
	for cont {
		select {
		case <-ticker:
			ticks++
		case <-time.After(2 * time.Second):
			timeouts++
		}
		if timeouts >= 2 || ticks >= 6 {
			cont = false
		}
	}
	assert.True(t, ticks >= 6, "There should be at least 6 ticks")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

}

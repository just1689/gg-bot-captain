package ai_test

import (
	"testing"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/ai"
	"github.com/stretchr/testify/assert"
	log "github.com/sirupsen/logrus"
	"fmt"
)

const NAME = "ai.PickTarget"

func TestPickTarget(t *testing.T) {

	//Setup the game
	mem.Init()
	myTag := uint(1)
	theirTag := uint(2)
	me := model.Tag{Tag: myTag}
	mem.Push(model.TableNameTag, me)
	myTank := model.Thing{Tag: me.Tag}
	mem.Push(model.TableNameThing, myTank)
	theirTank := model.Thing{Tag: theirTag}
	mem.Push(model.TableNameThing, theirTank)

	//The call
	target, ok := ai.PickTarget()

	//The test
	assert.Equal(t, ok, true, "Was not able to find a target. Should have been able to.")
	assert.NotEqual(t, myTag, target.Tag, "The target was me. It should not be me.")
	assert.Equal(t, theirTag, target.Tag, "The target was not them. It should be them.")

	if t.Failed() {
		log.Println(fmt.Sprintf("Testing %s failed ❌ ", NAME))
	}
}

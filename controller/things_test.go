package controller

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersistThingsMessage(t *testing.T) {

	name := "controller.things_test.go"

	mem.Init()
	myTag := uint(1)
	me := model.Tag{Tag: myTag}
	mem.Push(model.TableNameTag, me)
	myTank := model.Thing{Tag: me.Tag}
	msg := incoming.MessageShareDynamicThings{}
	things := []model.Thing{myTank}
	msg.Things = things
	bytes, _ := json.Marshal(myTank)
	signal := PersistThingsMessage(bytes)

	<-signal
	c := GetAllThings()

	for item := range c {
		assert.Equal(t, item.Tag, myTag, "The tag of the tank in the database should be the tag I inserted")
	}

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

}

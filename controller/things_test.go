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

	name := "controller.TestPersistThingsMessage"

	mem.Init()
	myTag := uint(1)
	insertTankBlocking(myTag)
	c := GetAllThings()

	for item := range c {
		assert.Equal(t, item.Tag, myTag, "The tag of the tank in the database should be the tag I inserted")
	}

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

}

func TestGetAllThings(t *testing.T) {

	name := "controller.TestGetAllThings"
	var tag uint
	var max uint
	var min uint

	mem.Init()
	min = 100
	max = 200
	totalInserted := 0
	for tag = min; tag <= max; tag++ {
		insertTankBlocking(tag)
		totalInserted++
	}

	c := GetAllThings()

	total := 0
	for item := range c {
		assert.True(t, item.Tag >= min || item.Tag <= max, "The tag should be one of those created")
		total++
	}
	assert.Equal(t, totalInserted, total, "The number inserted should be equal to the number read")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}

}

func insertTankBlocking(tag uint) {
	myTag := uint(tag)
	me := model.Tag{Tag: myTag}
	mem.Push(model.TableNameTag, me)
	myTank := model.Thing{Tag: me.Tag}
	msg := incoming.MessageShareDynamicThings{Things: []model.Thing{myTank}}
	bytes, _ := json.Marshal(msg)
	signal := PersistThingsMessage(bytes)
	<-signal

}

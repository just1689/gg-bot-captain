package controller

import (
	"encoding/json"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestHandleMyTagMessage(t *testing.T) {

	mem.Init()

	myTag := uint(10)
	tag := model.Tag{Tag: myTag}
	bytes, _ := json.Marshal(tag)
	HandleMyTagMessage(bytes)
	assert.Equal(t, myTag, GetMyTag(), "The tag set by HandleMyTagMessage should equal what we sent")

}

package controller

import (
	"testing"
	"github.com/just1689/gg-bot-captain/model"
	"encoding/json"
	"github.com/magiconair/properties/assert"
)

func TestHandleMyTagMessage(t *testing.T) {

	myTag := uint(10)
	tag := model.Tag{Tag: myTag}
	bytes, _ := json.Marshal(tag)
	c := HandleMyTagMessage(bytes)
	<-c
	assert.Equal(t, myTag, GetMyTag(), "The tag set by HandleMyTagMessage should equal what we sent")

}

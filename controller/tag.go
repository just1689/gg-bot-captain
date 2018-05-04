package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	log "github.com/sirupsen/logrus"
)

func handleMyTagMessage(b []byte) {
	tag, errorBuild := model.BuildTagFromString(b)
	if errorBuild != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
		return
	}
	mem.Push(model.TableNameTag, tag)
	SendMessageJoinServer()
	log.Infoln(fmt.Sprintf("My tag in the database is: %v", GetMyTag()))

}

//GetMyTag finds the users tag in the database
func GetMyTag() uint {
	ite, err := mem.GetAll(model.TableNameTag)
	c, e := model.IteratorToManyTags(ite, err)
	if e != nil {
		log.Errorln(fmt.Sprintf("Error GetMyTag: %v", e.Error()))
		return 0
	}
	for item := range c {
		return item.Tag
	}
	return 0

}

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
	if err != nil {
		log.Errorln(fmt.Sprintf("Error GetMyTag: %v", err.Error()))
		return 0
	}
	tag := model.IteratorToTag(ite, err)
	return tag.Tag

}

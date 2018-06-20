package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	"github.com/sirupsen/logrus"
)

//PersistThingsMessage puts a the things found in a []byte into the database, signalling when done
func PersistThingsMessage(b []byte) {
	msg, errorBuild := incoming.BuildMessageShareDynamicThingsFromString(b)
	if errorBuild != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
		return
	}
	for _, thing := range msg.Things {
		mem.Push(model.TableNameThing, thing)
	}

}

//GetAllThings gets all things in database
func GetAllThings() chan model.Thing {
	iterator, err := mem.GetAll(model.TableNameThing)
	i := model.IteratorToManyThings(iterator, err)
	return i

}

func GetMeAndTarget(meTag uint, targetTag uint) (me model.Thing, target model.Thing) {
	c := GetAllThings()
	for item := range c {
		if item.Tag == meTag {
			me = item
		}
		if item.Tag == targetTag {
			target = item
		}
	}
	return
}

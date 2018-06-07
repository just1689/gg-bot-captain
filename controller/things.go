package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	"github.com/sirupsen/logrus"
)

//PersistThingsMessage puts a the things found in a []byte into the database, signalling when done
func PersistThingsMessage(b []byte) chan bool {
	signal := make(chan bool)
	go func() {
		msg, errorBuild := incoming.BuildMessageShareDynamicThingsFromString(b)
		if errorBuild != nil {
			logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
			signal <- true
			return
		}
		for _, thing := range msg.Things {
			mem.Push(model.TableNameThing, thing)
		}
		signal <- true
	}()
	return signal

}

//GetAllThings gets all things in database
func GetAllThings() chan model.Thing {
	iterator, err := mem.GetAll(model.TableNameThing)
	i := model.IteratorToManyThings(iterator, err)
	return i

}

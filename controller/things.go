package controller

import (
	"fmt"
	"github.com/just1689/gg-bot-captain/mem"
	"github.com/just1689/gg-bot-captain/model"
	"github.com/just1689/gg-bot-captain/model/messages/incoming"
	"github.com/sirupsen/logrus"
)

func handleDynamicThings(b []byte) {
	things, errorBuild := incoming.BuildMessageShareDynamicThingsFromString(b)
	if errorBuild != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", errorBuild.Error()))
		return
	}

	go func(things *[]model.Thing) {
		c := make(chan model.Thing)
		go pushThingsAsync(c)
		count := 0
		for _, item := range *things {
			c <- item
			count++
		}
		close(c)
	}(&things.Things)
}

func pushThingsAsync(c chan model.Thing) {
	for thing := range c {
		mem.Push(model.TableNameThing, thing)
	}
}

//GetAllThings gets all things in database
func GetAllThings() chan model.Thing {
	iterator, err := mem.GetAll(model.TableNameThing)
	i := model.IteratorToManyThings(iterator, err)
	return i

}

package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/just1689/gg-bot-captain/util"
	log "github.com/sirupsen/logrus"
)

//TableNameTag is the name of the tag table
const TableNameTag = "Tag"

//Tag holds the player info
type Tag struct {
	Tag uint `json:"tag" bson:"tag"`
}

//BuildTagFromString builds from bytes
func BuildTagFromString(b []byte) (item Tag, err error) {
	if err = util.BytesToDecoder(b).Decode(&item); err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err

}

//IteratorToManyTags gets a list of tags
func IteratorToTag(iterator memdb.ResultIterator, err error) Tag {
	//This implementation is bad as
	// 1. it is blocking
	// 2. does not handle the case where 0 Tags are returned
	i := IteratorToManyTags(iterator, err)
	return <-i
}

//IteratorToManyTags gets a list of tags
func IteratorToManyTags(iterator memdb.ResultIterator, err error) chan Tag {
	c := make(chan Tag)

	go func() {
		if err != nil {
			log.Errorln(fmt.Sprintf("Error iterator to many tags: %v", err.Error()))
			close(c)
		}

		i := iteratorToChannel(iterator, err)
		counter := 0
		for next := range i {
			item := next.(Tag)
			c <- item
			counter++
		}
		log.Debugln(fmt.Sprintf("Tag iterator counts: %d", counter))
		close(c)
	}()
	return c

}

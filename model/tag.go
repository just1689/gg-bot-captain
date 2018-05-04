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
func IteratorToManyTags(iterator memdb.ResultIterator, err error) (chan Tag, error) {

	if err != nil {
		log.Errorln(fmt.Sprintf("Error iterator to many tags: %v", err.Error()))
		return nil, err
	}

	c := make(chan Tag)
	go func() {
		if iteratorUseful(iterator, err) == false {
			if err != nil {
				log.Errorln(fmt.Sprintf("Error iterator to many tags: %v", err.Error()))
			}
			close(c)
			return
		}
		counter := 0
		next := iterator.Next()
		for next != nil {
			counter++
			item := next.(Tag)
			c <- item
			next = iterator.Next()
		}
		log.Debugln(fmt.Sprintf("Tag iterator counts: %d", counter))
		close(c)
	}()
	return c, nil

}

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
	if err = util.Decode(b, item); err != nil {
		log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	return item, err

}

//IteratorToManyTags gets a list of tags
func IteratorToManyTags(iterator memdb.ResultIterator, err error) (items []Tag) {
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if iterator == nil {
		return items
	}
	more := true
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue
		}
		item := next.(Tag)
		items = append(items, item)
	}
	log.Debugln(fmt.Sprintf("Tag iterator counts: %d", len(items)))
	return items

}

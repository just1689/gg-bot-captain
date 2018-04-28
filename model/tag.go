package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

//TableNameTag is the name of the tag table
const TableNameTag = "Tag"

//ConversationShareTag is the tag name
const ConversationShareTag = "S_SHARE_TAG"

type Tag struct {
	Tag uint `json:"tag" bson:"tag"`
}

func BuildTagFromString(msg []byte) (Tag, error) {
	r := bytes.NewReader(msg)
	decoder := json.NewDecoder(r)
	var item Tag
	err := decoder.Decode(&item)
	if err != nil {
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

package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

//TableNameThing is the name of the table
const TableNameThing = "Thing"

//Thing is a skinned object with a point
type Thing struct {
	Point Point  `json:"point" bson:"point"`
	Tag   uint   `json:"tag" bson:"tag"`
	Speed string `json:"speed" bson:"speed"`
}

//IteratorToManyTags gets a list of tags
func IteratorToManyThings(iterator memdb.ResultIterator, err error) (items []Thing) {
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
		item := next.(Thing)
		items = append(items, item)
	}
	log.Debugln(fmt.Sprintf("Thing iterator counts: %d", len(items)))
	return items

}

package model

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
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
	if IteratorInvalid(iterator, err) {
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

//IteratorToFirstThing gets the first item of an iterator
func IteratorToFirstThing(iterator memdb.ResultIterator, err error) (item Thing, ok bool) {
	if IteratorInvalid(iterator, err) {
		return Thing{}, false
	}

	more := true
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue
		}
		item := next.(Thing)
		return item, true
	}
	return Thing{}, false

}

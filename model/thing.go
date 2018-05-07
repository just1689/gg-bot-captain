package model

import (
	"github.com/hashicorp/go-memdb"
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
func IteratorToManyThings(iterator memdb.ResultIterator, err error) (chan Thing) {
	c := make(chan Thing)
	go func() {
		i := iteratorToChannel(iterator, err)
		count := 0
		for next := range i {
			item := next.(Thing)
			c <- item
			count++
		}
		close(c)
		log.Debugln(fmt.Sprintf("Thing iterator counts: %d", count))

	}()
	return c

}

//IteratorToFirstThing gets the first item of an iterator
func IteratorToFirstThing(iterator memdb.ResultIterator, err error) (item Thing, ok bool) {
	if iteratorUseful(iterator, err) == false {
		return Thing{}, false
	}
	next := iterator.Next()
	if next != nil {
		item := next.(Thing)
		return item, true
	}
	return Thing{}, false

}

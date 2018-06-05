package model

import (
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

func iteratorUseful(iterator memdb.ResultIterator, err error) bool {
	if err != nil {
		log.Error(err.Error())
		return false
	}
	if nil == iterator {
		return false
	}
	return true
}

func iteratorToChannel(iterator memdb.ResultIterator, err error) chan interface{} {
	c := make(chan interface{})
	go func() {
		if iteratorUseful(iterator, err) == false {
			close(c)
		}
		next := iterator.Next()
		for next != nil {
			c <- next
			next = iterator.Next()
		}
		close(c)
	}()
	return c

}

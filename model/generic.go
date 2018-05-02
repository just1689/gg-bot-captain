package model

import (
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

//IteratorInvalid checks if an iterator is useful
func IteratorInvalid(iterator memdb.ResultIterator, err error) bool {
	if err != nil {
		log.Error(err.Error())
		return true
	}
	if iterator == nil {
		return true
	}
	return false
}

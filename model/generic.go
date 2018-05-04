package model

import (
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

//IteratorUseful checks if an iterator is useful
func IteratorUseful(iterator memdb.ResultIterator, err error) bool {
	if err != nil {
		log.Error(err.Error())
		return true
	}
	if iterator == nil {
		return true
	}
	return false
}

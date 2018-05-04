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

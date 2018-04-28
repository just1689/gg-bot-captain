package mem

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

var (
	db *memdb.MemDB
)

//Init sets up the database
func Init() {

	// Create a new data base
	dbConnect, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Errorln(fmt.Sprintf("Could not connect to database: %s", err.Error()))
		panic(err)
	}
	db = dbConnect

}

//GetUniqueById gets one row by unique ID field
func GetUniqueById(name string, id string) (interface{}, error) {
	txn := getTransaction(false)
	defer txn.Abort()
	raw, err := txn.First(name, "id", id)
	return raw, err

}

//GetAllByFieldAndValue gets rows in a table with a where clause for an index
func GetAllByFieldAndValue(name string, field string, value string) (memdb.ResultIterator, error) {
	txn := getTransaction(false)
	defer txn.Abort()
	raw, err := txn.Get(name, field, value)
	return raw, err

}

//GetAll retrieves a whole table
func GetAll(name string) (memdb.ResultIterator, error) {
	txn := getTransaction(false)
	defer txn.Abort()
	return txn.Get(name, "id")

}

func getTransaction(write bool) *memdb.Txn {
	return db.Txn(write)
}

//Push stores an object
func Push(table string, obj interface{}) error {
	log.Debugln(fmt.Sprintf("Trying to insert a: %s", table))
	txn := db.Txn(true)
	defer txn.Abort()
	if err := txn.Insert(table, obj); err != nil {
		log.Errorln(fmt.Sprintf("Error pushing to mem: %s", err))
		log.Errorln(fmt.Sprintf("Error pushing to mem: %s", err.Error()))
		defer txn.Abort()
		return err
	}
	txn.Commit()
	return nil
}

//Delete removes all rows with a value for an index
func Delete(name string, field string, id string) (int, error) {
	txn := getTransaction(true)
	i, err := txn.DeleteAll(name, field, id)
	if err != nil {
		txn.Abort()
	}
	txn.Commit()
	return i, err
}

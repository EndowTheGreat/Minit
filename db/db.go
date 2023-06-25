package db

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

func NewDB() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal("Failed to open db: ", err)
	}
	return db
}

package database

import (
	"log"
	"github.com/boltdb/bolt"
)

var Db *bolt.DB

func Setup() {
	var err error
	Db, err = bolt.Open(".gotasks.db", 0600, nil)
	//600 code means rw access to the user only => rw- --- ---
	if err != nil {
		log.Fatal(err)
	}
	Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
}
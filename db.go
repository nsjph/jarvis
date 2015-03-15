package main

import (
	"github.com/boltdb/bolt"
	"path/filepath"
)

func initDatabase(configDir string) *bolt.DB {

	dbpath := filepath.Join(configDir, "jarvis.db")

	db, err := bolt.Open(dbpath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	return db
}

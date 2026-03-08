package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("mongo", "")
	if err != nil {
		log.Fatalln("error open database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("error pinging database", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}

package database

import (
	"database/sql"
	"log"
)

const databaseFilepath = "database.sqlite"

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", databaseFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}

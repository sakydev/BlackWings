package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const databaseFilename = "database.sqlite"

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", getDatabaseFilePath())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getDatabaseFilePath() string {
	parentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/database/%s", parentDirectory, databaseFilename)
}

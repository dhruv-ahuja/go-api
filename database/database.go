package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	return db, err
}

func CreateTable(db *sql.DB) error {
	// path to the initial migration file
	sqlPath := "./migrations/001_init.up.sql"

	// query will contain the contents of our sql file in the form of byte slice
	// this can then be passed onto the database query
	// after conversion to string
	query, ioErr := os.ReadFile(sqlPath)
	if ioErr != nil {
		return ioErr
	}

	_, err := db.Exec(string(query))
	return err
}

func DropTable(db *sql.DB) error {
	// path to the migration file to drop table
	sqlPath := "./migrations/001_init.down.sql"

	query, ioErr := os.ReadFile(sqlPath)
	if ioErr != nil {
		return ioErr
	}

	_, err := db.Exec(string(query))
	return err
}

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

type Book struct {
	// all field tags must be exported through capitalization
	// if they are to be used in the JSON encodings
	ID int `json:"id"`
	// `omitempty` excludes the field from the JSON encoding if its empty
	ISBN   int    `json:"isbn,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year,omitempty"`
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

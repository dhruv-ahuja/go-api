package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func connectToDB(dbPath string) (*sql.DB, error) {
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

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY,
		isbn INTEGER,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		genre TEXT NOT NULL,
		year INTEGER
		);`

	_, err := db.Exec(query)
	return err
}

func dropTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS books;`
	_, err := db.Exec(query)
	return err
}

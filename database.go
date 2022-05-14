package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func connectToDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	return db, err
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

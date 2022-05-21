package database

import (
	"database/sql"
	"os"

	"github.com/dhruv-ahuja/go-api/helpers"
	_ "github.com/mattn/go-sqlite3"
)

// Init initializes the database connection and runs the CreateTable function
// for us in one place, helping declutter the main function.
func Init(dbPath string) *sql.DB {
	db, err := ConnectToDB(dbPath)

	helpers.CheckErr("error connecting to database: ", err)
	// directly feeding the CreateTable func to CheckErr
	helpers.CheckErr("error when creating table: ", createTable(db))

	return db
}

func ConnectToDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	return db, err
}

func createTable(db *sql.DB) error {
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

func dropTable(db *sql.DB) error {
	// path to the migration file to drop table
	sqlPath := "./migrations/001_init.down.sql"

	query, ioErr := os.ReadFile(sqlPath)
	if ioErr != nil {
		return ioErr
	}

	_, err := db.Exec(string(query))
	return err
}

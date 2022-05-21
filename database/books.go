// contains the code relating to the CRUD operations to be perfomed
// on the database

package database

import (
	"database/sql"
	"encoding/json"
)

type Book struct {
	// all field tags must be exported through capitalization
	// if they are to be used in the JSON encodings
	ID int `json:"id,omitempty"`
	// `omitempty` excludes the field from the JSON encoding if its empty
	ISBN   int    `json:"isbn,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genres string `json:"genres"`
	Year   int    `json:"year,omitempty"`
}

// GetBooks returns all books from the database
func GetBooks(db *sql.DB) ([]Book, error) {
	query := "SELECT * FROM books;"
	books := []Book{}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// rows is an iterable and we keep on iterating over the results till there
	// are none remaining
	for rows.Next() {
		var book Book
		// we scan and copy over each column value for each row fetched to
		// the structs' fields. the pointer points to their memory addresses
		// where the values are then written
		if err := rows.Scan(&book.ID, &book.ISBN, &book.Title, &book.Author,
			&book.Genres, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	// checking to see whether there was any error from the overall query
	// this is the only place from where we can learn if the query itself failed
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// AddBook receives the book in as POST body and adds it to the database
func AddBook(db *sql.DB, decoder *json.Decoder) (*Book, error) {
	book := &Book{}

	// this will return an error if there is a mismatch between the data received
	// in the POST request and the struct fields the data is being written to
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&book)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO books (isbn, title, author, genres, year) VALUES
(?, ?, ?, ?, ?) RETURNING *;`

	row := db.QueryRow(query, book.ISBN, book.Title, book.Author, book.Genres, book.Year)

	// res will store the result returned by the DB, to be sent back to the user
	// we do this so that the user can get to know the ID of the book that they have
	// just inserted
	res := &Book{}
	err = row.Scan(&res.ID, &res.ISBN, &res.Title, &res.Author, &res.Genres, &res.Year)
	if err != nil {
		return nil, err
	}

	return res, err
}

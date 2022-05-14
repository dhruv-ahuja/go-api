// contains the code relating to the CRUD operations to be perfomed
// on the database

package database

import "database/sql"

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

func GetBooks(db *sql.DB) ([]Book, error) {
	query := "SELECT * FROM books LIMIT 5;"
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
			&book.Genre, &book.Year); err != nil {
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

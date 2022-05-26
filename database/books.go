// contains the code relating to the CRUD operations to be perfomed
// on the database

package database

import (
	"database/sql"
	"encoding/json"
)

type BookStore interface {
	GetBook(int) (*Book, error)
	GetBooks() ([]Book, error)
	AddBook(*json.Decoder) (*Book, error)
	UpdateBook(*json.Decoder, int) (*Book, error)
	DeleteBook(int) (sql.Result, error)
}

// Book contains the properties that the books stored in the DB have.
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

// BookShelf is the struct that implements BookStore here.
// The 'Book' struct is used to store data when performing operations on the
// BookShelf. So, a BookShelf stores many books and acts as a small-scale
// representation of what a BookStore would be.
type BookShelf struct {
	DB *sql.DB
}

func NewShelf(db *sql.DB) *BookShelf {
	return &BookShelf{
		DB: db,
	}
}

// GetBook returns a book from the database given its ID
func (b *BookShelf) GetBook(bookID int) (*Book, error) {
	query := `SELECT * FROM books WHERE id=?;`
	book := &Book{}

	row := b.DB.QueryRow(query, bookID)
	// we always need to be using pointers with Scan, doesn't matter if the
	// destination struct has been intiliazed with its pointer or not
	if err := row.Scan(&book.ID, &book.ISBN, &book.Title, &book.Author,
		&book.Genres, &book.Year); err != nil {
		return nil, err
	}

	return book, nil
}

// GetBooks returns all books from the database
func (b *BookShelf) GetBooks() ([]Book, error) {
	query := "SELECT * FROM books;"
	books := []Book{}

	rows, err := b.DB.Query(query)
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
func (b *BookShelf) AddBook(decoder *json.Decoder) (*Book, error) {
	book := &Book{}

	// this will return an error if there is a mismatch between the data received
	// in the POST request and the struct fields the data is being written to
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(book); err != nil {
		return nil, err
	}

	query := `INSERT INTO books (isbn, title, author, genres, year) VALUES
(?, ?, ?, ?, ?) RETURNING *;`

	row := b.DB.QueryRow(query, book.ISBN, book.Title, book.Author,
		book.Genres, book.Year)

	// res will store the result returned by the DB, to be sent back to the user
	// we do this so that the user can get to know the ID of the book that they have
	// just inserted
	res := &Book{}
	if err := row.Scan(&res.ID, &res.ISBN, &res.Title, &res.Author,
		&res.Genres, &res.Year); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateBook receives the book to be updated as POST body and updates it in
// the database.
func (b *BookShelf) UpdateBook(decoder *json.Decoder, bookID int) (*Book, error) {
	book := &Book{}

	decoder.DisallowUnknownFields()
	if err := decoder.Decode(book); err != nil {
		return nil, err
	}

	query := `UPDATE books SET isbn=?, title=?, author=?, genres=?, year=?
	WHERE id=? RETURNING *;`

	// bookID is entered separately since its being received through the URL
	row := b.DB.QueryRow(query, book.ISBN, book.Title, book.Author,
		book.Genres, book.Year, bookID)

	res := &Book{}
	if err := row.Scan(&res.ID, &res.ISBN, &res.Title, &res.Author,
		&res.Genres, &res.Year); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteBook receives the book to deleted as POST body and remvoes it from
// the database.
func (b *BookShelf) DeleteBook(bookID int) (sql.Result, error) {
	query := `DELETE FROM books WHERE id=?;`

	res, err := b.DB.Exec(query, bookID)

	if err != nil {
		return nil, err
	}

	return res, nil
}

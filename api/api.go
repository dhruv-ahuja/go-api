package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhruv-ahuja/go-api/database"
	"github.com/dhruv-ahuja/go-api/helpers"
)

type Connection struct {
	DB *sql.DB
}

// defining the struct to be used during responses
// structs are converted to JSON using the `marshal` function
type jsonResponse struct {
	Message string   `json:"message"`
	Data    []string `json:"data,omitempty"`
}

func NewConnection(db *sql.DB) *Connection {
	return &Connection{
		DB: db,
	}
}

// Index is the entrypoint to the api
func (c *Connection) Index(w http.ResponseWriter, r *http.Request) {
	msg := jsonResponse{
		Message: "Hello, World!",
	}

	data, err := json.Marshal(msg)
	helpers.CheckErr("error converting data to JSON: ", err)

	// data marshaled to JSON is in the form of a slice of bytes, convert it to
	// string to make it usable for the writer
	fmt.Fprintln(w, string(data))
}

func (c *Connection) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.GetBooks(c.DB)
	helpers.CheckErr("error fetching books from the DB", err)

	data, err := json.Marshal(books)
	helpers.CheckErr("error converting to JSON: ", err)

	fmt.Fprint(w, string(data))
}

func (c *Connection) AddABook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		msg := jsonResponse{
			Message: "Add a book by sending a POST request.",
		}

		data, err := json.Marshal(msg)
		helpers.CheckErr("error converting data to JSON: ", err)

		fmt.Fprintln(w, string(data))

	case "POST":
		decoder := json.NewDecoder(r.Body)

		book, err := database.AddBook(c.DB, decoder)
		helpers.CheckErr("error adding book to database: ", err)

		// setting the header status code to 201/Created to indicate success
		// with creating a new resource
		w.WriteHeader(http.StatusCreated)
		// set the content type so that the user knows what type of data to expect
		w.Header().Set("Content-Type", "application/json")

		data, err := json.Marshal(book)
		helpers.CheckErr("error converting data to JSON: ", err)

		fmt.Fprintln(w, string(data))
	}
}

package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhruv-ahuja/go-api/database"
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

// index is the default endpoint for the api
func (c Connection) index(w http.ResponseWriter, r *http.Request) {
	msg := jsonResponse{
		Message: "Hello, World!",
	}

	data, err := json.Marshal(msg)
	checkErr("", err)

	fmt.Fprintln(w, string(data))
}

func (c Connection) allBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.GetBooks(c.DB)
	checkErr("error while fetching books from the DB", err)

	data, err := json.Marshal(books)
	checkErr("", err)

	fmt.Fprint(w, string(data))
}

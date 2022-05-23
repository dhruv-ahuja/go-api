package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dhruv-ahuja/go-api/database"
	"github.com/dhruv-ahuja/go-api/helpers"
	"github.com/go-chi/chi/v5"
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

	w.Header().Set("Content-Type", "application/json")
	// data marshaled to JSON is in the form of a slice of bytes, convert it to
	// string to make it usable for the writer
	fmt.Fprintln(w, string(data))
}

// AddABook performs the Create operation of the API
func (c *Connection) AddABook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	book, err := database.AddBook(c.DB, decoder)
	helpers.CheckErr("error adding book to database: ", err)

	data, err := json.Marshal(book)
	helpers.CheckErr("error converting data to JSON: ", err)

	// set the content type so that the user knows what type of data to expect
	w.Header().Set("Content-Type", "application/json")

	// setting the header status code to 201/Created to indicate success
	// with creating a new resource
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, string(data))
}

func (c *Connection) GetABook(w http.ResponseWriter, r *http.Request) {
	getID := chi.URLParam(r, "id")

	if getID != "" {
		bookID, err := strconv.Atoi(getID)
		helpers.CheckErr("error converting string to int: ", err)

		book, err := database.GetBook(c.DB, bookID)
		// if the error is 'ErrNoRows' that means that no data was found
		// for that ID
		if err != nil && err != sql.ErrNoRows {
			helpers.CheckErr("error fetching book from the DB: ", err)
		}

		// in case we don't receive any data from the db, we set the
		// 404 status code to indicate resource wasn't found
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		}

		data, err := json.Marshal(book)
		helpers.CheckErr("error converting to JSON: ", err)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(data))
	}
}

// GetAllBooks performs the Read operation of the API
func (c *Connection) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := database.GetBooks(c.DB)
	if err != nil && err != sql.ErrNoRows {
		helpers.CheckErr("error fetching books from the DB", err)
	}

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
	}

	data, err := json.Marshal(books)
	helpers.CheckErr("error converting to JSON: ", err)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

// UpdateABook performs the UPDATE operation of the API
func (c *Connection) UpdateABook(w http.ResponseWriter, r *http.Request) {
	// we send the ID in the URL and the book data in the request body
	// the ID is verified and only then is the entry overwritten with the new data
	getID := chi.URLParam(r, "id")

	if getID != "" {
		bookID, err := strconv.Atoi(getID)
		helpers.CheckErr("error converting string to int: ", err)

		decoder := json.NewDecoder(r.Body)
		book, err := database.UpdateBook(c.DB, decoder, bookID)
		if err != nil && err != sql.ErrNoRows {
			helpers.CheckErr("error updating book in database: ", err)
		}

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		}

		data, err := json.Marshal(book)
		helpers.CheckErr("error converting data to JSON: ", err)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(data))

	}
}

// DeleteABook performs the DELETE operation of the API
func (c *Connection) DeleteABook(w http.ResponseWriter, r *http.Request) {
	// URL param fetches the named parameter in the URL
	// in our case, `/books/{id}`: here 'id' is the url param
	getID := chi.URLParam(r, "id")

	if getID != "" {
		// since the param is sent with the URL, it is in the string form
		bookID, err := strconv.Atoi(getID)
		helpers.CheckErr("error converting string to int: ", err)

		// we have to check whether the ID was valid or not
		// the query result will help us do that
		res, err := database.DeleteBook(c.DB, bookID)
		helpers.CheckErr("error deleting book from database: ", err)

		// the ID was invalid if no rows were affected, so we just return a 404
		if rows, _ := res.RowsAffected(); rows == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			// StatusNoContent or Status 204 indicates that the request was fulfilled
			// we don't need to send any data back, ideal response for a delete request
			w.WriteHeader(http.StatusNoContent)
		}

	}
}

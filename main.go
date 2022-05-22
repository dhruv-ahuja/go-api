package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dhruv-ahuja/go-api/api"
	"github.com/dhruv-ahuja/go-api/database"
	"github.com/dhruv-ahuja/go-api/helpers"
	"github.com/joho/godotenv"
)

func main() {
	// loading .env file and checking for errors directly
	helpers.CheckErr("error loading .env file: ", godotenv.Load())

	dbPath := os.Getenv("DB")
	if dbPath == "" {
		dbPath = "./app.db"
	}

	// Init connects to the DB and also runs the createTable func
	db := database.Init(dbPath)
	defer db.Close()

	// creating a Connection instance to use to register handlers for the
	// web server
	c := api.NewConnection(db)

	fmt.Println("live on port 8080...")

	http.HandleFunc("/", c.Index)
	http.HandleFunc("/books", c.GetAllBooks)
	http.HandleFunc("/add/books", c.AddABook)
	http.HandleFunc("/put/books", c.UpdateABook)

	err := http.ListenAndServe("localhost:8080", nil)
	helpers.CheckErr("error when serving endpoints: ", err)
}

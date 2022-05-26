package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dhruv-ahuja/go-api/api"
	"github.com/dhruv-ahuja/go-api/database"
	"github.com/dhruv-ahuja/go-api/helpers"
	"github.com/go-chi/chi/v5"
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

	// implement chi as the router of choice to get more functionality over
	// the net/http router
	r := chi.NewRouter()

	// our BookShelf struct implements the Store interface
	var store database.BookStore
	bookShelf := database.NewShelf(db)
	// this store will be used by the Connection struct
	store = bookShelf

	// creating a Connection instance to use to register handlers for the
	// web server
	c := api.NewConnection(store, r)
	c.SetupRoutes(c.Router)

	fmt.Println("live on http://localhost:8080...")

	err := http.ListenAndServe("localhost:8080", c.Router)
	helpers.CheckErr("error when serving endpoints: ", err)
}

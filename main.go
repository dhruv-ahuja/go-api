package main

import (
	"encoding/json"
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

	// creating a Connection instance to use to register handlers for the
	// web server
	c := api.NewConnection(db)
	// implement chi as the router of choice to get more functionality over
	// the net/http router
	r := chi.NewRouter()

	fmt.Println("live on port 8080...")

	r.HandleFunc("/", c.Index)

	// making a custom 404 page
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		msg := api.JsonResponse{
			Message: "invalid URL or page not found",
		}

		data, err := json.Marshal(msg)
		helpers.CheckErr("Error converting data to JSON: ", err)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(data))
	})

	// using subrouter for the routes sharing the same URL but using different
	// request types
	r.Route("/books", func(r chi.Router) {
		r.Get("/", c.GetAllBooks)
		r.Post("/", c.AddABook)

		r.Get("/{id}", c.GetABook)
		r.Put("/{id}", c.UpdateABook)
		r.Delete("/{id}", c.DeleteABook)
	})

	err := http.ListenAndServe("localhost:8080", r)
	helpers.CheckErr("error when serving endpoints: ", err)
}

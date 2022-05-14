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

	db, err := database.ConnectToDB(dbPath)
	helpers.CheckErr("error connecting to database: ", err)
	defer db.Close()

	helpers.CheckErr("error when creating table: ", database.CreateTable(db))
	// api.CheckErr("error when dropping table: ", database.DropTable(db))

	fmt.Println("live on port 8080...")

	// creating a struct instance
	c := api.NewConnection(db)

	http.HandleFunc("/", c.Index)
	http.HandleFunc("/books", c.GetBooks)

	err = http.ListenAndServe("localhost:8080", nil)
	helpers.CheckErr("", err)
}

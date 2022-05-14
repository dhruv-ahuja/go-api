package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dhruv-ahuja/go-api/database"
	"github.com/joho/godotenv"
)

// defining the struct to be used during responses
// structs are converted to JSON using the `marshal` function
type jsonResponse struct {
	Message string   `json:"message"`
	Data    []string `json:"data,omitempty"`
}

func main() {
	// loading .env file and checking for errors directly
	checkErr("error loading .env file: ", godotenv.Load())

	dbPath := os.Getenv("DB")
	if dbPath == "" {
		dbPath = "./app.db"
	}

	db, err := database.ConnectToDB(dbPath)
	checkErr("error connecting to database: ", err)
	defer db.Close()

	checkErr("error when creating table: ", database.CreateTable(db))
	// checkErr("error when dropping table: ", database.DropTable(db))

	fmt.Println("live on port 8080...")

	http.HandleFunc("/", index)
	err = http.ListenAndServe("localhost:8080", nil)
	checkErr("", err)
}

// index is the default endpoint for the api
func index(w http.ResponseWriter, r *http.Request) {
	msg := jsonResponse{
		Message: "Hello, World!",
	}

	data, err := json.Marshal(msg)
	checkErr("", err)

	fmt.Fprintln(w, string(data))
}

// checkErr checks for error in given functions/methods. It also outputs an
// error message, if given one.
func checkErr(errMsg string, err error) {
	if err != nil {
		log.Fatal(errMsg, err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// defining the structs to be used during responses
// structs are converted to JSON
type jsonResponse struct {
	Message string   `json:"message"`
	Data    []string `json:"data,omitempty"`
}

func main() {
	fmt.Println("live on port 8080...")

	http.HandleFunc("/", index)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{
		"message": "hello",
	}
	data, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(data))
}

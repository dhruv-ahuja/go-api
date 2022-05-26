package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhruv-ahuja/go-api/helpers"
	"github.com/go-chi/chi/v5"
)

func (c *Connection) SetupRoutes(r *chi.Mux) {
	r.HandleFunc("/", c.HealthCheck)

	// making a custom 404 page
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		msg := JsonResponse{
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

}

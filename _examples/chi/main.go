package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/olacin/sentinel"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(sentinel.DefaultHandler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":3000", r)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world!</h1>")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", indexHandler)

	log.Print("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

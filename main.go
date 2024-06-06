package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/vite-golang-stream/templates"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home().Render(r.Context(), w)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", indexHandler)

	fs := http.FileServer(http.Dir("./priv/assets/assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	log.Print("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

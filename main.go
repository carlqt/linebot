package main

import (
	"log"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", index)
	r.Post("/payload", payload)
	r.Get("/user/:id/new", newUser)
	log.Println("listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func payload(w http.ResponseWriter, r *http.Request) {
	jsonRequest(r)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	w.Write([]byte("welcome " + userID))
}

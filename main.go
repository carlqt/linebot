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

	// log.Println(devexcuse.Excuse())

	r.Route("/line", func(r chi.Router) {
		r.Use(validateRequest)
		r.Use(validateSignature)
		r.Use(replyHandler)
		r.Post("/reply", lineReply)
	})

	log.Println("listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

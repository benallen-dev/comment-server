package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/a-h/templ"
	// "github.com/benallen-dev/comment-server/data"
	
	"github.com/benallen-dev/comment-server/pkg/handlers"
)

func main() {
	log.Println("Hello World")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Route("/comments", func(r chi.Router) {
		r.Get("/", handlers.GetAllComments)
		r.Get("/{slug}", handlers.GetCommentForSlug)
		r.Post("/{slug}", handlers.AddCommentToSlug)
	})

	http.ListenAndServe(":3000", r)

}

package handlers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"

	"github.com/benallen-dev/comment-server/pkg/data"
	"github.com/benallen-dev/comment-server/pkg/views"
)

// GetCommentHandler handles requests for comments for a given slug
func GetCommentForSlug(w http.ResponseWriter, r *http.Request) {

	slug := chi.URLParam(r, "slug")

	log.Println("GetCommentForSlug", slug)

	// Get the comments for the slug
	comments, err := data.GetCommentsForSlug(slug)
	if err != nil {
		log.Panic(err)
		return
	}

	templ.Handler(views.Comments(comments)).ServeHTTP(w, r)
}


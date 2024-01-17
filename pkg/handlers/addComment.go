package handlers

import (
	//	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/benallen-dev/comment-server/pkg/data"
	"github.com/benallen-dev/comment-server/pkg/views"
)

// GetCommentHandler handles requests for comments for a given slug
func AddCommentToSlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	// TODO: Figure out how to extract all these values from JSON data
	text := r.FormValue("text")
	author := r.FormValue("author")
	email := r.FormValue("email")

	dateTime := time.Now().Format("2006-01-02 15:04:05")
	id, _ := uuid.NewUUID() // TODO: Handle error

	newComment := data.Comment{
		ID:       id,
		Slug:     slug,
		Author:   author,
		Email:    email,
		DateTime: dateTime,
		Text:     text,
	}

	data.AddComment(newComment)

	log.Println("AddCommentForSlug", slug)

	// Return a comment component to add to the array on the page
	templ.Handler(views.Comment(newComment)).ServeHTTP(w, r)
}

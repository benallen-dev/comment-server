
package handlers

import (
//	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/benallen-dev/comment-server/pkg/data"
	"github.com/benallen-dev/comment-server/pkg/views"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	
	// Get the comments for the slug
	comments, err := data.GetAllComments()
	if err != nil {
		log.Panic(err)
		return
	}

	templ.Handler(views.Comments(comments)).ServeHTTP(w, r)
}

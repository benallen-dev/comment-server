
package handlers

import (
	"encoding/json"
	"log"
	"net/http"


	"github.com/benallen-dev/comment-server/pkg/data"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	
	// Get the comments for the slug
	comments, err := data.GetAllComments()
	if err != nil {
		log.Panic(err)
		return
	}

	jsonComments, err := json.Marshal(comments)
	if err != nil {
		log.Panic(err)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonComments)
}

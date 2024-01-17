package data

import (
	"encoding/json"
	"os"
	"log"
)

var commentsStore map[string][]Comment

func init() {
	commentsStore = make(map[string][]Comment)
	data, err := ReadStoreFromDisk()
	if err != nil {
		log.Println(err)
		return
	}

	commentsStore = data
}

func getFakeComment(slug string) Comment {
	return Comment{
		ID:       "1",
		Slug:     slug,
		Author:   "Ben",
		Email:    "ballen@anwb.nl",
		DateTime: "2024-01-11 12:00:00",
		Text:     "This is a fake comment",
	}
}

func GetCommentsForSlug(slug string) ([]Comment, error) {

	return []Comment{
		getFakeComment(slug),
		getFakeComment(slug),
		getFakeComment(slug),
	}, nil
}

func AddComment(comment Comment) error {
	commentsForSlug := commentsStore[comment.Slug]

	if commentsForSlug == nil {
		commentsStore[comment.Slug] = []Comment{comment}

	} else {
		commentsStore[comment.Slug] = append(commentsForSlug, comment)
	}

	err := WriteStoreToDisk()
	if err != nil {
		return err
	}

	return nil
}

func GetAllComments() ([]Comment, error) {
	comments := []Comment{}

	for _, commentsForSlug := range commentsStore {
		comments = append(comments, commentsForSlug...)
	}

	log.Println("GetAllComments", comments)

	return comments, nil
}

func WriteStoreToDisk() error {
	storeJson, err := json.Marshal(commentsStore)
	if err != nil {
		return err
	}

	err = os.WriteFile("comments.json", storeJson, 0644)
	return nil
}

func ReadStoreFromDisk() (map[string][]Comment, error) {
	storeJson, err := os.ReadFile("comments.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(storeJson, &commentsStore)
	if err != nil {
		return nil, err
	}

	return commentsStore, nil
}

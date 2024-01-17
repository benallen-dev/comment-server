package data

import (
	"encoding/json"
	"os"
)

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

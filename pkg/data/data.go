package data

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

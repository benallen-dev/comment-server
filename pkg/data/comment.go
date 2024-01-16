package data

import (
	"fmt"
)

// Comment represents a comment on a page.
// 
// Ideally you'd seperate the meta and content data but let's keep it simple.
type Comment struct {
	// The unique identifier for the comment.
	ID string `json:"id"`
	// The slug of the page the comment is on.
	Slug string `json:"slug"`
	// The person who made the comment.
	Author string `json:"name"`
	// The email address of the person who made the comment.
	Email string `json:"email"`
	// The date the comment was made.
	DateTime string `json:"date"`
	// The comment text.
	Text string `json:"text"`
}

func (c *Comment) String() string {
	return fmt.Sprintf("%s - %s <%s>: %s", c.DateTime, c.Author, c.Email, c.Text)
}

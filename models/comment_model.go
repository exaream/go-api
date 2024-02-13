package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	ArticleID int       `json:"article_id"`
	Body      string    `json:"body"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewComment(id, articleID int, body, userName string, createdAt, updatedAt time.Time) *Comment {
	return &Comment{
		ID:        id,
		ArticleID: articleID,
		Body:      body,
		UserName:  userName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

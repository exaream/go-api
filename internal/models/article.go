package models

import "time"

type Article struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Body        string     `json:"body"`
	UserName    string     `json:"user_name"`
	NiceNum     int        `json:"nice_num"`
	CommentList []*Comment `json:"comment_list"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"update_at"`
}

func NewArticle(id int, title, body, userName string, niceNum int, createdAt, updatedAt time.Time) *Article {
	return &Article{
		ID:        id,
		Title:     title,
		Body:      body,
		UserName:  userName,
		NiceNum:   niceNum,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

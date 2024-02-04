package models

import "time"

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Comment `json:"comment_list"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type Comment struct {
	ID        int       `json:"id"`
	ArticleID int       `json:"article_id"`
	Body      string    `json:"body"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

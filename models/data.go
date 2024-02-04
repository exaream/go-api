package models

import "time"

var (
	Article1 = Article{
		ID:          1,
		UserName:    "Alice",
		Title:       "title1",
		Body:        "body1",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:          2,
		UserName:    "Bob",
		Title:       "title2",
		Body:        "body2",
		NiceNum:     2,
		CommentList: []Comment{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	Comment1 = Comment{
		ID:        1,
		ArticleID: 1,
		UserName:  "Chris",
		Body:      "comment1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	Comment2 = Comment{
		ID:        2,
		ArticleID: 1,
		UserName:  "Dorothy",
		Body:      "comment2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

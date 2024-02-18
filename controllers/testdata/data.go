package testdata

import "github.com/exaream/go-api/models"

var ArticleList = []*models.Article{
	&models.Article{
		ID:       1,
		Title:    "article title 1",
		Body:     "article body 1",
		UserName: "Alice",
		NiceNum:  2,
	},
	&models.Article{
		ID:       2,
		Title:    "article title 2",
		Body:     "article body 2",
		UserName: "Alice",
		NiceNum:  4,
	},
}

var CommentList = []*models.Comment{
	&models.Comment{
		ID:        1,
		ArticleID: 1,
		Body:      "article 1, comment 1",
		UserName:  "Bob",
	},
	&models.Comment{
		ID:        2,
		ArticleID: 1,
		Body:      "article 1, comment 2",
		UserName:  "Chris",
	},
}

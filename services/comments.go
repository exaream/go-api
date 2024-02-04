package services

import (
	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
)

func PostComment(comment *models.Comment) (*models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return repositories.InsertComment(db, comment)
}

package testdata

import "github.com/exaream/go-api/internal/models"

type commentServiceMock struct{}

func NewCommentServiceMock() *commentServiceMock {
	return &commentServiceMock{}
}

func (s *commentServiceMock) PostComment(comment *models.Comment) (*models.Comment, error) {
	return CommentList[0], nil
}

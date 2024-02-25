package testdata

import "github.com/exaream/go-api/internal/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) ListArticle(page int) ([]*models.Article, error) {
	return ArticleList, nil
}

func (s *serviceMock) GetArticle(id int) (*models.Article, error) {
	return ArticleList[0], nil
}

func (s *serviceMock) PostArticle(article *models.Article) (*models.Article, error) {
	return ArticleList[0], nil
}

func (s *serviceMock) PostNice(article *models.Article) (*models.Article, error) {
	return ArticleList[0], nil
}

func (s *serviceMock) PostComment(comment *models.Comment) (*models.Comment, error) {
	return CommentList[0], nil
}

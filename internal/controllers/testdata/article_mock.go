package testdata

import "github.com/exaream/go-api/internal/models"

type articleServiceMock struct{}

func NewArticleServiceMock() *articleServiceMock {
	return &articleServiceMock{}
}

func (s *articleServiceMock) List(page int) ([]*models.Article, error) {
	return ArticleList, nil
}

func (s *articleServiceMock) GetByID(id int) (*models.Article, error) {
	return ArticleList[0], nil
}

func (s *articleServiceMock) Post(article *models.Article) (*models.Article, error) {
	return ArticleList[0], nil
}

func (s *articleServiceMock) PostNice(article *models.Article) (*models.Article, error) {
	return ArticleList[0], nil
}

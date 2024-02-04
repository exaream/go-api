package services

import (
	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
)

func (s *AppService) GetArticleList(page int) ([]*models.Article, error) {
	return repositories.SelectArticleList(s.db, page)
}

func (s *AppService) GetArticleDetail(articleID int) (*models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return nil, err
	}

	comments, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return nil, err
	}

	article.CommentList = comments

	return article, nil
}

func (s *AppService) PostArticle(article *models.Article) (*models.Article, error) {
	return repositories.InsertArticle(s.db, article)
}

func (s *AppService) PostNice(article *models.Article) (*models.Article, error) {
	return repositories.UpdateNiceNum(s.db, article.ID)
}

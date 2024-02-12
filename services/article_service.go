package services

import (
	"database/sql"
	"errors"

	"github.com/exaream/go-api/apperrors"
	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
)

func (s *AppService) GetArticleList(page int) ([]*models.Article, error) {
	list, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, apperrors.FailedToSelect.Wrap(err, "failed to get article list")
	}

	if len(list) == 0 {
		return nil, apperrors.NotFound.Wrap(ErrNotFound, "there is no article")
	}

	return list, nil
}

func (s *AppService) GetArticleDetail(articleID int) (*models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound.Wrap(err, "there is no target article")
		}

		return nil, apperrors.FailedToSelect.Wrap(err, "failed to get article detail")
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		// if no comments, just return article
		if errors.Is(err, sql.ErrNoRows) {
			return article, nil
		}

		return nil, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *AppService) PostArticle(article *models.Article) (*models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return nil, apperrors.FailedToInsert.Wrap(err, "failed to insert article")
	}

	return newArticle, nil
}

func (s *AppService) PostNice(article *models.Article) (*models.Article, error) {
	article, err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFoundToUpdate.Wrap(err, "there is no target article")
		}

		return nil, apperrors.FailedToUpdate.Wrap(err, "failed to update nice count")
	}

	return article, nil
}

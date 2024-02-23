package services

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/exaream/go-api/apperrors"
	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
)

func (s *AppService) ListArticle(page int) ([]*models.Article, error) {
	list, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, apperrors.FailedToSelect.Wrap(err, "failed to get article list")
	}

	if len(list) == 0 {
		return nil, apperrors.NotFound.Wrap(ErrNotFound, "there is no article")
	}

	return list, nil
}

func (s *AppService) GetArticle(articleID int) (*models.Article, error) {
	var (
		article                    *models.Article
		commentList                []*models.Comment
		articleErr, commentErr     error
		articleMutex, commentMutex sync.Mutex
		wg                         sync.WaitGroup
	)

	wg.Add(2)

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		articleMutex.Lock()
		article, articleErr = repositories.SelectArticleDetail(db, articleID)
		articleMutex.Unlock()
	}(s.db, articleID)

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		commentMutex.Lock()
		commentList, commentErr = repositories.SelectCommentList(db, articleID)
		commentMutex.Unlock()
	}(s.db, articleID)

	wg.Wait()

	if articleErr != nil {
		if errors.Is(articleErr, sql.ErrNoRows) {
			return nil, apperrors.NotFound.Wrap(articleErr, "there is no target article")
		}

		return nil, apperrors.FailedToSelect.Wrap(articleErr, "failed to get article detail")
	}

	if commentErr != nil {
		// if no comments, just return article
		if errors.Is(commentErr, sql.ErrNoRows) {
			return article, nil
		}

		return nil, apperrors.FailedToSelect.Wrap(commentErr, "failed to get comment list")
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

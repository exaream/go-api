package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/exaream/go-api/internal/apperrors"
	"github.com/exaream/go-api/internal/models"
	"github.com/exaream/go-api/internal/repositories"
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
	type articleRes struct {
		article *models.Article
		err     error
	}

	var article *models.Article
	var articleErr error
	articleChan := make(chan articleRes)
	defer close(articleChan)

	type commentRes struct {
		commentList []*models.Comment
		err         error
	}

	var commentList []*models.Comment
	var commentErr error
	commentChan := make(chan commentRes)
	defer close(commentChan)

	go func(ch chan<- articleRes, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleRes{article, err}
	}(articleChan, s.db, articleID)

	go func(ch chan<- commentRes, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentRes{commentList, err}
	}(commentChan, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case res := <-articleChan:
			article, articleErr = res.article, res.err
		case res := <-commentChan:
			commentList, commentErr = res.commentList, res.err
		}
	}

	if articleErr != nil {
		if errors.Is(articleErr, sql.ErrNoRows) {
			return nil, apperrors.NotFound.Wrap(articleErr, fmt.Sprintf("not found article_id: %d", articleID))
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
			return nil, apperrors.NotFoundToUpdate.Wrap(err, fmt.Sprintf("not found article_id: %d", article.ID))
		}

		return nil, apperrors.FailedToUpdate.Wrap(err, "failed to update nice count")
	}

	return article, nil
}

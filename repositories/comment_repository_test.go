package repositories_test

import (
	"testing"

	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
	"github.com/exaream/go-api/repositories/testdata"
)

func TestInsertComment(t *testing.T) {
	wawntID := 3
	want := &models.Comment{
		ArticleID: 1,
		Body:      "test comment",
		UserName:  "test user",
	}

	got, err := repositories.InsertComment(testDB, want)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != wawntID {
		t.Errorf("id, want: %v, got: %v", wawntID, got.ID)
	}

	if got.ArticleID != want.ArticleID {
		t.Errorf("article_id, want: %v, got: %v", want.ArticleID, got.ArticleID)
	}

	if got.Body != want.Body {
		t.Errorf("body, want: %v, got: %v", want.Body, got.Body)
	}

	if got.UserName != want.UserName {
		t.Errorf("user_name, want: %v, got: %v", want.UserName, got.UserName)
	}

	t.Cleanup(func() {
		const query1 = `DELETE FROM comments WHERE id = ?;`
		if _, err := testDB.Exec(query1, got.ID); err != nil {
			t.Fatal(err)
		}

		const query2 = `ALTER TABLE comments AUTO_INCREMENT = 2;`
		if _, err := testDB.Exec(query2); err != nil {
			t.Fatal(err)
		}
	})
}

func TestSelectCommentList(t *testing.T) {
	articleID := 1

	res, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range res {
		want := testdata.CommentList[i]

		if got.ArticleID != articleID {
			t.Errorf("article_id, want: %v, got: %v", articleID, got.ArticleID)
		}

		if got.Body != want.Body {
			t.Errorf("body, want: %v, got: %v", want.Body, got.Body)
		}

		if got.UserName != want.UserName {
			t.Errorf("user_name, want: %v, got: %v", want.UserName, got.UserName)
		}
	}
}

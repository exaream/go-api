package repositories_test

import (
	"testing"

	"github.com/exaream/go-api/models"
	"github.com/exaream/go-api/repositories"
	"github.com/exaream/go-api/repositories/testdata"
)

func TestInsertArticle(t *testing.T) {
	want := &models.Article{
		Title:    "new article",
		Body:     "new body",
		UserName: "user",
		NiceNum:  0,
	}

	wantID := 3

	got, err := repositories.InsertArticle(testDB, want)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != wantID {
		t.Errorf("id, want: %v, got: %v", wantID, got.ID)
	}

	if got.Title != want.Title {
		t.Errorf("title, want: %v, got: %v", want.Title, got.Title)
	}

	if got.Body != want.Body {
		t.Errorf("body, want: %v, got: %v", want.Body, got.Body)
	}

	if got.UserName != want.UserName {
		t.Errorf("user_name, want: %v, got: %v", want.UserName, got.UserName)
	}

	if got.NiceNum != want.NiceNum {
		t.Errorf("nice_num, want: %v, got: %v", want.NiceNum, got.NiceNum)
	}

	t.Cleanup(func() {
		const query1 = `DELETE FROM articles WHERE id = ?;`
		if _, err := testDB.Exec(query1, got.ID); err != nil {
			t.Fatal(err)
		}

		const query2 = `ALTER TABLE articles AUTO_INCREMENT = 2;`
		if _, err := testDB.Exec(query2); err != nil {
			t.Fatal(err)
		}
	})
}

func TestSelectArticleList(t *testing.T) {
	want := len(testdata.ArticleList)
	list, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if got := len(list); got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		name string
		want *models.Article
	}{
		{
			name: "first article",
			want: testdata.ArticleList[0],
		},
		{
			name: "second article",
			want: testdata.ArticleList[1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, tt.want.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != tt.want.ID {
				t.Errorf("id, want: %v, got: %v", tt.want.ID, got.ID)
			}

			if got.Title != tt.want.Title {
				t.Errorf("title, want: %v, got: %v", tt.want.Title, got.Title)
			}

			if got.Body != tt.want.Body {
				t.Errorf("body, want: %v, got: %v", tt.want.Body, got.Body)
			}

			if got.UserName != tt.want.UserName {
				t.Errorf("user_name, want: %v, got: %v", tt.want.UserName, got.UserName)
			}

			if got.NiceNum != tt.want.NiceNum {
				t.Errorf("nice_num, want: %v, got: %v", tt.want.NiceNum, got.NiceNum)
			}
		})
	}
}

func TestUpdateNice(t *testing.T) {
	articleID := 1

	got, err := repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	want := testdata.ArticleList[articleID-1].NiceNum + 1
	if got.NiceNum != want {
		t.Errorf("nice_num, want: %v, got: %v", want, got.NiceNum)
	}
}

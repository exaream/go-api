package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListArticle(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query string
		code  int
	}{
		{"number", "1", http.StatusOK},
		{"alphabet", "a", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			articleCtrl.ListArticle(res, req)

			if res.Code != tt.code {
				t.Errorf("code, got: %d, want: %d", res.Code, tt.code)
			}
		})
	}
}

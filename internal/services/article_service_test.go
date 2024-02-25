package services_test

import "testing"

func BenchmarkGetArticle(b *testing.B) {
	articleID := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := articleService.GetArticle(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}

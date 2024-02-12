package middlewares

import "net/http"

func Apply(
	handler http.HandlerFunc,
	middlewareList []func(http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	for _, middleware := range middlewareList {
		handler = middleware(handler)
	}

	return handler
}

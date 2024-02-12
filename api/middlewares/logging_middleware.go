package middlewares

import (
	"log"
	"net/http"
)

type responseLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResponseLoggingWriter(w http.ResponseWriter) *responseLoggingWriter {
	return &responseLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rlw *responseLoggingWriter) WriteHeader(code int) {
	rlw.code = code
	rlw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		rlw := NewResponseLoggingWriter(w)
		next(rlw, r)
		log.Println("status_code", rlw.code)
	}
}

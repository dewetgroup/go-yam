package yam

import (
	"net/http"
)

func CreateMockMiddleware(control *int) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			*control = 1
			next.ServeHTTP(writer, request)
		})
	}
}

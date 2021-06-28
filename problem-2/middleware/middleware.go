package middleware

import (
	"net/http"

	"github.com/aditya-nagare/take-home-test/problem-2/log"
)

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req = req.WithContext(log.WithRqID(req.Context()))

		next.ServeHTTP(w, req)
	})
}

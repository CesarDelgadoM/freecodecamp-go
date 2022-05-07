package middleware

import "net/http"

func SetMiddlewareJson(hf http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("content-type", "application/json")
		hf(rw, r)
	}
}

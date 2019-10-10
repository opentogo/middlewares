package middlewares

import (
	"net/http"
)

func FrameOptions(option string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		frameOptions := r.Header.Get(headerFrameOptions)
		w.Header().Set(headerFrameOptions, frameOptions)

		if frameOptions != "" {
			next.ServeHTTP(w, r)
			return
		}
		for _, contentType := range htmlContentTypes {
			if contentType == r.Header.Get(headerContentType) {
				w.Header().Set(headerFrameOptions, option)
				break
			}
		}
		next.ServeHTTP(w, r)
	}
}

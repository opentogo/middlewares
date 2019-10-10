package middlewares

import (
	"fmt"
	"net/http"
)

func XSS(mode string, nosniff bool, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			contentOptions = r.Header.Get(headerContentTypeOptions)
			xss            = r.Header.Get(headerXSSProtection)
		)
		w.Header().Set(headerContentTypeOptions, contentOptions)
		w.Header().Set(headerXSSProtection, xss)

		if xss == "" {
			for _, contentType := range htmlContentTypes {
				if contentType == r.Header.Get(headerContentType) {
					w.Header().Set(headerXSSProtection, fmt.Sprintf("1; mode=%s", mode))
					break
				}
			}
		}
		if contentOptions == "" && nosniff {
			w.Header().Set(headerContentTypeOptions, "nosniff")
		}
		next.ServeHTTP(w, r)
	}
}

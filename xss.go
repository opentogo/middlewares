package middlewares

import (
	"fmt"
	"net/http"
)

type XSS struct {
	mode    string
	nosniff bool
}

func NewXSS(mode string, nosniff bool) XSS {
	return XSS{
		mode:    mode,
		nosniff: nosniff,
	}
}

func (m XSS) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentTypeOptions, r.Header.Get(headerContentTypeOptions))
		w.Header().Set(headerXSSProtection, r.Header.Get(headerXSSProtection))

		if m.isNotEmpty(r) {
			next.ServeHTTP(w, r)
			return
		}
		for _, contentType := range htmlContentTypes {
			if contentType == r.Header.Get(headerContentType) {
				w.Header().Set(headerXSSProtection, fmt.Sprintf("1; mode=%s", m.mode))
				break
			}
		}
		if m.nosniff {
			w.Header().Set(headerContentTypeOptions, "nosniff")
		}

		next.ServeHTTP(w, r)
	}
}

func (m XSS) isNotEmpty(r *http.Request) bool {
	return r.Header.Get(headerContentTypeOptions) != "" && r.Header.Get(headerXSSProtection) != ""
}

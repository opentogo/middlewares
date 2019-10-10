package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

func TestXSS(t *testing.T) {
	var (
		r = httptest.NewRequest(http.MethodGet, "/", nil)
		w = httptest.NewRecorder()
		h = func(w http.ResponseWriter, r *http.Request) {}
	)

	t.Run("setting 'X-XSS-Protection' header for HTML content-types", func(t *testing.T) {
		for _, contentType := range htmlContentTypes {
			w.Header().Del(headerContentTypeOptions)
			w.Header().Del(headerXSSProtection)

			t.Run(contentType, func(t *testing.T) {
				r.Header.Set(headerContentType, contentType)
				http.HandlerFunc(XSS("block", true, h)).ServeHTTP(w, r)

				assert.Equal(t, "1; mode=block", w.Header().Get(headerXSSProtection))
			})
		}
	})

	t.Run("setting 'X-XSS-Protection' header for JSON content-type", func(t *testing.T) {
		w.Header().Del(headerContentTypeOptions)
		w.Header().Del(headerXSSProtection)

		r.Header.Set(headerContentType, "application/json")
		http.HandlerFunc(XSS("block", true, h)).ServeHTTP(w, r)

		assert.Equal(t, "", w.Header().Get(headerXSSProtection))
	})

	t.Run("setting 'X-Content-Type-Options' header", func(t *testing.T) {
		w.Header().Del(headerContentTypeOptions)
		w.Header().Del(headerXSSProtection)

		r.Header.Set(headerContentType, "text/html")

		http.HandlerFunc(XSS("block", true, h)).ServeHTTP(w, r)

		assert.Equal(t, "nosniff", w.Header().Get(headerContentTypeOptions))
	})

	t.Run("checking override the 'X-XSS-Protection' header", func(t *testing.T) {
		w.Header().Del(headerContentTypeOptions)
		w.Header().Del(headerXSSProtection)

		r.Header.Set(headerContentTypeOptions, "sniff")
		r.Header.Set(headerXSSProtection, "1; mode=foo")
		r.Header.Set(headerContentType, "text/html")

		http.HandlerFunc(XSS("block", true, h)).ServeHTTP(w, r)

		assert.Equal(t, "sniff", w.Header().Get(headerContentTypeOptions))
		assert.Equal(t, "1; mode=foo", w.Header().Get(headerXSSProtection))
	})
}

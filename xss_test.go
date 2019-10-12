package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

func TestXSS(t *testing.T) {
	var (
		res = httptest.NewRecorder()
		xss = NewXSS("block", true)
	)

	t.Run("setting 'X-XSS-Protection' header for HTML content-types", func(t *testing.T) {
		for _, contentType := range htmlContentTypes {
			res.Header().Del(headerContentTypeOptions)
			res.Header().Del(headerXSSProtection)

			t.Run(contentType, func(t *testing.T) {
				req.Header.Set(headerContentType, contentType)

				http.HandlerFunc(xss.Handler(testHandler)).ServeHTTP(res, req)

				assert.Equal(t, "1; mode=block", res.Header().Get(headerXSSProtection))
			})
		}
	})

	t.Run("setting 'X-XSS-Protection' header for JSON content-type", func(t *testing.T) {
		res.Header().Del(headerContentTypeOptions)
		res.Header().Del(headerXSSProtection)

		req.Header.Set(headerContentType, "application/json")
		http.HandlerFunc(xss.Handler(testHandler)).ServeHTTP(res, req)

		assert.Equal(t, "", res.Header().Get(headerXSSProtection))
	})

	t.Run("setting 'X-Content-Type-Options' header", func(t *testing.T) {
		res.Header().Del(headerContentTypeOptions)
		res.Header().Del(headerXSSProtection)

		req.Header.Set(headerContentType, "text/html")

		http.HandlerFunc(xss.Handler(testHandler)).ServeHTTP(res, req)

		assert.Equal(t, "nosniff", res.Header().Get(headerContentTypeOptions))
	})

	t.Run("checking override the 'X-XSS-Protection' header", func(t *testing.T) {
		res.Header().Del(headerContentTypeOptions)
		res.Header().Del(headerXSSProtection)

		req.Header.Set(headerContentTypeOptions, "sniff")
		req.Header.Set(headerXSSProtection, "1; mode=foo")
		req.Header.Set(headerContentType, "text/html")

		http.HandlerFunc(xss.Handler(testHandler)).ServeHTTP(res, req)

		assert.Equal(t, "sniff", res.Header().Get(headerContentTypeOptions))
		assert.Equal(t, "1; mode=foo", res.Header().Get(headerXSSProtection))
	})
}

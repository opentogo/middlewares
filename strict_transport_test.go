package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

func TestStrictTransport(t *testing.T) {
	res := httptest.NewRecorder()

	t.Run("setting 'Strict-Transport-Security' header", func(t *testing.T) {
		m := NewStrictTransport(31536000, false, false)

		http.HandlerFunc(m.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, "max-age=31536000", res.Header().Get(headerStrictTransport))
	})

	t.Run("switching on the include_subdomains option", func(t *testing.T) {
		m := NewStrictTransport(31536000, true, false)

		http.HandlerFunc(m.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, "max-age=31536000; includeSubDomains", res.Header().Get(headerStrictTransport))
	})

	t.Run("switching on the preload option", func(t *testing.T) {
		m := NewStrictTransport(31536000, false, true)

		http.HandlerFunc(m.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, "max-age=31536000; preload", res.Header().Get(headerStrictTransport))
	})

	t.Run("should allow switching on all the options", func(t *testing.T) {
		m := NewStrictTransport(31536000, true, true)

		http.HandlerFunc(m.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, "max-age=31536000; includeSubDomains; preload", res.Header().Get(headerStrictTransport))
	})
}

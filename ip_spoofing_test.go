package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

func TestIPSpoofing(t *testing.T) {
	var (
		res        = httptest.NewRecorder()
		ipSpoofing = NewIPSpoofing()
	)

	t.Run("accepts requests without 'X-Forward-For' header", func(t *testing.T) {
		http.HandlerFunc(ipSpoofing.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("accepts requests with proper 'X-Forward-For' header", func(t *testing.T) {
		req.Header.Set(headerClientIP, "1.2.3.4")
		req.Header.Set(headerForwardedFor, "192.168.1.20, 1.2.3.4, 127.0.0.1")

		http.HandlerFunc(ipSpoofing.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("denies requests where the client spoofs 'X-Forward-For' but not the IP", func(t *testing.T) {
		req.Header.Set(headerClientIP, "1.2.3.4")
		req.Header.Set(headerForwardedFor, "1.2.3.5")

		http.HandlerFunc(ipSpoofing.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("denies requests where the client spoofs the IP but not X-Forward-For", func(t *testing.T) {
		req.Header.Set(headerClientIP, "1.2.3.5")
		req.Header.Set(headerForwardedFor, "192.168.1.20, 1.2.3.4, 127.0.0.1")

		http.HandlerFunc(ipSpoofing.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("denies requests where IP and X-Forward-For are spoofed but not X-Real-IP", func(t *testing.T) {
		req.Header.Set(headerClientIP, "1.2.3.5")
		req.Header.Set(headerForwardedFor, "1.2.3.5")
		req.Header.Set(headerRealIP, "1.2.3.4")

		http.HandlerFunc(ipSpoofing.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusForbidden, res.Code)
	})
}

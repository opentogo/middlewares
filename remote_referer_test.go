package middlewares

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/allisson/go-assert"
)

func TestRemoteReferer(t *testing.T) {
	var (
		err           error
		res           = httptest.NewRecorder()
		remoteReferer = NewRemoteReferer(nil)
	)

	t.Run("accepts POST requests without remote referer", func(t *testing.T) {
		req.Method = http.MethodPost

		http.HandlerFunc(remoteReferer.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("accepts GET requests with remote referer", func(t *testing.T) {
		req.URL, err = url.Parse("http://example.org")
		assert.Nil(t, err)

		req.Method = http.MethodGet
		req.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(remoteReferer.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("accepts POST requests with different hosts", func(t *testing.T) {
		req.URL, err = url.Parse("http://foo.org")
		assert.Nil(t, err)

		req.Method = http.MethodPost
		req.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(remoteReferer.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("accepts POST requests with same referer filtering only GET requests", func(t *testing.T) {
		req.URL, err = url.Parse("http://example.org")
		assert.Nil(t, err)

		req.Method = http.MethodPost
		req.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(NewRemoteReferer([]string{http.MethodGet}).Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("denies POST requests with remote referer", func(t *testing.T) {
		req.URL, err = url.Parse("http://example.org")
		assert.Nil(t, err)

		req.Method = http.MethodPost
		req.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(remoteReferer.Handler(testHandler)).ServeHTTP(res, req)
		assert.Equal(t, http.StatusForbidden, res.Code)
	})
}

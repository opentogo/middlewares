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
		r             = httptest.NewRequest(http.MethodGet, "/", nil)
		w             = httptest.NewRecorder()
		handler       = func(w http.ResponseWriter, r *http.Request) {}
		remoteReferer = NewRemoteReferer(nil)
	)

	t.Run("accepts POST requests without remote referer", func(t *testing.T) {
		r.Method = http.MethodPost

		http.HandlerFunc(remoteReferer.Handler(handler)).ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("accepts GET requests with remote referer", func(t *testing.T) {
		r.URL, err = url.Parse("http://example.org")
		assert.Nil(t, err)

		r.Method = http.MethodGet
		r.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(remoteReferer.Handler(handler)).ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("accepts POST requests with different hosts", func(t *testing.T) {
		r.URL, err = url.Parse("http://foo.org")
		assert.Nil(t, err)

		r.Method = http.MethodPost
		r.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(remoteReferer.Handler(handler)).ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("accepts POST requests with same referer filtering only GET requests", func(t *testing.T) {
		r.URL, err = url.Parse("http://example.org")
		assert.Nil(t, err)

		r.Method = http.MethodPost
		r.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(NewRemoteReferer([]string{http.MethodGet}).Handler(handler)).ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("denies POST requests with remote referer", func(t *testing.T) {
		r.URL, err = url.Parse("http://example.org")
		assert.Nil(t, err)

		r.Method = http.MethodPost
		r.Header.Set(headerRemoteReferer, "http://example.org")

		http.HandlerFunc(remoteReferer.Handler(handler)).ServeHTTP(w, r)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}

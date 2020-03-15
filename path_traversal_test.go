package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

func TestPathTraversal(t *testing.T) {
	t.Run("does not touch /foo/bar", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/foo/bar")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/foo/bar", req.Header.Get(headerPathInfo))
	})

	t.Run("does not touch /", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/", req.Header.Get(headerPathInfo))
	})

	t.Run("does not touch /.f", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/.f")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/.f", req.Header.Get(headerPathInfo))
	})

	t.Run("does not touch /a.x", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/a.x")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/a.x", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /.. with /", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/..")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /a/../b with /b", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/a/../b")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/b", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /a/../b/ with /b/", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/a/../b/")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/b/", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /a/. with /a/", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/a/.")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/a/", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /%2e. with /", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/%2e.")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /a/%2E%2e/b with /b", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/a/%2E%2e/b")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/b", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /a%2f%2E%2e%2Fb/ with /b/", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/a%2f%2E%2e%2Fb/")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/b/", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces // with /", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "//")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/", req.Header.Get(headerPathInfo))
	})

	t.Run("replaces /%2fetc%2Fpasswd with /etc/passwd", func(t *testing.T) {
		req.Header.Set(headerPathInfo, "/%2fetc%2Fpasswd")

		http.HandlerFunc(NewPathTraversal().Handler(testHandler)).ServeHTTP(httptest.NewRecorder(), req)
		assert.Equal(t, "/etc/passwd", req.Header.Get(headerPathInfo))
	})
}

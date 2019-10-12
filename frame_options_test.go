package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

var (
	req         = httptest.NewRequest(http.MethodGet, "/", nil)
	testHandler = func(w http.ResponseWriter, r *http.Request) {}
)

func TestFrameOptions(t *testing.T) {
	var (
		res          = httptest.NewRecorder()
		frameOptions = NewFrameOptions("SAMEORIGIN")
	)

	t.Run("setting 'X-Frame-Options' header for HTML content-types", func(t *testing.T) {
		for _, contentType := range htmlContentTypes {
			res.Header().Del(headerFrameOptions)

			t.Run(contentType, func(t *testing.T) {
				req.Header.Set(headerContentType, contentType)
				http.HandlerFunc(frameOptions.Handler(testHandler)).ServeHTTP(res, req)

				assert.Equal(t, "SAMEORIGIN", res.Header().Get(headerFrameOptions))
			})
		}
	})

	t.Run("setting 'X-Frame-Options' header for JSON content-type", func(t *testing.T) {
		res.Header().Del(headerFrameOptions)

		req.Header.Set(headerContentType, "application/json")
		http.HandlerFunc(frameOptions.Handler(testHandler)).ServeHTTP(res, req)

		assert.Equal(t, "", res.Header().Get(headerFrameOptions))
	})

	t.Run("checking override the 'X-Frame-Options' header", func(t *testing.T) {
		res.Header().Del(headerFrameOptions)

		req.Header.Set(headerFrameOptions, "ALLOW")
		req.Header.Set(headerContentType, "text/html")

		http.HandlerFunc(frameOptions.Handler(testHandler)).ServeHTTP(res, req)

		assert.Equal(t, "ALLOW", res.Header().Get(headerFrameOptions))
	})
}

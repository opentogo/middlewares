package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

func TestFrameOptions(t *testing.T) {
	var (
		r = httptest.NewRequest(http.MethodGet, "/", nil)
		w = httptest.NewRecorder()
		h = func(w http.ResponseWriter, r *http.Request) {}
	)

	t.Run("setting 'X-Frame-Options' header for HTML content-types", func(t *testing.T) {
		for _, contentType := range htmlContentTypes {
			w.Header().Del(headerFrameOptions)

			t.Run(contentType, func(t *testing.T) {
				r.Header.Set(headerContentType, contentType)
				http.HandlerFunc(FrameOptions("SAMEORIGIN", h)).ServeHTTP(w, r)

				assert.Equal(t, "SAMEORIGIN", w.Header().Get(headerFrameOptions))
			})
		}
	})

	t.Run("setting 'X-Frame-Options' header for JSON content-type", func(t *testing.T) {
		w.Header().Del(headerFrameOptions)

		r.Header.Set(headerContentType, "application/json")
		http.HandlerFunc(FrameOptions("SAMEORIGIN", h)).ServeHTTP(w, r)

		assert.Equal(t, "", w.Header().Get(headerFrameOptions))
	})

	t.Run("checking override the 'X-Frame-Options' header", func(t *testing.T) {
		w.Header().Del(headerFrameOptions)

		r.Header.Set(headerFrameOptions, "ALLOW")
		r.Header.Set(headerContentType, "text/html")

		http.HandlerFunc(FrameOptions("SAMEORIGIN", h)).ServeHTTP(w, r)

		assert.Equal(t, "ALLOW", w.Header().Get(headerFrameOptions))
	})
}

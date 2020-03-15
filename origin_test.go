package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allisson/go-assert"
)

func TestOrigin(t *testing.T) {
	origin := NewOrigin(nil)

	for _, method := range []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	} {
		t.Run(fmt.Sprintf("accepts %q requests with no origin", method), func(t *testing.T) {
			res := httptest.NewRecorder()
			req.Method = method

			http.HandlerFunc(origin.Handler(testHandler)).ServeHTTP(res, req)
			assert.Equal(t, http.StatusOK, res.Code)
		})
	}

	for _, method := range []string{
		http.MethodGet,
		http.MethodHead,
	} {
		t.Run(fmt.Sprintf("accepts %q requests with non-whitelisted origin", method), func(t *testing.T) {
			res := httptest.NewRecorder()

			req.Method = method
			req.Header.Set(headerOrigin, "http://malicious.com")

			http.HandlerFunc(origin.Handler(testHandler)).ServeHTTP(res, req)
			assert.Equal(t, http.StatusOK, res.Code)
		})
	}

	for _, method := range []string{
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	} {
		t.Run(fmt.Sprintf("denies %q requests with non-whitelisted origin", method), func(t *testing.T) {
			res := httptest.NewRecorder()

			req.Method = method
			req.Header.Set(headerOrigin, "http://malicious.com")

			http.HandlerFunc(origin.Handler(testHandler)).ServeHTTP(res, req)
			assert.Equal(t, http.StatusForbidden, res.Code)
		})

		t.Run(fmt.Sprintf("accepts %q requests with whitelisted origin", method), func(t *testing.T) {
			res := httptest.NewRecorder()

			req.Method = method
			req.Header.Set(headerOrigin, "http://friend.com")

			m := NewOrigin([]string{"http://friend.com"})

			http.HandlerFunc(m.Handler(testHandler)).ServeHTTP(res, req)
			assert.Equal(t, http.StatusOK, res.Code)
		})
	}
}

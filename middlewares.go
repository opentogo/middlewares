package middlewares

import "net/http"

const (
	headerContentLength      = "Content-Length"
	headerContentType        = "Content-Type"
	headerContentTypeOptions = "X-Content-Type-Options"
	headerFrameOptions       = "X-Frame-Options"
	headerRemoteReferer      = "Referer"
	headerXSSProtection      = "X-XSS-Protection"
)

var htmlContentTypes = []string{
	"text/html",
	"text/html;charset=utf-8",
	"application/xhtml",
	"application/xhtml+xml",
}

var safeMethods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodOptions,
	http.MethodTrace,
}

type Middleware interface {
	Handler(http.HandlerFunc) http.HandlerFunc
}

func NewMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware.Handler(handler)
	}
	return handler
}

/*
handler := NewMiddleware(
	handler,
	NewXSS("block", true),
	NewFrameOptions("SAMEORIGIN"),
	NewRemoteReferer([]string{http.MethodGet})
)
*/

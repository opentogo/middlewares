package middlewares

import "net/http"

const (
	headerClientIP           = "X-Client-IP"
	headerContentLength      = "Content-Length"
	headerContentType        = "Content-Type"
	headerContentTypeOptions = "X-Content-Type-Options"
	headerForwardedFor       = "X-Forwarded-For"
	headerFrameOptions       = "X-Frame-Options"
	headerPathInfo           = "PATH_INFO"
	headerRealIP             = "X-Real-IP"
	headerRemoteReferer      = "Referer"
	headerStrictTransport    = "Strict-Transport-Security"
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
	NewFrameOptions("SAMEORIGIN"),
	NewIPSpoofing(),
	NewPathTraversal(),
	NewRemoteReferer([]string{http.MethodGet})
	NewStrictTransport(31536000, false, false),
	NewXSS("block", true),
)
*/

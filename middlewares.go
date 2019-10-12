package middlewares

import "net/http"

const (
	headerClientIP           = "X-Client-IP"
	headerContentLength      = "Content-Length"
	headerContentType        = "Content-Type"
	headerContentTypeOptions = "X-Content-Type-Options"
	headerForwardedFor       = "X-Forwarded-For"
	headerFrameOptions       = "X-Frame-Options"
	headerOrigin             = "Origin"
	headerOriginHTTP         = "HTTP_ORIGIN"
	headerOriginXHTTP        = "HTTP_X_ORIGIN"
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

var middlewares []Middleware

type Middleware interface {
	Handler(http.HandlerFunc) http.HandlerFunc
}

func Use(middleware ...Middleware) {
	if len(middleware) > 0 {
		middlewares = append(middlewares, middleware...)
	}
}

func Handle(handler http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware.Handler(handler)
	}
	return handler
}

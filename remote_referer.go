package middlewares

import (
	"net/http"
	"net/url"
	"strings"
)

// RemoteReferer type. It doesn't accept unsafe HTTP requests if the Referer
// header is set to a different host
type RemoteReferer struct {
	methods []string
}

// NewRemoteReferer creates new instance of RemoteReferer
func NewRemoteReferer(methods []string) RemoteReferer {
	return RemoteReferer{
		methods: methods,
	}
}

// Handler checks if HTTP request is safe.
func (m RemoteReferer) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if m.isURLSafe(r) || m.isMethodFiltered(r) || m.isMethodSafe(r) {
			next.ServeHTTP(w, r)
			return
		}
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(headerContentType, "text/plain")
			w.Header().Set(headerContentLength, "0")
			w.WriteHeader(http.StatusForbidden)
		}).ServeHTTP(w, r)
	}
}

func (m RemoteReferer) isURLSafe(r *http.Request) bool {
	referer, err := url.Parse(r.Header.Get(headerRemoteReferer))
	return err != nil || referer.Host == "" || referer.Host != r.URL.Host
}

func (m RemoteReferer) isMethodFiltered(r *http.Request) bool {
	return m.methods != nil && len(m.methods) > 0 && !strings.Contains(strings.Join(m.methods, ","), r.Method)
}

func (m RemoteReferer) isMethodSafe(r *http.Request) bool {
	return strings.Contains(strings.Join(safeMethods, ","), r.Method)
}

package middlewares

import (
	"net/http"
	"strings"
)

// IPSpoofing type used to detect IP spoofing attacks.
type IPSpoofing struct{}

// NewIPSpoofing creates a new instance of IPSpoofing
func NewIPSpoofing() IPSpoofing {
	return IPSpoofing{}
}

// Handler checks if request of safe against IPSpoofing attacks
func (m IPSpoofing) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			clientIP     = r.Header.Get(headerClientIP)
			forwardedFor = r.Header.Get(headerForwardedFor)
			realIP       = r.Header.Get(headerRealIP)
			isSafe       = false
		)
		if clientIP == "" && realIP == "" {
			next.ServeHTTP(w, r)
			return
		}
		for _, ip := range strings.Split(forwardedFor, ", ") {
			if ip == clientIP || ip == realIP {
				isSafe = true
				break
			}
		}
		if !isSafe {
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(headerContentType, "text/plain")
				w.Header().Set(headerContentLength, "0")
				w.WriteHeader(http.StatusForbidden)
			}).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}
}

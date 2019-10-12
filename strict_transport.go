package middlewares

import (
	"fmt"
	"net/http"
)

type StrictTransport struct {
	maxAge            int
	includeSubdomains bool
	preload           bool
}

func NewStrictTransport(maxAge int, includeSubdomains, preload bool) StrictTransport {
	return StrictTransport{
		maxAge:            maxAge,
		includeSubdomains: includeSubdomains,
		preload:           preload,
	}
}

func (m StrictTransport) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(
			headerStrictTransport,
			fmt.Sprintf(
				"max-age=%d%s%s",
				m.maxAge,
				m.subdomainsOption(),
				m.preloadOption(),
			),
		)
		next.ServeHTTP(w, r)
	}
}

func (m StrictTransport) subdomainsOption() (option string) {
	if m.includeSubdomains {
		option = "; includeSubDomains"
	}
	return
}

func (m StrictTransport) preloadOption() (option string) {
	if m.preload {
		option = "; preload"
	}
	return
}

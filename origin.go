package middlewares

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Origin struct {
	whitelist []string
}

func NewOrigin(whitelist []string) Origin {
	return Origin{
		whitelist: whitelist,
	}
}

func (m Origin) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var origin string
		for _, value := range []string{
			r.Header.Get(headerOrigin),
			r.Header.Get(headerOriginHTTP),
			r.Header.Get(headerOriginXHTTP),
		} {
			if value != "" {
				origin = value
			}
		}
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}
		if strings.Contains(strings.Join(safeMethods, ","), r.Method) {
			next.ServeHTTP(w, r)
			return
		}
		if baseURL(origin) == origin {
			next.ServeHTTP(w, r)
			return
		}
		if m.whitelist != nil {
			for _, item := range m.whitelist {
				if item == origin {
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(headerContentType, "text/plain")
			w.Header().Set(headerContentLength, "0")
			w.WriteHeader(http.StatusForbidden)
		}).ServeHTTP(w, r)
	}
}

func baseURL(origin string) (baseURL string) {
	var (
		err    error
		parsed *url.URL
	)
	if parsed, err = url.Parse(origin); err != nil {
		return
	}
	if parsed.Port() != "80" || parsed.Port() != "443" {
		baseURL = fmt.Sprintf(":%s", parsed.Port())
	}
	baseURL = fmt.Sprintf("%s://%s%s", parsed.Scheme, parsed.Host, baseURL)
	return
}

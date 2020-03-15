package middlewares

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// PathTraversal type. It protects against unauthorized access to file system attacks,
// unescapes '/' and '.' from PATH_INFO
type PathTraversal struct{}

// NewPathTraversal create new instance of PathTraversal
func NewPathTraversal() PathTraversal {
	return PathTraversal{}
}

// Handler checks for url if file path is being accessed.
func (m PathTraversal) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set(headerPathInfo, m.cleaner(r.Header.Get(headerPathInfo)))
		next.ServeHTTP(w, r)
	}
}

func (m PathTraversal) cleaner(path string) (cleanPath string) {
	var (
		err   error
		parts []string
	)
	if path, err = url.PathUnescape(path); err != nil {
		return
	}
	path = strings.Replace(path, `\\`, "/", -1)
	for _, part := range strings.Split(path, "/") {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
			continue
		}
		parts = append(parts, part)
	}

	cleanPath = fmt.Sprintf("/%s", strings.Join(parts, "/"))
	if len(parts) > 0 && regexp.MustCompile(`\/\.{0,2}$`).MatchString(path) {
		cleanPath += "/"
	}
	return
}

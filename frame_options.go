package middlewares

import "net/http"

// FrameOptions Defines who should be allowed to embed the page in a frame.
// It protects against clickjacking, setting header to tell the browser
// avoid embedding the page in a frame.
type FrameOptions struct {
	option string
}

// NewFrameOptions create new instance of FrameOptions
func NewFrameOptions(option string) FrameOptions {
	return FrameOptions{
		option: option,
	}
}

// Handler adds frameOptions to response header and returns http.HandlerFunc
func (m FrameOptions) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		frameOptions := r.Header.Get(headerFrameOptions)
		if frameOptions != "" {
			w.Header().Set(headerFrameOptions, frameOptions)

			next.ServeHTTP(w, r)
			return
		}
		for _, contentType := range htmlContentTypes {
			if contentType == r.Header.Get(headerContentType) {
				w.Header().Set(headerFrameOptions, m.option)
				break
			}
		}
		next.ServeHTTP(w, r)
	}
}

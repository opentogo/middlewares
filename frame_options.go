package middlewares

import "net/http"

type FrameOptions struct {
	option string
}

func NewFrameOptions(option string) FrameOptions {
	return FrameOptions{
		option: option,
	}
}

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

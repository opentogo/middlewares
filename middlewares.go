package middlewares

const (
	headerContentType        = "Content-Type"
	headerContentTypeOptions = "X-Content-Type-Options"
	headerFrameOptions       = "X-Frame-Options"
	headerXSSProtection      = "X-XSS-Protection"
)

var htmlContentTypes = []string{
	"text/html",
	"text/html;charset=utf-8",
	"application/xhtml",
	"application/xhtml+xml",
}

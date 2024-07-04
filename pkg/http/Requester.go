package http

import "net/http"

type Rest interface {
	GET() (*http.Response, error)
}

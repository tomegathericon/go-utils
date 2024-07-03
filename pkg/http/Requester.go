package http

import "net/http"

type Rester interface {
	GET() (*http.Response, error)
}

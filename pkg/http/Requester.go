package http

import "net/http"

type Requester interface {
	GET() (*http.Response, error)
}

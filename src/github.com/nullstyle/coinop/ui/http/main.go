package http

import (
	"golang.org/x/net/context"
	"net/http"
)

// UI represents an http "user interface", that is, an API.
type UI struct {
	Ctx context.Context
	R   *http.Request
	W   http.ResponseWriter
}

package http

import (
	"net/http"
)

// Driver represents an http clienta
type Driver struct {
	HTTP *http.Client `inject:""`
}

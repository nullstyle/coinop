package http

import (
	"net/http"
)

type Horizon struct {
	BaseURL string       `inject:"horizon-url"`
	Client  *http.Client `inject:""`
}

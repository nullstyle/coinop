package http

import (
	"fmt"
	"net/http"
)

// DeliveryError represents a failed http request
type DeliveryError struct {
	URL  string
	Resp *http.Response
}

func (err *DeliveryError) Error() string {
	return fmt.Sprintf("failed call to %s: got %d", err.URL, err.Resp.StatusCode)
}

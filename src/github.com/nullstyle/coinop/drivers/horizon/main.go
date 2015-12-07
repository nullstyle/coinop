package horizon

import (
	"net/http"
)

type Driver struct {
	BaseURL string       `inject:"horizon-url"`
	HTTP    *http.Client `inject:""`
}

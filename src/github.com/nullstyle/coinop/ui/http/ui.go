package http

import (
	"encoding/json"
)

func (ui *UI) JSON(v interface{}) error {
	ui.W.Header().Set("Content-Type", "application/json; charset=UTF-8")

	return json.NewEncoder(ui.W).Encode(v)
}

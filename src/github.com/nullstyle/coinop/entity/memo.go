package entity

import (
	"errors"
)

// Valid returns nil if the memo is valid
func (m *Memo) Valid() error {
	switch m.Type {
	case "none":
		if m.Value != "" {
			return errors.New("invlid memo: `none` type can't have value")
		}
	case "text", "id", "hash", "return":
		return nil
	}
	return errors.New("invalid memo type: " + m.Type)
}

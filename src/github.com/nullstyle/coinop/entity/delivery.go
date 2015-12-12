package entity

import (
	"errors"
)

// Valid returns true if the delivery is valid
func (d *Delivery) Valid() error {
	if d.URL == nil {
		return errors.New("invalid delivery: empty url")
	}

	return d.Payment.Valid()
}

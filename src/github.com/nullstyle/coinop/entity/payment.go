package entity

import (
	"errors"
	"fmt"
)

// Valid returns true if the payment is valid
func (d *Payment) Valid() error {
	if d.PagingToken == "" {
		return errors.New("invalid payment: empty paging token")
	}

	if d.From == AccountID("") {
		return errors.New("invalid payment: from is empty")
	}

	if d.To == AccountID("") {
		return errors.New("invalid payment: to is empty")
	}

	if err := d.Memo.Valid(); err != nil {
		return fmt.Errorf("invalid payment: %s", err)
	}

	if d.Asset == (Asset{}) {
		return errors.New("invalid payment: zero-value asset")
	}

	if d.Amount == (Amount{}) {
		return errors.New("invalid payment: zero-value amount")
	}

	return nil
}

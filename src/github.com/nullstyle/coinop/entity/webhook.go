package entity

import (
	"errors"
	"github.com/stellar/go-stellar-base/strkey"
)

// Valid returns true if the webhook is valid
func (hook *Webhook) Valid() error {
	if hook.URL == nil {
		return errors.New("invalid webhook: empty url")
	}

	_, err := strkey.Decode(strkey.VersionByteAccountID, hook.DestinationFilter)
	if err != nil {
		return errors.New("invalid webhook: bad destination filter")
	}

	if hook.MemoFilter == nil {
		return nil
	}

	return hook.MemoFilter.Valid()
}

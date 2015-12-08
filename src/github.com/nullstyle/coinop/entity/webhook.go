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

	_, err := strkey.Decode(
		strkey.VersionByteAccountID,
		string(hook.DestinationFilter),
	)
	if err != nil {
		return errors.New("invalid webhook: bad destination filter")
	}

	if hook.MemoFilter == nil {
		return nil
	}

	return hook.MemoFilter.Valid()
}

// IsTriggeredBy returns true if `p` matches `hook`, signifying it should
// trigger a delivery.
func (hook *Webhook) IsTriggeredBy(p Payment) bool {
	if hook.DestinationFilter != p.To {
		return false
	}

	if hook.MemoFilter != nil && *hook.MemoFilter != p.Memo {
		return false
	}

	return true
}

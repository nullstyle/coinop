package horizon

import (
	"github.com/nullstyle/coinop/usecase"
)

// StreamPayments streams payments into `fn`, starting at `cursor` by listening
// to a horizon instance.
func (driver *Driver) StreamPayments(
	cursor string,
	fn usecase.PaymentHandler,
) error {
	return nil
}

var _ usecase.PaymentProvider = &Driver{}

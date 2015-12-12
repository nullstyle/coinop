package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

// WatchPayments watches for payments on a digital payment network and triggers
// the ProcessPayment subprocess in response
type WatchPayments struct {
	Payments  PaymentProvider
	Processor paymentProcessor
}

// Exec runs the use case.
func (kase *WatchPayments) Exec(cursor string) (err error) {
	if cursor == "" {
		cursor = "now"
	}

	err = kase.Payments.StreamPayments(cursor, func(p entity.Payment) error {
		return kase.Processor.Exec(p)
	})
	return
}

type paymentProcessor interface {
	Exec(p entity.Payment) (err error)
}

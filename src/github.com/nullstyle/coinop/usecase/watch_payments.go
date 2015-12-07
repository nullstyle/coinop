package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

// WatchPayments watches for payments on a digital payment network and triggers
// the ProcessPayment subprocess in response
type WatchPayments struct {
	DB        PaymentRepository
	Payments  PaymentProvider
	Processor ProcessPayment
}

// Exec runs the use case.
func (kase *WatchPayments) Exec() (err error) {
	var cursor string
	cursor, err = kase.DB.LoadCursor()
	if err != nil {
		return
	}

	if cursor == "" {
		cursor = "now"
	}

	err = kase.Payments.StreamPayments(cursor, func(p entity.Payment) error {
		err := kase.Processor.Exec(p)
		if err != nil {
			return err
		}

		return kase.DB.SaveCursor(p.PagingToken)
	})
	return
}

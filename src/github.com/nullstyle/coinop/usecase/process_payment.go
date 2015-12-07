package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

// ProcessPayment receives a payment and creates a new delivery for each
// matching hook
type ProcessPayment struct {
	Hooks      WebhookRepository
	Deliveries DeliveryRepository
}

// Exec runs the use case.
func (kase *ProcessPayment) Exec(p entity.Payment) (err error) {
	return
}

package fake

import (
	"github.com/nullstyle/coinop/entity"
)

// DeliveryEntity returns a new random valid delivery
func DeliveryEntity() (result entity.Delivery) {
	result.URL = WebhookEntity().URL
	result.Payment = PaymentEntity()

	if err := result.Valid(); err != nil {
		panic(err)
	}

	return
}

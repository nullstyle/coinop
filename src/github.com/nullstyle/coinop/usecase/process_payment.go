package usecase

import (
	"log"

	"github.com/nullstyle/coinop/entity"
	"golang.org/x/net/context"
)

// ProcessPayment receives a payment and creates a new delivery for each
// matching hook
type ProcessPayment struct {
	Hooks      WebhookRepository
	Deliveries DeliveryRepository
	Sender     DeliverySender
}

// Exec runs the use case.
func (kase *ProcessPayment) Exec(p entity.Payment) (err error) {
	hooks, err := kase.Hooks.ForDestination(p.To)
	if err != nil {
		return
	}

	ds := []entity.Delivery{}
	for _, hook := range hooks {
		if !hook.IsTriggeredBy(p) {
			continue
		}
		d := entity.Delivery{URL: hook.URL, Payment: p}
		ds = append(ds, d)
	}

	err = kase.Deliveries.SaveDeliveries(p.PagingToken, ds)
	if err != nil {
		return
	}

	for _, d := range ds {
		go kase.deliver(d)
	}

	return
}

func (kase *ProcessPayment) deliver(d entity.Delivery) {
	pd := ProcessDelivery{DB: kase.Deliveries, Sender: kase.Sender}
	err := pd.Exec(context.Background(), d)
	if err != nil {
		log.Printf("warn: failed delivery: %s", err)
	}
}

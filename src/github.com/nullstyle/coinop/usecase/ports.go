package usecase

import (
	"github.com/nullstyle/coinop/entity"
	"time"
)

// DeliveryRepository tracks delivery state in persistent storage and
// coordinates updates to that state.
type DeliveryRepository interface {
	// StartDelivery indicates to the repository that the caller wants to perform
	// the provided delivery.  An implementation should ensure that the token
	// returned will authorize an update to the delivery during a window of
	// exclusive access to the delivery bounded by the returned deadline.
	StartDelivery(entity.Delivery) (token int64, deadline time.Time, err error)

	// MarkDeliverySuccess marks the provided delivery as successfully completed.
	MarkDeliverySuccess(token int64, delivery entity.Delivery) error

	// MarkDeliveryFailed marks the provided delivery as a failure, provided it
	// has not been marked as successful or failed in the past.
	MarkDeliveryFailed(token int64, delivery entity.Delivery) error

	// FailedDeliveries returns a slice of deliveries that have been marked as
	// failed.
	FailedDeliveries() ([]entity.Delivery, error)
}

// DeliverySender makes http requests against urls with payment information
type DeliverySender interface {
	SendDelivery(entity.Delivery) error
}

// PaymentHandler represents a function that responds to a payment
type PaymentHandler func(entity.Payment) error

// PaymentProvider streams payments into the system
type PaymentProvider interface {
	StreamPayments(cursor string, fn PaymentHandler) error
}

// PaymentRepository saves the last seen cursor from a PaymentProvider
type PaymentRepository interface {
	SaveCursor(string) error
	LoadCursor() (string, error)
}

// WebhookPresenter displays webhooks
type WebhookPresenter interface {
	List([]entity.Webhook) error
}

// WebhookRepository saves webhooks in persistent storage
type WebhookRepository interface {
	SaveWebhook(*entity.Webhook) error
	DestroyWebhook(ID RepoID) error
	ListWebhooks() ([]entity.Webhook, error)
}

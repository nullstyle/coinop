package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

// DeliveryRepository tracks delivery state in persistent storage
type DeliveryRepository interface {
	StartDelivery(entity.Delivery) error
	MarkDeliveryFailed(entity.Delivery) error
	FailedDeliveries() ([]entity.Delivery, error)
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

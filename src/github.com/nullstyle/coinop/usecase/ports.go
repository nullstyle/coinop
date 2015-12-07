package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

type WebhookPresenter interface {
	List([]entity.Webhook) error
}

// WebhookRepository saves webhooks in persistent storage
type WebhookRepository interface {
	SaveWebhook(*entity.Webhook) error
	DestroyWebhook(ID RepoID) error
	ListWebhooks() ([]entity.Webhook, error)
}

// DeliveryRepository tracks delivery state in persistent storage
type DeliveryRepository interface {
	StartDelivery(entity.Delivery) error
	MarkDeliveryFailed(entity.Delivery) error
	FailedDeliveries() ([]entity.Delivery, error)
}

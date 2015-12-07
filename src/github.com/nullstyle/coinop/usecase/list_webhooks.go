package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

// ListWebhooks lists webhooks
type ListWebhooks struct {
	DB  WebhookRepository
	Out WebhookPresenter
}

// Exec runs the use case.
func (kase *ListWebhooks) Exec() (err error) {
	var hooks []entity.Webhook
	hooks, err = kase.DB.ListWebhooks()
	if err != nil {
		return
	}

	err = kase.Out.List(hooks)
	return
}

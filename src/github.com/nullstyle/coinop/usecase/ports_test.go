package usecase_test

import (
	"github.com/nullstyle/coinop/entity"
	. "github.com/nullstyle/coinop/usecase"
)

type mockWebhookRepository struct {
	Items  []entity.Webhook
	Err    error
	NoSave bool
}

func (repo *mockWebhookRepository) SaveWebhook(hook *entity.Webhook) error {
	if repo.Err != nil {
		return repo.Err
	}

	if repo.NoSave {
		return nil
	}

	if hook.IsNew() {
		hook.ID = &RepoID{
			T: "webhook",
			V: int64(len(repo.Items) + 1),
		}
		repo.Items = append(repo.Items, *hook)
	} else {
		id := hook.ID.(*RepoID).V
		repo.Items[id] = *hook
	}

	return nil
}

func (repo *mockWebhookRepository) DestroyWebhook(ID RepoID) error {
	return nil
}

func (repo *mockWebhookRepository) ListWebhooks() ([]entity.Webhook, error) {
	return append([]entity.Webhook{}, repo.Items...), nil
}

// ensure interface fulfillment
var _ WebhookRepository = &mockWebhookRepository{}

// var _ DeliveryRepository = &mockDeliveryRepository{}

package usecase_test

import (
	"github.com/nullstyle/coinop/entity"
	. "github.com/nullstyle/coinop/usecase"
)

type mockDeliverySender struct {
	Err error
}

func (send *mockDeliverySender) SendDelivery(entity.Delivery) error {
	return send.Err
}

type mockPaymentProvider struct {
	Queue []entity.Payment
	Err   error
}

func (p *mockPaymentProvider) StreamPayments(
	_ string,
	fn PaymentHandler,
) error {
	if p.Err != nil {
		return p.Err
	}

	for _, p := range p.Queue {
		err := fn(p)
		if err != nil {
			return err
		}
	}

	return nil
}

type mockDeliveryRepository struct {
	Cursor string
	Err    error
}

func (repo *mockPaymentRepository) SaveDeliveries(
	cursor string,
	ds []entity.Delivery,
) error {
	if repo.Err != nil {
		return repo.Err
	}

	repo.Cursor = cursor
	//TODO: dave deliveries
	return nil
}

func (repo *mockDeliveryRepository) LoadCursor() (string, error) {
	return repo.Cursor, repo.Err
}

type mockWebhookPresenter struct {
	Seen []entity.Webhook
	Err  error
}

func (p *mockWebhookPresenter) List(hooks []entity.Webhook) error {
	if p.Err != nil {
		return p.Err
	}

	p.Seen = append(p.Seen, hooks...)
	return nil
}

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
var _ DeliveryRepository = &mockDeliveryRepository{}
var _ DeliverySender = &mockDeliverySender{}

var _ PaymentProvider = &mockPaymentProvider{}

var _ WebhookPresenter = &mockWebhookPresenter{}
var _ WebhookRepository = &mockWebhookRepository{}

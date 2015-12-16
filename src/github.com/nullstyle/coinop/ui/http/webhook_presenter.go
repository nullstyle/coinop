package http

import (
	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
)

// List prints entities to stdout
func (ui *UI) List(hooks []entity.Webhook) error {
	var page WebhookPage
	page.Embedded.Records = make([]Webhook, len(hooks))

	for i, hook := range hooks {
		err := page.Embedded.Records[i].Scan(hook)
		if err != nil {
			return err
		}
	}

	return ui.JSON(page)
}

var _ usecase.WebhookPresenter = &UI{}

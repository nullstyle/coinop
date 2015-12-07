package console

import (
	"fmt"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
)

// List prints entities to stdout
func (d *Driver) List(hooks []entity.Webhook) error {
	for _, hook := range hooks {
		fmt.Printf("%s - %s\n", hook.ShortID(), hook.URL.String())
		fmt.Printf("  %s\n", hook.DestinationFilter)
		if hook.MemoFilter != nil {
			fmt.Printf("  memo(%s): %s\n",
				hook.MemoFilter.Type,
				hook.MemoFilter.Value,
			)
		}
	}

	return nil
}

var _ usecase.WebhookPresenter = &Driver{}

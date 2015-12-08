package editor

import (
	"net/url"

	"github.com/nullstyle/coinop/entity"
)

// Entity returns the core entity representation of the webhook
func (hook *Webhook) Entity() (result entity.Webhook, err error) {
	result.URL, err = url.Parse(hook.URL)
	if err != nil {
		return
	}

	result.DestinationFilter = entity.AccountID(hook.DestinationFilter)

	if hook.MemoTypeFilter != "" {
		result.MemoFilter = &entity.Memo{
			Type:  hook.MemoTypeFilter,
			Value: hook.MemoFilter,
		}
	}

	err = result.Valid()
	return
}

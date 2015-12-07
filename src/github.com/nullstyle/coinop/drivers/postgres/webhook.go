package postgres

import (
	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
	"net/url"
)

// Entity converts a postgres table row into a core entity
func (hook *Webhook) Entity() (result entity.Webhook, err error) {
	result.ID = &usecase.RepoID{T: "webhook", V: hook.ID}
	result.URL, err = url.Parse(hook.URL)
	if err != nil {
		return
	}
	result.DestinationFilter = hook.DestinationFilter
	if hook.MemoTypeFilter.Valid {
		result.MemoFilter = &entity.Memo{
			Type:  hook.MemoTypeFilter.String,
			Value: hook.MemoFilter.String,
		}
	}

	err = result.Valid()
	return
}

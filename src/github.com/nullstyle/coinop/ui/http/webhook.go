package http

import (
	"github.com/nullstyle/coinop/entity"
)

// Scan populates `json` from `ent`
func (json *Webhook) Scan(ent entity.Webhook) (err error) {
	if ent.ID != nil {
		json.ID = ent.ID.ShortID()
	}

	json.URL = ent.URL.String()
	json.DestinationFilter = string(ent.DestinationFilter)

	if ent.MemoFilter != nil {
		json.MemoTypeFilter = ent.MemoFilter.Type
		json.MemoFilter = ent.MemoFilter.Value
	}

	return nil
}

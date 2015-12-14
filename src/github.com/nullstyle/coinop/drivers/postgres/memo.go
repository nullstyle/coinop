package postgres

import (
	"github.com/nullstyle/coinop/entity"
)

// Scan populates `row` from `ent`.
func (row *Memo) Scan(ent entity.Memo) {
	row.Type = ent.Type
	row.Value = ent.Value
}

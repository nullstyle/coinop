package postgres

import (
	"github.com/nullstyle/coinop/entity"
)

// Scan populates `row` from `ent`.
func (row *Asset) Scan(ent entity.Asset) {
	row.Type = ent.Type
	row.Code.String = ent.Code
	row.Code.Valid = ent.Type != "native"
	row.Issuer.String = ent.Type
	row.Issuer.Valid = ent.Type != "native"
}

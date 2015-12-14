package postgres

import (
	"github.com/nullstyle/coinop/entity"
	"github.com/shopspring/decimal"
)

// Scan populates `row` from `ent`.
func (row *Payment) Scan(ent entity.Payment) (err error) {
	row.PagingToken = ent.PagingToken
	row.From = string(ent.From)
	row.To = string(ent.To)
	row.Memo.Scan(ent.Memo)
	row.Asset.Scan(ent.Asset)

	row.Amount = ent.Amount.Mul(decimal.New(1, 7)).IntPart()
	return
}

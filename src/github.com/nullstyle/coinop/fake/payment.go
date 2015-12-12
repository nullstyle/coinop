package fake

import (
	"github.com/nullstyle/coinop/entity"
)

// PaymentEntity returns a new random valid payment
func PaymentEntity() (result entity.Payment) {
	result.PagingToken = "1"
	result.To = AccountID()
	result.From = AccountID()
	result.Asset = entity.Asset{Type: "native"}
	result.Memo = entity.Memo{Type: "none"}
	result.Amount = entity.MustParseAmount("1.123")
	return
}

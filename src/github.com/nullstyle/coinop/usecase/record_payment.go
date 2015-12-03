package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

// RecordPayment Use Case
//
// A client should be able to record a new payment, so that the flow of
// commodities between users can be tracked.
//
// In:
// 	- Payee (user id)
// 	- Payer (user id)
// 	- Source (account id)
// 	- Commodity
// 	- Amount
// 	- TransactionID (optional)
// 	- Memo (optional)
//
// Output: none
//
//
// Steps:
// 1. A coinop client requests a new payment to be recorded by providing the
//  	sender, the receiver and the payment details.
// 2. System loads payee and payer User accounts and confirms they are in good
//  	standing.
// 3. TODO
//
type RecordPayment struct {
	UserRepo    UserRepository
	PaymentRepo interface{}
	AccountRepo interface{}
}

// Exec runs the use case.
func (kase *RecordPayment) Exec(from, to entity.AccountID) (err error) {
	return
}

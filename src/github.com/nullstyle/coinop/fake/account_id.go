package fake

import (
	"github.com/nullstyle/coinop/entity"
	"github.com/stellar/go-stellar-base/keypair"
)

// AccountID returns a new random account id
func AccountID() entity.AccountID {
	kp, err := keypair.Random()
	if err != nil {
		panic(err)
	}

	return entity.AccountID(kp.Address())
}

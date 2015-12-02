package entity

import (
	"github.com/stellar/go-stellar-base/xdr"
)

type Account struct {
	ID       uint64
	Names    []string
	Balances []Balance
}

type Balance struct {
	Asset  xdr.Asset
	Amount string
}

package entity

import (
	"github.com/shopspring/decimal"
)

// ParseAmount parses `in` into an Amount value
func ParseAmount(in string) (ret Amount, err error) {
	dec, err := decimal.NewFromString(in)
	ret = Amount{dec}
	return
}

// MustParseAmount parses `in` into an Amount value, panicing if it fails
func MustParseAmount(in string) Amount {
	ret, err := ParseAmount(in)
	if err != nil {
		panic(err)
	}
	return ret
}

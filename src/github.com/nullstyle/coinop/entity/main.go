package entity

import (
	"github.com/shopspring/decimal"
	"net/url"
	"time"
)

// AccountID is a stellar account id
type AccountID string

// Amount represents a fixed quantity
type Amount decimal.Decimal

// Asset represents an asset on the stellar network
type Asset struct {
	Type   string
	Code   string
	Issuer string
}

// Delivery represents a delivery of a payment to a specific url and state
// information.
type Delivery struct {
	ID
	URL          *url.URL
	Payment      Payment
	StartedAt    time.Time
	LastFailedAt time.Time
	SucceededAt  time.Time
}

// Memo represents a memo attached to a stellar transaction
type Memo struct {
	Type  string
	Value string
}

// OperationID represents an ID for an operation on the stellar network.
type OperationID string

// Payment represents a single payment that occurred on the
type Payment struct {
	ID     OperationID
	From   AccountID
	To     AccountID
	Memo   Memo
	Asset  Asset
	Amount Amount
}

// Webhook represents the desire to have an http request made when a payment
// occurs on the stellar network that matches a configured filter.
type Webhook struct {
	ID
	URL               *url.URL
	DestinationFilter string
	MemoFilter        *Memo
}

package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"time"
)

// Asset represents the asset portion of a row in the deliveries table of the
// postgres db
type Asset struct {
	Type   string         `db:"payment_asset_type"`
	Code   sql.NullString `db:"payment_asset_code"`
	Issuer sql.NullString `db:"payment_asset_issuer"`
}

// Delivery represents a database row from the `deliveries` table
type Delivery struct {
	ID      int64  `db:"id"`
	Version int64  `db:"version"`
	URL     string `db:"url"`
	Payment
	StartedAt    *time.Time `db:"started_at"`
	LastFailedAt *time.Time `db:"last_failed_at"`
	SucceededAt  *time.Time `db:"succeeded_at"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}

// Driver represents a connection to a postgres database
type Driver struct {
	DB *sqlx.DB `inject:""`
}

// KV represents a generic key-value pair.
type KV struct {
	Key   string         `db:"key"`
	Value types.JsonText `db:"value"`
}

// Memo represents the memo portion of a row in the deliveries table of the
// postgres db
type Memo struct {
	Type  string `db:"payment_memo_type"`
	Value string `db:"payment_memo"`
}

// Payment represents the payment portion of a row in the deliveries table of
// the postgres db
type Payment struct {
	PagingToken string `db:"payment_paging_token"`
	From        string `db:"payment_from"`
	To          string `db:"payment_to"`
	Memo
	Asset
	Amount int64 `db:"payment_amount"`
}

// Webhook represents a database row from the `webhooks` table
type Webhook struct {
	ID                int64          `db:"id"`
	URL               string         `db:"url"`
	DestinationFilter string         `db:"destination_filter"`
	MemoTypeFilter    sql.NullString `db:"memo_type_filter"`
	MemoFilter        sql.NullString `db:"memo_filter"`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         *time.Time     `db:"updated_at"`
}

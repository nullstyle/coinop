package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

//go:generate go-bindata -pkg postgres -o bindata.go migrations/

// // Delivery represents a database row from the `deliveries` table
// type Delivery struct {
// 	ID   int64  `db:"id"`
// 	URL  string `db:"url"`
// 	From string
// }

// Driver represents a connection to a postgres database
type Driver struct {
	DB *sqlx.DB `inject:""`
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

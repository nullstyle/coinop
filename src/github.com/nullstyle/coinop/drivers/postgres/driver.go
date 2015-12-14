package postgres

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"github.com/nullstyle/coinop/drivers/postgres/migrations"
	"github.com/rubenv/sql-migrate"
)

// MigrateSchema updates the schema in `db`.
func (db *Driver) MigrateSchema() error {
	migrate.SetSchema("coinop")
	_, err := migrate.Exec(db.DB.DB, "postgres", migrations.Source, migrate.Up)
	return err
}

// RebuildSchema rebuilds the schema in `db` by migrating down then up again.
func (db *Driver) RebuildSchema() error {
	migrate.SetSchema("coinop")
	_, err := migrate.Exec(db.DB.DB, "postgres", migrations.Source, migrate.Down)
	if err != nil {
		return err
	}

	_, err = migrate.Exec(db.DB.DB, "postgres", migrations.Source, migrate.Up)
	return err
}

// SaveKV saves `in` at `key` in the postgres db
func (db *Driver) SaveKV(key string, in interface{}) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = db.saveKV(tx, key, in)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetKV loads the value at `key` from the postgres db into `value`
func (db *Driver) GetKV(key string, value interface{}) error {
	var pair KV
	err := db.DB.Get(&pair, Queries.KV.Load, key)
	if err != nil {
		return err
	}

	return pair.Value.Unmarshal(value)
}

// saveKV saves a kv pair within the provided transaction
func (db *Driver) saveKV(tx *sqlx.Tx, key string, in interface{}) error {
	vj, err := json.Marshal(in)
	if err != nil {
		return err
	}

	_, err = tx.Exec(Queries.KV.Insert, key, types.JsonText(vj))
	if isUniqueErr(err) {
		_, err = tx.Exec(Queries.KV.Update, key, types.JsonText(vj))
	}

	return err
}

package postgres

import (
	"github.com/rubenv/sql-migrate"
)

var migrations migrate.MigrationSource = &migrate.AssetMigrationSource{
	Asset:    Asset,
	AssetDir: AssetDir,
	Dir:      "migrations",
}

// MigrateSchema updates the schema in `db`.
func (db *Driver) MigrateSchema() error {
	migrate.SetSchema("coinop")
	_, err := migrate.Exec(db.DB.DB, "postgres", migrations, migrate.Up)
	return err
}

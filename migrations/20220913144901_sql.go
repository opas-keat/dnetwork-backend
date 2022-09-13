package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upSql, downSql)
}

func upSql(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downSql(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

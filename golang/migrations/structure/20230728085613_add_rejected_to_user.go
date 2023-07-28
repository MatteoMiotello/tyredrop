package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddRejectedToUser, downAddRejectedToUser)
}

func upAddRejectedToUser(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users ADD COLUMN rejected bool NOT NULL DEFAULT false")

	if err != nil {
		return err
	}
	return nil
}

func downAddRejectedToUser(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddStatusToInvoice, downAddStatusToInvoice)
}

func upAddStatusToInvoice(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE invoices ADD COLUMN status varchar NULL")
	if err != nil {
		return err
	}
	return nil
}

func downAddStatusToInvoice(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE invoices DROP COLUMN status")
	if err != nil {
		return err
	}
	return nil
}

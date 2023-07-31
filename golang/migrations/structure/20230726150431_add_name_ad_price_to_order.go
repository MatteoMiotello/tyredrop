package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddNameAdPriceToOrder, downAddNameAdPriceToOrder)
}

func upAddNameAdPriceToOrder(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE orders ADD COLUMN address_name varchar NOT NULL DEFAULT ''")

	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE orders ADD COLUMN price_amount int4 NOT NULL DEFAULT 0")
	if err != nil {
		return err
	}

	return nil
}

func downAddNameAdPriceToOrder(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddOrderNumber, downAddOrderNumber)
}

func upAddOrderNumber(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE orders ADD COLUMN order_number varchar NULL")
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE orders SET order_number = ( SELECT concat( 'TITW', upper( CAST( to_hex( order1.id + 120000 ) AS varchar( 100 ) ) ) ) FROM orders as order1 WHERE order1.id = orders.id )")
	if err != nil {
		return err
	}
	return nil
}

func downAddOrderNumber(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE orders DROP COLUMN order_number")

	if err != nil {
		return err
	}

	return nil
}

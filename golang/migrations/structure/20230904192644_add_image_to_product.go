package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddImageToProduct, downAddImageToProduct)
}

func upAddImageToProduct(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE products ADD COLUMN image_url varchar NULL")
	if err != nil {
		return err
	}
	return nil
}

func downAddImageToProduct(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE products DROP COLUMN image_url")
	if err != nil {
		return err
	}
	return nil
}

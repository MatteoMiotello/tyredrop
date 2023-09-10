package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddSpecificationToPriceMarkup, downAddSpecificationToPriceMarkup)
}

func upAddSpecificationToPriceMarkup(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE product_price_markup ADD COLUMN product_specification_value_id bigint NULL")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE product_price_markup ADD CONSTRAINT fk_specification_value_id  FOREIGN KEY (product_specification_value_id) REFERENCES product_specification_values(id)")
	if err != nil {
		return err
	}
	return nil
}

func downAddSpecificationToPriceMarkup(tx *sql.Tx) error {
	return nil
}

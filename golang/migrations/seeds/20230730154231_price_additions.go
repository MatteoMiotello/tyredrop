package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upPriceAdditions, downPriceAdditions)
}

func upPriceAdditions(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.price_addition_types").
		Cols("currency_id", "addition_type", "addition_code", "addition_name", "addition_value").
		Values(1, "FIXED", "PFU_R_13_18", "PFU", 260).
		Values(1, "FIXED", "PFU_R_19_20", "PFU", 370).
		Values(1, "FIXED", "PFU_R_21_23", "PFU", 470).
		Values(1, "FIXED", "PFU_R_17.5", "PFU", 770).
		Values(1, "FIXED", "PFU_R_22.5_23.5", "PFU", 1870).
		Values(1, "FIXED", "PFU_MOTO", "PFU", 180).
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func downPriceAdditions(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM product_item_price_additions")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM price_addition_types")
	if err != nil {
		return err
	}
	return nil
}

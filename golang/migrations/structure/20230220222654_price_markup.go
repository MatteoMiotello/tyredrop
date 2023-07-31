package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upPriceMarkup, downPriceMarkup)
}

func upPriceMarkup(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.product_price_markup").
		PKColumn().
		FKColumn("public.product_categories", "product_category_id", true).
		FKColumn("public.brands", "brand_id", true).
		FKColumn("public.products", "product_id", true).
		Column("markup_percentage", types.Int, false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func downPriceMarkup(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.product_price_markup")
	if err != nil {
		return err
	}
	return nil
}

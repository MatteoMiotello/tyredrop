package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upPriceAdditions, downPriceAdditions)
}

func upPriceAdditions(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.price_addition_types").
		PKColumn().
		FKColumn("public.currencies", "currency_id", false).
		Column("addition_type", types.Varchar.Options("55"), false).
		Column("addition_name", types.Varchar.Options("55"), false).
		Column("addition_code", types.Varchar.Options("55"), false).
		Column("addition_value", types.Int, false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	query = sqlbuilder.CreateTable("public.product_item_price_additions").
		PKColumn().
		FKColumn("public.product_item_prices", "product_item_price_id", false).
		FKColumn("public.price_addition_types", "price_addition_type_id", false).
		Column("addition_value", types.Int, false).
		CreatedColumn().
		String()

	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.order_rows ADD COLUMN additions_amount int4 not null default 0")

	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.orders ADD COLUMN taxes_amount int4 not null default 0")

	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.orders ADD COLUMN price_amount_total int4 not null default 0")

	return nil
}

func downPriceAdditions(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS public.product_item_price_additions")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS public.price_addition_types")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.order_rows DROP COLUMN additions_amount")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.orders DROP COLUMN taxes_amount")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.orders DROP COLUMN price_amount_total")
	if err != nil {
		return err
	}
	return nil
}

package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upOrder, downOrder)
}

func upOrder(tx *sql.Tx) error {
	orderQuery := sqlbuilder.CreateTable("public.orders").
		PKColumn().
		FKColumn("public.currencies", "currency_id", false).
		FKColumn("public.tax_rates", "tax_rate_id", false).
		FKColumn("public.user_billings", "user_billing_id", false).
		Column("status", types.Varchar.Options("45"), false).
		Column("address_line_1", types.Varchar.Options("255"), false).
		Column("address_line_2", types.Varchar.Options("255"), false).
		Column("city", types.Varchar.Options("45"), false).
		Column("province", types.Varchar.Options("45"), false).
		Column("country", types.Varchar.Options("45"), false).
		Column("cap", types.Varchar.Options("5"), false).
		UpdatedColumn().
		CreatedColumn().
		String()

	orderRowQuery := sqlbuilder.CreateTable("public.order_rows").
		PKColumn().
		FKColumn("public.orders", "order_id", false).
		FKColumn("public.products", "product_id", false).
		FKColumn("public.suppliers", "supplier_id", false).
		Column("amount", types.Int, false).
		Column("tracking_number", types.Varchar.Options("255"), true).
		Column("carrier", types.Varchar.Options("45"), true).
		Column("sent_at", types.Timestamptz, true).
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(orderQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(orderRowQuery)
	if err != nil {
		return err
	}
	return nil
}

func downOrder(tx *sql.Tx) error {
	_, err := tx.Exec(" DROP TABLE IF EXISTS public.orders")
	if err != nil {
		return err
	}
	_, err = tx.Exec(" DROP TABLE IF EXISTS public.order_rows")
	if err != nil {
		return err
	}
	return nil
}

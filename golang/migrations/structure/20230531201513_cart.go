package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCart, downCart)
}

func upCart(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.carts").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		FKColumn("public.product_items", "product_item_id", false).
		Column("quantity", types.Int, false).
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

func downCart(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS public.carts")
	if err != nil {
		return err
	}
	return nil
}

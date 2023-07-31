package structure

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upRefactorCartOrderRow, downRefactorCartOrderRow)
}

func upRefactorCartOrderRow(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM public.carts")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.carts DROP CONSTRAINT fk_product_item")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.carts DROP COLUMN product_item_id")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.carts ADD COLUMN product_item_price_id int8 NOT NULL")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.carts ADD CONSTRAINT fk_product_price FOREIGN KEY (product_item_price_id) REFERENCES public.product_item_prices(id);")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM public.order_rows")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM public.orders")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.order_rows DROP CONSTRAINT fk_product_item")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.order_rows DROP COLUMN product_item_id")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.order_rows ADD COLUMN product_item_price_id int8 NOT NULL")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE public.order_rows ADD CONSTRAINT fk_product_price FOREIGN KEY (product_item_price_id) REFERENCES public.product_item_prices(id);")
	if err != nil {
		return err
	}

	return nil
}

func downRefactorCartOrderRow(tx *sql.Tx) error {

	return nil
}

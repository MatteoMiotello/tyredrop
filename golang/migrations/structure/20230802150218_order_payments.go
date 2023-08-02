package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upOrderPayments, downOrderPayments)
}

func upOrderPayments(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.payments").
		PKColumn().
		FKColumn("public.user_payment_methods", "user_payment_method_id", false).
		FKColumn("public.currencies", "currency_id", false).
		Column("amount", types.Int, false).
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE orders ADD COLUMN payment_id bigint NULL")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE orders ADD CONSTRAINT fk_payment_id  FOREIGN KEY (payment_id) REFERENCES payments(id)")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE user_payment_methods ALTER COLUMN value DROP NOT NULL")
	if err != nil {
		return err
	}

	return nil
}

func downOrderPayments(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

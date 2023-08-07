package seeds

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upPaymentMethods, downPaymentMethods)
}

func upPaymentMethods(tx *sql.Tx) error {
	query, values := sqlbuilder.InsertInto("public.payment_methods").
		Cols("code", "name").
		Values("SEPA", "SEPA").
		BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := tx.Exec(query, values...)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE user_payment_methods SET payment_method_id = 1 WHERE type = 'SEPA'")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE  user_payment_methods DROP COLUMN type")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE user_payment_methods ALTER COLUMN payment_method_id SET not null ")
	if err != nil {
		return err
	}

	return nil
}

func downPaymentMethods(tx *sql.Tx) error {
	return nil
}

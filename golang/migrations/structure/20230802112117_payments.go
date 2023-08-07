package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upPayments, downPayments)
}

func upPayments(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.payment_methods").
		PKColumn().
		Column("code", types.Varchar.Options("45"), false).
		Column("name", types.Varchar.Options("45"), false).
		Column("receiver", types.Varchar.Options("255"), true).
		Column("bank_name", types.Varchar.Options("255"), true).
		Column("iban", types.Varchar.Options("255"), true).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE user_payment_methods ADD COLUMN payment_method_id bigint NULL DEFAULT NULL")
	if err != nil {
		return err
	}

	_, err = tx.Exec("ALTER TABLE user_payment_methods ADD CONSTRAINT fk_payment_method_id  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)")
	if err != nil {
		return err
	}

	return nil
}

func downPayments(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

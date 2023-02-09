package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upFirstmigration, downFirstmigration)
}

func upFirstmigration(tx *sql.Tx) error {
	currencyQuery := sqlbuilder.CreateTable("public.currencies").
		PKColumn().
		Column("ISO_code", types.Varchar.Options("3"), false).
		Column("symbol", types.Varchar.Options("1"), false).
		Column("tag", types.Varchar.Options("45"), false).
		Column("precision", types.Int, false).
		CreatedColumn().
		String()

	languageQuery := sqlbuilder.CreateTable("public.languages").
		PKColumn().
		FKColumn("public.currencies", "currency_id", false).
		Column("name", types.Varchar.Options("45"), false).
		Column("ISO_code", types.Varchar.Options("3"), false).
		Column("tag", types.Varchar.Options("10"), false).
		Column("STD_timezone", types.Varchar.Options("3"), true).
		CreatedColumn().
		String()

	currencyLanguageQuery := sqlbuilder.CreateTable("public.currency_languages").
		PKColumn().
		FKColumn("public.currencies", "currency_id", false).
		FKColumn("public.languages", "language_id", false).
		Column("name", types.Varchar.Options("45"), false).
		CreatedColumn().
		String()

	_, err := tx.Exec(currencyQuery)
	if err != nil {
		return err
	}

	_, err = tx.Exec(languageQuery)
	if err != nil {
		return err
	}

	_, err = tx.Exec(currencyLanguageQuery)
	if err != nil {
		return err
	}

	return nil
}

func downFirstmigration(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.currency_languages")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.languages")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.currencies")
	if err != nil {
		return err
	}
	return nil
}

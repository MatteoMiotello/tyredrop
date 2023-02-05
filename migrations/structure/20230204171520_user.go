package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upUser, downUser)
}

func upUser(tx *sql.Tx) error {
	userRoleQuery := sqlbuilder.CreateTable("public.user_roles").
		PKColumn().
		Column("role_code", types.Varchar.Options("45"), false).
		Column("admin", types.Bool, false).
		CreatedColumn().
		String()

	userRoleLanguage := sqlbuilder.CreateTable("public.user_role_languages").
		PKColumn().
		FKColumn("public.languages", "language_id", false).
		Column("name", types.Varchar.Options("45"), false).
		CreatedColumn().
		String()

	userQuery := sqlbuilder.CreateTable("public.users").
		PKColumn().
		FKColumn("public.user_roles", "user_role_id", false).
		FKColumn("public.languages", "default_language_id", false).
		Column("email", types.Varchar.Options("100"), false).
		Column("username", types.Varchar.Options("100"), true).
		Column("password", types.Varchar.Options("255"), true).
		Column("name", types.Varchar.Options("100"), false).
		Column("surname", types.Varchar.Options("100"), false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	taxRateQuery := sqlbuilder.CreateTable("public.tax_rates").
		PKColumn().
		Column("markup_percentage", types.Int, false).
		Column("name", types.Varchar.Options("45"), false).
		CreatedColumn().
		String()

	legalEntitiesTypeQuery := sqlbuilder.CreateTable("public.legal_entity_types").
		PKColumn().
		Column("name", types.Varchar.Options("255"), false).
		CreatedColumn().
		String()

	userBilling := sqlbuilder.CreateTable("public.user_billings").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		FKColumn("public.tax_rates", "default_tax_rate_id", false).
		FKColumn("public.legal_entity_types", "legal_entity_type_id", false).
		Column("name", types.Varchar.Options("255"), false).
		Column("surname", types.Varchar.Options("255"), false).
		Column("fiscal_code", types.Varchar.Options("16"), false).
		Column("vat_number", types.Varchar.Options("11"), false).
		Column("address", types.Varchar.Options("255"), false).
		Column("city", types.Varchar.Options("45"), false).
		Column("province", types.Varchar.Options("45"), false).
		Column("cap", types.Varchar.Options("5"), false).
		Column("country", types.Varchar.Options("45"), false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(userRoleQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(userRoleLanguage)
	if err != nil {
		return err
	}
	_, err = tx.Exec(userQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(taxRateQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(legalEntitiesTypeQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(userBilling)
	if err != nil {
		return err
	}
	return nil
}

func downUser(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.user_billings IF EXISTS ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.legal_entity_types IF EXISTS ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.tax_rates IF EXISTS ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.users IF EXISTS ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.user_role_languages IF EXISTS ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.user_roles IF EXISTS ")
	if err != nil {
		return err
	}
	return nil
}

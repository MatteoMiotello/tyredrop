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
		FKColumn("public.user_roles", "user_role_id", false).
		Column("name", types.Varchar.Options("45"), false).
		CreatedColumn().
		String()

	userQuery := sqlbuilder.CreateTable("public.users").
		PKColumn().
		FKColumn("public.user_roles", "user_role_id", false).
		FKColumn("public.languages", "default_language_id", false).
		Column("email", types.Varchar.Options("100"), false).
		Column("username", types.Varchar.Options("100"), true).
		Column("password", types.Varchar.Options("255"), false).
		Column("name", types.Varchar.Options("100"), false).
		Column("surname", types.Varchar.Options("100"), true).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	taxRateQuery := sqlbuilder.CreateTable("public.taxes").
		PKColumn().
		Column("markup_percentage", types.Int, false).
		Column("name", types.Varchar.Options("45"), false).
		CreatedColumn().
		String()

	legalEntitiesTypeQuery := sqlbuilder.CreateTable("public.legal_entity_types").
		PKColumn().
		Column("name", types.Varchar.Options("255"), false).
		Column("is_person", types.Bool, false).
		CreatedColumn().
		String()

	userBilling := sqlbuilder.CreateTable("public.user_billings").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		FKColumn("public.legal_entity_types", "legal_entity_type_id", false).
		Column("name", types.Varchar.Options("255"), false).
		Column("surname", types.Varchar.Options("255"), true).
		Column("fiscal_code", types.Varchar.Options("16"), false).
		Column("vat_number", types.Varchar.Options("16"), false).
		Column("address_line_1", types.Varchar.Options("255"), false).
		Column("address_line_2", types.Varchar.Options("255"), true).
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
	_, err := tx.Exec("DROP TABLE public.user_billings  ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.legal_entity_types  ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.taxes  ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.users  ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.user_role_languages  ")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.user_roles  ")
	if err != nil {
		return err
	}
	return nil
}

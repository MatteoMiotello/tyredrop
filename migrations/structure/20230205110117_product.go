package structure

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upProduct, downProduct)
}

func upProduct(tx *sql.Tx) error {
	brandQuery := sqlbuilder.CreateTable("public.brands").
		PKColumn().
		Column("brand_code", types.Varchar.Options("45"), false).
		Column("name", types.Varchar.Options("255"), false).
		Column("image_logo", types.Varchar.Options("255"), true).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	productCategoryQuery := sqlbuilder.CreateTable("public.product_categories").
		PKColumn().
		Column("category_code", types.Varchar.Options("255"), false).
		Column("color", types.Varchar.Options("7"), true).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	productCategoryLanguage := sqlbuilder.CreateTable("public.product_category_languages").
		PKColumn().
		FKColumn("public.languages", "language_id", false).
		FKColumn("public.product_categories", "product_category_id", false).
		Column("name", types.Varchar.Options("255"), false).
		UpdatedColumn().
		CreatedColumn().
		String()

	productQuery := sqlbuilder.CreateTable("public.products").
		PKColumn().
		FKColumn("public.product_categories", "product_category_id", false).
		FKColumn("public.brands", "brand_id", false).
		Column("product_code", types.Varchar.Options("255"), true).
		Column("manufacturer_code", types.Varchar.Options("255"), true).
		Column("eprel_updated_at", types.Timestamptz, true).
		Column("completed", types.Bool, false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	productLanguageQuery := sqlbuilder.CreateTable("public.product_languages").
		PKColumn().
		FKColumn("public.products", "product_id", false).
		FKColumn("public.languages", "language_id", false).
		Column("name", types.Varchar.Options("45"), false).
		Column("description", types.Varchar.Options("255"), true).
		UpdatedColumn().
		CreatedColumn().
		String()

	productSpecificationQuery := sqlbuilder.CreateTable("public.product_specifications").
		PKColumn().
		FKColumn("public.product_categories", "product_category_id", false).
		Column("specification_code", types.Varchar.Options("45"), false).
		Column("type", types.Varchar.Options("45"), false).
		Column("mandatory", types.Bool, false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	productSpecificationLanguageQuery := sqlbuilder.CreateTable("public.product_specification_languages").
		PKColumn().
		FKColumn("public.product_specifications", "product_specification_id", false).
		FKColumn("public.languages", "language_id", false).
		Column("name", types.Varchar.Options("255"), false).
		Column("description", types.Varchar.Options("255"), true).
		UpdatedColumn().
		CreatedColumn().
		String()

	productSpecificationValueQuery := sqlbuilder.CreateTable("public.product_specification_values").
		PKColumn().
		FKColumn("public.products", "product_id", false).
		FKColumn("public.product_specifications", "product_specification_id", false).
		Column("specification_value", types.Varchar.Options("500"), false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	productItemQuery := sqlbuilder.CreateTable("public.product_items").
		PKColumn().
		FKColumn("public.products", "product_id", false).
		FKColumn("public.suppliers", "supplier_id", false).
		Column("supplier_price", types.Int, false).
		Column("supplier_quantity", types.Int, false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	productItemPriceQuery := sqlbuilder.CreateTable("public.product_item_prices").
		PKColumn().
		FKColumn("public.product_items", "product_item_id", false).
		FKColumn("public.currencies", "currency_id", false).
		Column("price", types.Int, false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(brandQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productCategoryQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productCategoryLanguage)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productLanguageQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productSpecificationQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productSpecificationLanguageQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productSpecificationValueQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productItemQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec(productItemPriceQuery)
	if err != nil {
		return err
	}
	_, err = tx.Exec("CREATE INDEX spec_code_idx ON public.product_specifications (specification_code);")
	if err != nil {
		return err
	}

	return nil
}

func downProduct(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.product_item_prices")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.product_items")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.product_specification_values")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.product_specification_languages")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.product_specifications")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.product_languages")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.products")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.product_category_languages")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.product_categories")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.brands")
	if err != nil {
		return err
	}
	return nil
}

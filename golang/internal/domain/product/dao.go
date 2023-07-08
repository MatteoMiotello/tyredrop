package product

import (
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type Dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		db.DaoFromExecutor(executor),
	}
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(limit int, offset int) *Dao {
	return db.Paginate(d, limit, offset)
}

func (d Dao) ForUpdate() *Dao {
	return db.ForUpdate(d)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d *Dao) ProductSpecificationValues(ctx context.Context, product *models.Product) (models.ProductSpecificationValueSlice, error) {
	return product.ProductSpecificationValues(
		d.GetMods()...,
	).All(ctx, d.Db)
}

func (d *Dao) Brand(ctx context.Context, product *models.Product) (*models.Brand, error) {
	return product.Brand(
		d.GetMods()...,
	).One(ctx, d.Db)
}

func (d *Dao) FindAll(ctx context.Context) (models.ProductSlice, error) {
	return models.Products(
		d.GetMods()...,
	).All(ctx, db.DB)
}

func (d *Dao) CountAll(ctx context.Context) (int64, error) {
	return models.Products(
		d.GetMods()...,
	).Count(ctx, d.Db)
}

func (d *Dao) FindOneById(ctx context.Context, id int64) (*models.Product, error) {
	return models.Products(
		d.GetMods(
			models.ProductWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d *Dao) FindOneByCode(ctx context.Context, productCode string) (*models.Product, error) {
	return models.Products(
		d.GetMods(
			models.ProductWhere.ProductCode.EQ(null.StringFrom(productCode)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindAllPriceForProductItem(ctx context.Context, pi *models.ProductItem) (models.ProductItemPriceSlice, error) {
	return models.ProductItemPrices(
		d.GetMods(
			models.ProductItemPriceWhere.ProductItemID.EQ(pi.ID),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) DeleteProductItemPrice(ctx context.Context, price *models.ProductItemPrice) error {
	_, err := price.Delete(ctx, d.Db, false)

	if err != nil {
		return err
	}

	return nil
}

func (d Dao) FindNextRemainingEprelProduct(ctx context.Context, categoryCodes ...string) (*models.Product, error) {
	return models.Products(
		d.GetMods(
			qm.LeftOuterJoin("product_categories on products.product_category_id = product_categories.id"),
			models.ProductWhere.EprelProductCode.IsNotNull(),
			models.ProductCategoryWhere.CategoryCode.IN(categoryCodes),
			models.ProductWhere.EprelUpdatedAt.IsNull(),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) AddProductSpecificationValues(ctx context.Context, product *models.Product, values ...*models.ProductSpecificationValue) error {
	return product.AddProductSpecificationValues(ctx, d.Db, true, values...)
}

func (d Dao) Search(ctx context.Context, input *model.ProductSearchInput, currency *models.Currency) (models.ProductSlice, error) {
	var mods []qm.QueryMod

	if input != nil {
		if input.Code != nil {
			mods = append(mods, models.ProductWhere.ProductCode.EQ(null.StringFrom(*input.Name)))
		}

		if input.Brand != nil {
			mods = append(mods, models.BrandWhere.BrandCode.EQ(*input.Brand))
		}

		if input.Name != nil {
			mods = append(mods, qm.Where("products.name ILIKE ?", "%"+*input.Name+"%"))
		}

		for _, spec := range input.Specifications {
			mods = append(mods, qm.Where("product_specifications.specification_code = ? AND product_specification_values.specification_value = ?", spec.Code, spec.Value))
		}
	}

	return models.Products(
		d.GetMods(
			qm.LeftOuterJoin("product_items on product_items.product_id = products.id"),
			qm.LeftOuterJoin("product_item_prices on product_item_prices.product_item_id = product_items.id"),
			qm.LeftOuterJoin("product_specification_values on product_specification_values.product_id = products.id"),
			qm.LeftOuterJoin("product_specifications on product_specification_values.product_specification_id = product_specifications.id"),
			qm.LeftOuterJoin("brands on brands.id = products.brand_id"),
			qm.Expr(mods...),
			models.ProductItemPriceWhere.CurrencyID.EQ(currency.ID),
			qm.GroupBy("products.id, product_item_prices.price"),
			qm.OrderBy("product_item_prices.price ASC"),
		)...,
	).All(ctx, d.Db)
}

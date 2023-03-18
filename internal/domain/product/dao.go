package product

import (
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (d *Dao) SetDao(dao *db.Dao) {
	d.Dao = dao
}

func (d *Dao) GetDao() *db.Dao {
	return d.Dao
}

func (d *Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *Dao) Paginate(first int, offset int) *Dao {
	db.Paginate(d, first, offset)

	return d
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

func (d *Dao) FindOneById(ctx context.Context, id int64) (*models.Product, error) {
	return models.Products(
		d.GetMods()...,
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
			models.ProductCategoryWhere.CategoryCode.IN(categoryCodes),
			models.ProductWhere.EprelUpdatedAt.IsNull(),
		)...,
	).One(ctx, d.Db)
}

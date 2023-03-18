package product

import (
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type ItemDao struct {
	db.Dao
}

func NewItemDao(exec boil.ContextExecutor) *ItemDao {
	return &ItemDao{
		db.DaoFromExecutor(exec),
	}
}

func (d *ItemDao) Load(relationship string, mods ...qm.QueryMod) *ItemDao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *ItemDao) Paginate(first int, offset int) *ItemDao {
	db.Paginate(d, first, offset)

	return d
}

func (d *ItemDao) Product(ctx context.Context, productItem *models.ProductItem) (*models.Product, error) {
	return productItem.Product(
		d.GetMods()...,
	).One(ctx, d.Db)
}

func (d *ItemDao) FindProductItemById(ctx context.Context, id int64) (*models.ProductItem, error) {
	return models.ProductItems(
		d.GetMods(
			models.ProductItemWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d *ItemDao) ProductItemPrices(ctx context.Context, productItem *models.ProductItem) (models.ProductItemPriceSlice, error) {
	return productItem.ProductItemPrices(
		d.GetMods()...,
	).All(ctx, d.Db)
}

func (d *ItemDao) FindByProductAndSupplier(ctx context.Context, product *models.Product, supplier *models.Supplier) (models.ProductItemSlice, error) {
	return models.ProductItems(
		d.GetMods(
			models.ProductItemWhere.ProductID.EQ(product.ID),
			models.ProductItemWhere.SupplierID.EQ(supplier.ID),
		)...,
	).All(ctx, d.Db)
}

func (d *ItemDao) FindByCategory(ctx context.Context, category *models.ProductCategory) (models.ProductItemSlice, error) {
	return models.ProductItems(
		d.GetMods(
			qm.LeftOuterJoin("products on products.id = product_items.product_id"),
			models.ProductWhere.ProductCategoryID.EQ(category.ID),
		)...,
	).All(ctx, d.Db)
}

func (d *ItemDao) FindLessExpensiveByProductCode(ctx context.Context, code string) (*models.ProductItem, error) {
	return models.ProductItems(
		d.GetMods(
			qm.LeftOuterJoin("products on products.id = product_items.product_id"),
			qm.LeftOuterJoin("product_item_prices on product_items.id = product_item_prices.product_item_id"),
			qm.OrderBy("product_item_prices.price ASC"),
			models.ProductWhere.ProductCode.EQ(null.StringFrom(code)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

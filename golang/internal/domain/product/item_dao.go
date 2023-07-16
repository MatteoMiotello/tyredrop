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

type ItemDao struct {
	*db.Dao
}

func NewItemDao(exec boil.ContextExecutor) *ItemDao {
	return &ItemDao{
		db.DaoFromExecutor(exec),
	}
}

func (d ItemDao) Clone() db.DaoMod {
	return ItemDao{
		d.Dao.Clone(),
	}
}

func (d ItemDao) Load(relationship string, mods ...qm.QueryMod) *ItemDao {
	return db.Load(d, relationship, mods...)
}

func (d ItemDao) Paginate(first int, offset int) *ItemDao {
	return db.Paginate(d, first, offset)
}

func (d ItemDao) WithDeletes() *ItemDao {
	return db.WithDeletes(d)
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

func (d *ItemDao) ProductItemPrice(ctx context.Context, item *models.ProductItem, currency *models.Currency) (*models.ProductItemPrice, error) {
	return item.ProductItemPrices(
		d.GetMods(
			models.ProductItemPriceWhere.CurrencyID.EQ(currency.ID),
		)...,
	).One(ctx, d.Db)
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

type SpecificationInput struct {
	name  string
	value string
}

func (d *ItemDao) FindLessExpensiveBySpecifications(ctx context.Context, inputs ...*SpecificationInput) (models.ProductItemSlice, error) {
	return models.ProductItems(
		d.GetMods()...,
	).All(ctx, d.Db)
}

func (d *ItemDao) FindProductItems(ctx context.Context, input *model.ProductSearchInput, currency *models.Currency) (models.ProductItemSlice, error) {
	var mods []qm.QueryMod

	mods = append(mods, qm.LeftOuterJoin("products on products.id = product_items.product_id"))
	mods = append(mods, qm.LeftOuterJoin(
		"( "+
			"SELECT product_items.product_id, MIN( product_item_prices.price ) AS min_price "+
			"FROM product_items "+
			"LEFT JOIN product_item_prices on product_items.id = product_item_prices.product_item_id "+
			"AND product_item_prices.currency_id = ? "+
			"AND product_item_prices.deleted_at IS NULL "+
			"GROUP BY product_items.product_id "+
			" ) pi ON pi.product_id = products.id ", currency.ID,
	))
	mods = append(mods, qm.LeftOuterJoin("product_item_prices ON product_item_prices.product_item_id = product_items.id AND product_item_prices.price = pi.min_price"))
	mods = append(mods, qm.LeftOuterJoin("brands on brands.id = products.brand_id"))
	mods = append(mods, models.ProductItemPriceWhere.CurrencyID.EQ(currency.ID))
	mods = append(mods, qm.GroupBy("product_items.id, product_item_prices.id"))
	mods = append(mods, models.ProductItemWhere.SupplierQuantity.GTE(4))

	if input != nil {
		if input.Code != nil {
			mods = append(mods, models.ProductWhere.ProductCode.EQ(null.StringFrom(*input.Code)))
		}

		if input.Brand != nil {
			mods = append(mods, models.BrandWhere.BrandCode.EQ(*input.Brand))
		}

		if input.Name != nil {
			mods = append(mods, qm.Where("products.name ILIKE ?", "%"+*input.Name+"%"))
		}

		if input.VehicleCode != nil {
			mods = append(mods, qm.LeftOuterJoin("vehicle_types on vehicle_types.id = products.vehicle_type_id"))
			mods = append(mods, models.VehicleTypeWhere.Code.EQ(*input.VehicleCode))
		}

		if input.Specifications != nil {
			for _, spec := range input.Specifications {
				mods = append(mods, qm.And(
					"products.id IN ( SELECT product_product_specification_values.product_id FROM product_product_specification_values "+
						"LEFT JOIN product_specification_values ON product_product_specification_values.product_specification_value_id = product_specification_values.id "+
						"LEFT JOIN product_specifications ON product_specifications.id = product_specification_values.product_specification_id AND product_specifications.deleted_at IS NULL "+
						"WHERE product_specifications.specification_code = ? "+
						"AND product_specification_values.specification_value = ? ) ",
					spec.Code, spec.Value,
				))
			}
		}
	}

	mods = append(mods, qm.OrderBy(models.ProductItemPriceColumns.Price))

	return models.ProductItems(
		d.GetMods(
			mods...,
		)...,
	).All(ctx, d.Db)
}

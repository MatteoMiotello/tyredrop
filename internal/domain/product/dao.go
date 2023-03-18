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
	db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		db.DaoFromExecutor(executor),
	}
}

func (d *Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *Dao) Paginate(first int, offset int) *Dao {
	db.Paginate(d, first, offset)

	return d
}

func (d Dao) ProductSpecificationValues(ctx context.Context, product *models.Product) (models.ProductSpecificationValueSlice, error) {
	return product.ProductSpecificationValues().All(ctx, d.Db)
}

func (d Dao) Brand(ctx context.Context, product *models.Product) (*models.Brand, error) {
	return product.Brand().One(ctx, d.Db)
}

func (d Dao) Currency(ctx context.Context, price *models.ProductItemPrice) (*models.Currency, error) {
	return price.Currency().One(ctx, d.Db)
}

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Product, error) {
	return models.FindProduct(ctx, d.Db, id)
}

func (d Dao) FindOneByProductCode(ctx context.Context, productCode string) (*models.Product, error) {
	return models.Products(
		d.GetMods(
			models.ProductWhere.ProductCode.EQ(null.StringFrom(productCode)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindCategoryByCode(ctx context.Context, code string) (*models.ProductCategory, error) {
	return models.ProductCategories(
		d.GetMods(
			models.ProductCategoryWhere.CategoryCode.EQ(code),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindCategoryById(ctx context.Context, id int64) (*models.ProductCategory, error) {
	return models.ProductCategories(
		d.GetMods(
			models.ProductCategoryWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) DeleteProductItem(ctx context.Context, item *models.ProductItem) error {
	_, err := item.Delete(ctx, d.Db, false)
	if err != nil {
		return err
	}

	return nil
}

func (d Dao) FindProductItemsByProductAndSupplier(ctx context.Context, product *models.Product, supplier *models.Supplier) (models.ProductItemSlice, error) {
	return models.ProductItems(
		d.GetMods(
			models.ProductItemWhere.ProductID.EQ(product.ID),
			models.ProductItemWhere.SupplierID.EQ(supplier.ID),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindProductSpecificationValueByProductAndCode(ctx context.Context, product *models.Product, specificationCode string) (*models.ProductSpecificationValue, error) {
	return models.ProductSpecificationValues(
		d.GetMods(
			qm.LeftOuterJoin("product_specifications on product_specifications.id = product_specification_values.product_specification_id"),
			models.ProductSpecificationValueWhere.ProductID.EQ(product.ID),
			models.ProductSpecificationWhere.SpecificationCode.EQ(specificationCode),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindProductSpecificationValuesByProductId(ctx context.Context, id int64) (models.ProductSpecificationValueSlice, error) {
	return models.ProductSpecificationValues(
		models.ProductSpecificationValueWhere.ProductID.EQ(id),
	).All(ctx, d.Db)
}

func (d Dao) FindProductSpecificationValuesByCodes(ctx context.Context, product *models.Product, specificationCodes []string) (models.ProductSpecificationValueSlice, error) {
	return models.ProductSpecificationValues(
		d.GetMods(
			qm.LeftOuterJoin("product_specifications on product_specifications.id = product_specification_values.product_specification_id"),
			models.ProductSpecificationValueWhere.ProductID.EQ(product.ID),
			models.ProductSpecificationWhere.SpecificationCode.IN(specificationCodes),
			qm.Load(models.ProductSpecificationValueRels.ProductSpecification),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindOneProductSpecificationByCode(ctx context.Context, specCode string) (*models.ProductSpecification, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.SpecificationCode.EQ(specCode),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupByProductId(ctx context.Context, product *models.Product) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.ProductID.EQ(null.Int64From(product.ID)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupByBrandId(ctx context.Context, id int64) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.BrandID.EQ(null.Int64From(id)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupByProductCategoryId(ctx context.Context, categoryId int64) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.ProductCategoryID.EQ(null.Int64From(categoryId)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupDefault(ctx context.Context) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.ProductCategoryID.IsNull(),
			models.ProductPriceMarkupWhere.BrandID.IsNull(),
			models.ProductPriceMarkupWhere.ProductID.IsNull(),
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

func (d Dao) FindProductsByCategoryCodes(ctx context.Context, categoryCodes ...string) (models.ProductSlice, error) {
	return models.Products(
		d.GetMods(
			qm.LeftOuterJoin("product_categories on products.product_category_id = product_categories.id"),
			models.ProductCategoryWhere.CategoryCode.IN(categoryCodes),
		)...,
	).All(ctx, d.Db)
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

func (d Dao) FindMandatoryProductSpecificationsByProduct(ctx context.Context, product *models.Product) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ProductCategoryID.EQ(product.ProductCategoryID),
			models.ProductSpecificationWhere.Mandatory.EQ(true),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindProductSpecificationsByProduct(ctx context.Context, product *models.Product) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ProductCategoryID.EQ(product.ProductCategoryID),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindProductSpecificationsByProductCategoryId(ctx context.Context, id int64) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ProductCategoryID.EQ(id),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindProductItemsByCategory(ctx context.Context, category *models.ProductCategory) (models.ProductItemSlice, error) {
	return models.ProductItems(
		d.GetMods(
			qm.LeftOuterJoin("products on products.id = product_items.product_id"),
			models.ProductWhere.ProductCategoryID.EQ(category.ID),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindProductItemByBrand(ctx context.Context, brand *models.Brand) (models.ProductItemSlice, error) {
	return models.ProductItems(
		d.GetMods(
			qm.LeftOuterJoin("products on products.id = product_items.product_id"),
			models.ProductWhere.BrandID.EQ(brand.ID),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindProductItemByProduct(ctx context.Context, product *models.Product) (models.ProductItemSlice, error) {
	return product.ProductItems(
		d.GetMods()...,
	).All(ctx, d.Db)
}

func (d Dao) FindProductItemById(ctx context.Context, id int64) (*models.ProductItem, error) {
	return models.ProductItems(
		d.GetMods(
			models.ProductItemWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d *Dao) FindLessExpensiveProductItemByProductCode(ctx context.Context, code string) (*models.ProductItem, error) {
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

func (d *Dao) FindProductSpecificationsByCategoryId(ctx context.Context, id int64) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ProductCategoryID.EQ(id),
		)...,
	).All(ctx, d.Db)
}

func (d *Dao) FindProductSpecificationById(ctx context.Context, id int64) (*models.ProductSpecification, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ID.EQ(id),
		)...,
	).One(ctx, db.DB)
}

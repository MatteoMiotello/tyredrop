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

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Product, error) {
	return models.FindProduct(ctx, d.Db, id)
}

func (d Dao) FindOneByProductCode(ctx context.Context, productCode string) (*models.Product, error) {
	return models.Products(models.ProductWhere.ProductCode.EQ(null.StringFrom(productCode)), qm.Limit(1)).One(ctx, d.Db)
}

func (d Dao) FindCategoryByCode(ctx context.Context, code string) (*models.ProductCategory, error) {
	return models.ProductCategories(models.ProductCategoryWhere.CategoryCode.EQ(code), qm.Limit(1)).One(ctx, d.Db)
}

func (d Dao) DeleteProductItem(ctx context.Context, item *models.ProductItem) error {
	_, err := item.Delete(ctx, d.Db, false)
	if err != nil {
		return err
	}

	return nil
}

func (d Dao) FindProductItems(ctx context.Context, product *models.Product, supplier *models.Supplier) (models.ProductItemSlice, error) {
	return models.ProductItems(
		models.ProductItemWhere.ProductID.EQ(product.ID),
		models.ProductItemWhere.SupplierID.EQ(supplier.ID),
	).All(ctx, d.Db)
}

func (d Dao) FindProductSpecificationValue(ctx context.Context, product *models.Product, specificationCode string) (*models.ProductSpecificationValue, error) {
	return models.ProductSpecificationValues(
		qm.LeftOuterJoin("product_specifications on product_specifications.id = product_specification_values.product_specification_id"),
		models.ProductSpecificationValueWhere.ProductID.EQ(product.ID),
		models.ProductSpecificationWhere.SpecificationCode.EQ(specificationCode),
	).One(ctx, d.Db)
}

func (d Dao) FindProductSpecificationValuesByCodes(ctx context.Context, product *models.Product, specificationCodes []string) (models.ProductSpecificationValueSlice, error) {
	return models.ProductSpecificationValues(
		qm.LeftOuterJoin("product_specifications on product_specifications.id = product_specification_values.product_specification_id"),
		models.ProductSpecificationValueWhere.ProductID.EQ(product.ID),
		models.ProductSpecificationWhere.SpecificationCode.IN(specificationCodes),
		qm.Load(models.ProductSpecificationValueRels.ProductSpecification),
	).All(ctx, d.Db)
}

func (d Dao) FindOneProductSpecificationByCode(ctx context.Context, specCode string) (*models.ProductSpecification, error) {
	return models.ProductSpecifications(models.ProductSpecificationWhere.SpecificationCode.EQ(specCode), qm.Limit(1)).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupByProductId(ctx context.Context, product *models.Product) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(models.ProductPriceMarkupWhere.ProductID.EQ(null.Int64From(product.ID)), qm.Limit(1)).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupByBrandId(ctx context.Context, id int64) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(models.ProductPriceMarkupWhere.BrandID.EQ(null.Int64From(id)), qm.Limit(1)).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupByProductCategoryId(ctx context.Context, categoryId int64) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(models.ProductPriceMarkupWhere.ProductCategoryID.EQ(null.Int64From(categoryId)), qm.Limit(1)).One(ctx, d.Db)
}

func (d Dao) FindPriceMarkupDefault(ctx context.Context) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		qm.Where(models.ProductPriceMarkupColumns.ProductCategoryID+" is NULL"),
		qm.And(models.ProductPriceMarkupColumns.BrandID+" is NULL"),
		qm.And(models.ProductPriceMarkupColumns.ProductID+" is NULL"),
	).One(ctx, d.Db)
}

func (d Dao) FindAllPriceForProductItem(ctx context.Context, pi *models.ProductItem) (models.ProductItemPriceSlice, error) {
	return models.ProductItemPrices(models.ProductItemPriceWhere.ProductItemID.EQ(pi.ID)).All(ctx, d.Db)
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
		qm.LeftOuterJoin("product_categories on products.product_category_id = product_categories.id"),
		models.ProductCategoryWhere.CategoryCode.IN(categoryCodes),
	).All(ctx, d.Db)
}

func (d Dao) FindNextRemainingEprelProduct(ctx context.Context, categoryCodes ...string) (*models.Product, error) {
	return models.Products(
		qm.LeftOuterJoin("product_categories on products.product_category_id = product_categories.id"),
		models.ProductCategoryWhere.CategoryCode.IN(categoryCodes),
		models.ProductWhere.EprelUpdatedAt.IsNull(),
	).One(ctx, d.Db)
}

func (d Dao) FindMandatoryProductSpecificationsByProduct(ctx context.Context, product *models.Product) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		models.ProductSpecificationWhere.ProductCategoryID.EQ(product.ProductCategoryID),
		models.ProductSpecificationWhere.Mandatory.EQ(true),
	).All(ctx, d.Db)
}

func (d Dao) FindProductSpecificationsByProduct(ctx context.Context, product *models.Product) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		models.ProductSpecificationWhere.ProductCategoryID.EQ(product.ProductCategoryID),
	).All(ctx, d.Db)
}

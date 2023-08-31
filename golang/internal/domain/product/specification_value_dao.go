package product

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type SpecificationValueDao struct {
	*db.Dao
}

func NewSpecificationValueDao(exec boil.ContextExecutor) *SpecificationValueDao {
	return &SpecificationValueDao{
		db.DaoFromExecutor(exec),
	}
}

func (d SpecificationValueDao) Clone() db.DaoMod {
	return SpecificationValueDao{
		d.Dao.Clone(),
	}
}

func (d SpecificationValueDao) Load(relationship string, mods ...qm.QueryMod) *SpecificationValueDao {
	return db.Load(d, relationship, mods...)
}

func (d SpecificationValueDao) Paginate(first int, offset int) *SpecificationValueDao {
	return db.Paginate(d, first, offset)
}

func (d *SpecificationValueDao) ProductSpecification(ctx context.Context, specificationValue *models.ProductSpecificationValue) (*models.ProductSpecification, error) {
	return specificationValue.ProductSpecification(
		d.GetMods()...,
	).One(ctx, d.Db)
}

func (d *SpecificationValueDao) FindByProductAndCode(ctx context.Context, product *models.Product, specificationCode string) (*models.ProductSpecificationValue, error) {
	return models.ProductSpecificationValues(
		d.GetMods(
			qm.LeftOuterJoin("product_product_specification_values on product_product_specification_values.product_specification_value_id = product_specification_values.id"),
			qm.LeftOuterJoin("product_specifications on product_specifications.id = product_specification_values.product_specification_id"),
			qm.Where("product_specifications.specification_code = ?", specificationCode),
			models.ProductProductSpecificationValueWhere.ProductID.EQ(product.ID),
		)...,
	).One(ctx, d.Db)
}

func (d *SpecificationValueDao) FindByProductId(ctx context.Context, id int64) (models.ProductSpecificationValueSlice, error) {
	return models.ProductSpecificationValues(
		qm.LeftOuterJoin("product_product_specification_values on product_product_specification_values.product_specification_value_id = product_specification_values.id"),
		models.ProductProductSpecificationValueWhere.ProductID.EQ(id),
	).All(ctx, d.Db)
}

func (d *SpecificationValueDao) FindBySpecificationId(ctx context.Context, id int64) (models.ProductSpecificationValueSlice, error) {
	return models.ProductSpecificationValues(
		d.GetMods(
			qm.Distinct(models.ProductSpecificationValueColumns.SpecificationValue),
			models.ProductSpecificationValueWhere.ProductSpecificationID.EQ(id),
		)...,
	).All(ctx, d.Db)
}

func (d *SpecificationValueDao) FindBySpecificationAndValue(ctx context.Context, specification *models.ProductSpecification, value string) (*models.ProductSpecificationValue, error) {
	return models.ProductSpecificationValues(
		d.GetMods(
			models.ProductSpecificationValueWhere.ProductSpecificationID.EQ(specification.ID),
			models.ProductSpecificationValueWhere.SpecificationValue.EQ(value),
		)...,
	).One(ctx, d.Db)
}

func (d *SpecificationValueDao) SearchBySpecificationAndValue(ctx context.Context, code string, value *string, vehicleCode *string) (models.ProductSpecificationValueSlice, error) {
	var mods []qm.QueryMod

	mods = append(mods, qm.LeftOuterJoin("product_specifications on product_specifications.id = product_specification_values.product_specification_id"))
	mods = append(mods, qm.LeftOuterJoin("product_product_specification_values on product_specification_values.id = product_product_specification_values.product_specification_value_id"))
	mods = append(mods, models.ProductSpecificationWhere.SpecificationCode.EQ(code))
	mods = append(mods, qm.GroupBy("product_specification_values.specification_value, product_specification_values.id"))

	if value != nil {
		mods = append(mods, qm.Where(models.ProductSpecificationValueColumns.SpecificationValue+" ILIKE ?", `%`+*value+`%`))
	}

	if vehicleCode != nil && len(*vehicleCode) > 0 {
		mods = append(mods, qm.LeftOuterJoin("products on product_product_specification_values.product_id = products.id"))
		mods = append(mods, qm.LeftOuterJoin("vehicle_types on products.vehicle_type_id = vehicle_types.id"))
		mods = append(mods, models.VehicleTypeWhere.Code.EQ(*vehicleCode))
	}

	return models.ProductSpecificationValues(
		d.GetMods(
			mods...,
		)...,
	).All(ctx, d.Db)
}

func (d SpecificationValueDao) FindOneById(ctx context.Context, id int64) (*models.ProductSpecificationValue, error) {
	return models.ProductSpecificationValues(
		d.GetMods(
			models.ProductSpecificationValueWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

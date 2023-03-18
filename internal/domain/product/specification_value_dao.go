package product

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type SpecificationValueDao struct {
	db.Dao
}

func NewSpecificationValueDao(exec boil.ContextExecutor) *SpecificationValueDao {
	return &SpecificationValueDao{
		db.DaoFromExecutor(exec),
	}
}

func (d *SpecificationValueDao) Load(relationship string, mods ...qm.QueryMod) *SpecificationValueDao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *SpecificationValueDao) Paginate(first int, offset int) *SpecificationValueDao {
	db.Paginate(d, first, offset)

	return d
}

func (d *SpecificationValueDao) ProductSpecification(ctx context.Context, specificationValue *models.ProductSpecificationValue) (*models.ProductSpecification, error) {
	return specificationValue.ProductSpecification(
		d.GetMods()...,
	).One(ctx, d.Db)
}

func (d *SpecificationValueDao) FindByProductAndCode(ctx context.Context, product *models.Product, specificationCode string) (*models.ProductSpecificationValue, error) {
	return models.ProductSpecificationValues(
		d.GetMods(
			qm.LeftOuterJoin("product_specifications on product_specifications.id = product_specification_values.product_specification_id"),
			models.ProductSpecificationValueWhere.ProductID.EQ(product.ID),
			models.ProductSpecificationWhere.SpecificationCode.EQ(specificationCode),
		)...,
	).One(ctx, d.Db)
}

func (d *SpecificationValueDao) FindByProductId(ctx context.Context, id int64) (models.ProductSpecificationValueSlice, error) {
	return models.ProductSpecificationValues(
		models.ProductSpecificationValueWhere.ProductID.EQ(id),
	).All(ctx, d.Db)
}

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

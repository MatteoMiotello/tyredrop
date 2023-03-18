package product

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/language"
	"pillowww/titw/models"
)

type SpecificationDao struct {
	*db.Dao
}

type DaoModded struct {
	*SpecificationDao
	mods []qm.QueryMod
}

func NewSpecificationDao(exec boil.ContextExecutor) *SpecificationDao {
	return &SpecificationDao{
		db.DaoFromExecutor(exec),
	}
}

func (d *SpecificationDao) SetDao(dao *db.Dao) {
	d.Dao = dao
}

func (d *SpecificationDao) GetDao() *db.Dao {
	return d.Dao
}

func (d *SpecificationDao) Load(relationship string, mods ...qm.QueryMod) *SpecificationDao {
	new := &SpecificationDao{
		d.Dao,
	}

	db.Load(new, relationship, mods...)

	return new
}

func (d *SpecificationDao) Paginate(first int, offset int) *SpecificationDao {
	new := &SpecificationDao{
		d.Dao,
	}

	db.Paginate(new, first, offset)

	return new
}

func (d *SpecificationDao) ProductSpecificationLanguage(ctx context.Context, spec *models.ProductSpecification, language *language.Language) (*models.ProductSpecificationLanguage, error) {
	return spec.ProductSpecificationLanguages(
		d.GetMods(
			models.ProductSpecificationLanguageWhere.LanguageID.EQ(language.L.ID),
		)...,
	).One(ctx, d.Db)
}

func (d *SpecificationDao) FindOneByCode(ctx context.Context, specCode string) (*models.ProductSpecification, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.SpecificationCode.EQ(specCode),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d *SpecificationDao) FindMandatoryByProduct(ctx context.Context, product *models.Product) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ProductCategoryID.EQ(product.ProductCategoryID),
			models.ProductSpecificationWhere.Mandatory.EQ(true),
		)...,
	).All(ctx, d.Db)
}

func (d *SpecificationDao) FindByProduct(ctx context.Context, product *models.Product) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ProductCategoryID.EQ(product.ProductCategoryID),
		)...,
	).All(ctx, d.Db)
}

func (d *SpecificationDao) FindByCategoryId(ctx context.Context, id int64) (models.ProductSpecificationSlice, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ProductCategoryID.EQ(id),
		)...,
	).All(ctx, d.Db)
}

func (d *SpecificationDao) FindById(ctx context.Context, id int64) (*models.ProductSpecification, error) {
	return models.ProductSpecifications(
		d.GetMods(
			models.ProductSpecificationWhere.ID.EQ(id),
		)...,
	).One(ctx, db.DB)
}

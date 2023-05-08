package product

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type CategoryDao struct {
	*db.Dao
}

func NewCategoryDao(executor boil.ContextExecutor) *CategoryDao {
	return &CategoryDao{
		db.DaoFromExecutor(executor),
	}
}

func (d CategoryDao) Clone() db.DaoMod {
	return CategoryDao{
		d.Dao.Clone(),
	}
}

func (d CategoryDao) Load(relationship string, mods ...qm.QueryMod) *CategoryDao {
	return db.Load(d, relationship, mods...)
}

func (d CategoryDao) Paginate(first int, offset int) *CategoryDao {
	return db.Paginate(d, first, offset)
}

func (d *CategoryDao) FindByCode(ctx context.Context, code string) (*models.ProductCategory, error) {
	return models.ProductCategories(
		d.GetMods(
			models.ProductCategoryWhere.CategoryCode.EQ(code),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d *CategoryDao) FindCategoryById(ctx context.Context, id int64) (*models.ProductCategory, error) {
	return models.ProductCategories(
		d.GetMods(
			models.ProductCategoryWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d *CategoryDao) FindAll(ctx context.Context) (models.ProductCategorySlice, error) {
	return models.ProductCategories(
		d.GetMods()...,
	).All(ctx, d.Db)
}

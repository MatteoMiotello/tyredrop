package product

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type CategoryDao struct {
	db.Dao
}

func NewCategoryDao(executor boil.ContextExecutor) *CategoryDao {
	return &CategoryDao{
		db.DaoFromExecutor(executor),
	}
}

func (d *CategoryDao) Load(relationship string, mods ...qm.QueryMod) *CategoryDao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *CategoryDao) Paginate(first int, offset int) *CategoryDao {
	db.Paginate(d, first, offset)

	return d
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

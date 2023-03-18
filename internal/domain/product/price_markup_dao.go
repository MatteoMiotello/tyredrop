package product

import (
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type PriceMarkupDao struct {
	*db.Dao
}

func NewPriceMarkupDao(executor boil.ContextExecutor) *PriceMarkupDao {
	return &PriceMarkupDao{
		db.DaoFromExecutor(executor),
	}
}

func (d *PriceMarkupDao) SetDao(dao *db.Dao) {
	d.Dao = dao
}

func (d *PriceMarkupDao) GetDao() *db.Dao {
	return d.Dao
}

func (d *PriceMarkupDao) Load(relationship string, mods ...qm.QueryMod) *PriceMarkupDao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *PriceMarkupDao) Paginate(first int, offset int) *PriceMarkupDao {
	db.Paginate(d, first, offset)

	return d
}

func (d *PriceMarkupDao) FindPriceMarkupByProductId(ctx context.Context, product *models.Product) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.ProductID.EQ(null.Int64From(product.ID)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d *PriceMarkupDao) FindPriceMarkupByBrandId(ctx context.Context, id int64) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.BrandID.EQ(null.Int64From(id)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d *PriceMarkupDao) FindPriceMarkupByProductCategoryId(ctx context.Context, categoryId int64) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.ProductCategoryID.EQ(null.Int64From(categoryId)),
			qm.Limit(1),
		)...,
	).One(ctx, d.Db)
}

func (d *PriceMarkupDao) FindPriceMarkupDefault(ctx context.Context) (*models.ProductPriceMarkup, error) {
	return models.ProductPriceMarkups(
		d.GetMods(
			models.ProductPriceMarkupWhere.ProductCategoryID.IsNull(),
			models.ProductPriceMarkupWhere.BrandID.IsNull(),
			models.ProductPriceMarkupWhere.ProductID.IsNull(),
		)...,
	).One(ctx, d.Db)
}

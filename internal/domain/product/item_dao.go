package product

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type ItemDao struct {
	db.Dao
}

func NewItemDao(exec boil.ContextExecutor) *ItemDao {
	return &ItemDao{
		db.DaoFromExecutor(exec),
	}
}

func (d *ItemDao) Load(relationship string, mods ...qm.QueryMod) *ItemDao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *ItemDao) Paginate(first int, offset int) *ItemDao {
	db.Paginate(d, first, offset)

	return d
}

func (d *ItemDao) Product(ctx context.Context, productItem *models.ProductItem) (*models.Product, error) {
	return productItem.Product(
		d.GetMods()...,
	).One(ctx, d.Db)
}

func (d *ItemDao) ProductItemPrices(ctx context.Context, productItem *models.ProductItem) (models.ProductItemPriceSlice, error) {
	return productItem.ProductItemPrices(
		d.GetMods()...,
	).All(ctx, d.Db)
}

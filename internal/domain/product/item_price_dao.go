package product

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type ItemPriceDao struct {
	*db.Dao
}

func NewItemPriceDao(executor boil.ContextExecutor) *ItemPriceDao {
	return &ItemPriceDao{
		db.DaoFromExecutor(executor),
	}
}

func (d *ItemPriceDao) SetDao(dao *db.Dao) {
	d.Dao = dao
}

func (d *ItemPriceDao) GetDao() *db.Dao {
	return d.Dao
}

func (d *ItemPriceDao) Load(relationship string, mods ...qm.QueryMod) *ItemPriceDao {
	db.Load(d, relationship, mods...)

	return d
}

func (d *ItemPriceDao) Paginate(first int, offset int) *ItemPriceDao {
	db.Paginate(d, first, offset)

	return d
}

func (d *ItemPriceDao) Currency(ctx context.Context, price *models.ProductItemPrice) (*models.Currency, error) {
	return price.Currency(
		d.GetMods()...,
	).One(ctx, d.Db)
}

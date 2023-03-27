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

func (d ItemPriceDao) Load(relationship string, mods ...qm.QueryMod) *ItemPriceDao {
	return db.Load(d, relationship, mods...)
}

func (d ItemPriceDao) Paginate(first int, offset int) *ItemPriceDao {
	return db.Paginate(d, first, offset)
}

func (d ItemPriceDao) Clone() db.DaoMod {
	return ItemPriceDao{
		d.Dao.Clone(),
	}
}

func (d *ItemPriceDao) Currency(ctx context.Context, price *models.ProductItemPrice) (*models.Currency, error) {
	return price.Currency(
		d.GetMods()...,
	).One(ctx, d.Db)
}

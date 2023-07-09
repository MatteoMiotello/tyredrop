package order

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type Dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		db.DaoFromExecutor(executor),
	}
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(limit int, offset int) *Dao {
	return db.Paginate(d, limit, offset)
}

func (d Dao) ForUpdate() *Dao {
	return db.ForUpdate(d)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Order, error) {
	return models.Orders(
		d.GetMods(
			models.OrderWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindAllOrderRowsByOrderId(ctx context.Context, id int64) (models.OrderRowSlice, error) {
	return models.OrderRows(
		d.GetMods(
			models.OrderRowWhere.OrderID.EQ(id),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindAllByBillingId(ctx context.Context, id int64) (models.OrderSlice, error) {
	return models.Orders(
		d.GetMods(
			models.OrderWhere.UserBillingID.EQ(id),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindDefaultTax(ctx context.Context) (*models.Taxis, error) {
	return models.Taxes(
		d.GetMods()...,
	).One(ctx, d.Db)
}

func (d Dao) GetUserBilling(ctx context.Context, order *models.Order) (*models.UserBilling, error) {
	return order.UserBilling(d.GetMods()...).One(ctx, d.Db)
}

package order

import (
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
	"strings"
	"time"
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

func (d Dao) Order(orderMods []*model.OrderingInput) *Dao {
	var mods []db.OrderMods

	for _, mod := range orderMods {
		mods = append(mods, db.OrderMods{Column: mod.Column, Desc: mod.Desc})
	}

	return db.Order(d, mods)
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

func (d Dao) FindAll(ctx context.Context, from *string, to *string, number *string, status *model.OrderStatus) (models.OrderSlice, error) {
	var mods []qm.QueryMod

	if from != nil && len(*from) > 0 {
		fromTime, err := time.Parse("2006-01-02", *from)

		if err != nil {
			return nil, err
		}

		mods = append(mods, models.OrderWhere.CreatedAt.GTE(fromTime))
	}

	if to != nil && len(*to) > 0 {
		toTime, err := time.Parse("2006-01-02", *to)

		if err != nil {
			return nil, err
		}

		mods = append(mods, models.OrderWhere.CreatedAt.LTE(toTime))
	}

	if number != nil && len(*number) > 0 {
		sanitizedNumber := strings.TrimLeft(*number, "#")

		mods = append(mods, models.OrderWhere.OrderNumber.EQ(null.StringFrom(sanitizedNumber)))
	}

	if status != nil {
		mods = append(mods, models.OrderWhere.Status.EQ(status.String()))
	}

	return models.Orders(
		d.GetMods(
			mods...,
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindAllOrderRowsByOrderId(ctx context.Context, id int64) (models.OrderRowSlice, error) {
	return models.OrderRows(
		d.GetMods(
			models.OrderRowWhere.OrderID.EQ(id),
		)...,
	).All(ctx, d.Db)
}

func (d Dao) FindOrderRowById(ctx context.Context, id int64) (*models.OrderRow, error) {
	return models.OrderRows(
		d.GetMods(
			models.OrderRowWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindAllByBillingId(ctx context.Context, id int64, from *string, to *string, number *string) (models.OrderSlice, error) {
	var mods []qm.QueryMod

	if from != nil && len(*from) > 0 {
		fromTime, err := time.Parse("2006-01-02", *from)

		if err != nil {
			return nil, err
		}

		mods = append(mods, models.OrderWhere.CreatedAt.GTE(fromTime))
	}

	if to != nil && len(*to) > 0 {
		toTime, err := time.Parse("2006-01-02", *to)

		if err != nil {
			return nil, err
		}

		mods = append(mods, models.OrderWhere.CreatedAt.LTE(toTime))
	}

	if number != nil && len(*number) > 0 {
		sanitizedNumber := strings.TrimLeft(*number, "#")

		mods = append(mods, models.OrderWhere.OrderNumber.EQ(null.StringFrom(sanitizedNumber)))
	}

	mods = append(mods, models.OrderWhere.UserBillingID.EQ(id))

	return models.Orders(
		d.GetMods(
			mods...,
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

type TotalOrders struct {
	TotalPrice int `boil:"price_amount"`
}

func (d Dao) UserTotalOrders(ctx context.Context, billing *models.UserBilling, from time.Time, to time.Time) (*TotalOrders, error) {
	tp := new(TotalOrders)

	err := models.Orders(
		d.GetMods(
			qm.Select("SUM( price_amount ) as price_amount"),
			models.OrderWhere.UserBillingID.EQ(billing.ID),
			models.OrderWhere.CreatedAt.GTE(from),
			models.OrderWhere.CreatedAt.LTE(to),
			models.OrderWhere.Status.IN(model.OrderProcessedStatusCollection),
		)...,
	).Bind(ctx, d.Db, tp)

	if err != nil {
		return nil, err
	}

	return tp, nil
}

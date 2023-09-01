package invoice

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
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

func (d Dao) ForUpdate() *Dao {
	return db.ForUpdate(d)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Invoice, error) {
	return models.Invoices(
		d.GetMods(
			models.InvoiceWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindAll(ctx context.Context, userBillingId *int64, from *string, to *string, number *string) (models.InvoiceSlice, error) {
	var mods []qm.QueryMod

	mods = append(mods, qm.OrderBy(models.InvoiceColumns.CreatedAt+" DESC"))

	if userBillingId != nil {
		mods = append(mods, models.InvoiceWhere.UserBillingID.EQ(*userBillingId))
	}

	if from != nil {
		fromTime, err := time.Parse("2006-01-02", *from)

		if err != nil {
			return nil, err
		}

		mods = append(mods, models.InvoiceWhere.CreatedAt.GTE(fromTime))
	}

	if to != nil {
		toTime, err := time.Parse("2006-01-02", *to)

		if err != nil {
			return nil, err
		}

		mods = append(mods, models.InvoiceWhere.CreatedAt.LTE(toTime))
	}

	if number != nil {
		mods = append(mods, qm.Where("invoices.number LIKE ?", fmt.Sprintf("%%%s%%", *number)))
	}

	return models.Invoices(d.GetMods(
		mods...,
	)...).
		All(ctx, d.Db)
}

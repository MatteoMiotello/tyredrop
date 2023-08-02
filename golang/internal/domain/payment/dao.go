package payment

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

func (d Dao) SetDao(dao *db.Dao) db.DaoMod {
	d.Dao = dao
	return d
}

func (d Dao) GetDao() *db.Dao {
	return d.Dao
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(first int, offset int) *Dao {
	return db.Paginate(d, first, offset)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d Dao) FindPaymentMethodByCode(ctx context.Context, code string) (*models.PaymentMethod, error) {
	return models.PaymentMethods(
		d.GetMods(models.PaymentMethodWhere.Code.EQ(code))...,
	).One(ctx, d.Db)
}

func (d Dao) FindUserPaymentMethodById(ctx context.Context, id int64) (*models.UserPaymentMethod, error) {
	return models.UserPaymentMethods(
		d.GetMods(
			models.UserPaymentMethodWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindUserPaymentMethodByUserAndMethodId(ctx context.Context, userId int64, methodId int64) (*models.UserPaymentMethod, error) {
	return models.UserPaymentMethods(
		d.GetMods(
			models.UserPaymentMethodWhere.UserID.EQ(userId),
			models.UserPaymentMethodWhere.PaymentMethodID.EQ(methodId),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindAllPaymentMethods(ctx context.Context) (models.PaymentMethodSlice, error) {
	return models.PaymentMethods(
		d.GetMods()...,
	).All(ctx, d.Db)
}

func (d Dao) FindOneById(ctx context.Context, id int64) (*models.Payment, error) {
	return models.Payments(
		d.GetMods(models.PaymentWhere.ID.EQ(id))...,
	).One(ctx, d.Db)
}

func (d Dao) FindPaymentMethodById(ctx context.Context, id int64) (*models.PaymentMethod, error) {
	return models.PaymentMethods(d.GetMods(
		models.PaymentMethodWhere.ID.EQ(id),
	)...).One(ctx, d.Db)
}

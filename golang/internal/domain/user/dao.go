package user

import (
	context2 "context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/net/context"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
	"time"
)

type RoleSet string

const (
	ADMIN_ROLE    RoleSet = "ADMIN"
	SUPPLIER_ROLE RoleSet = "SUPPLIER"
	USER_ROLE     RoleSet = "USER"
)

type Dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		db.DaoFromExecutor(executor),
	}
}

func (u Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(u, relationship, mods...)
}

func (u Dao) Paginate(limit int, offset int) *Dao {
	return db.Paginate(u, limit, offset)
}

func (u Dao) ForUpdate() *Dao {
	return db.ForUpdate(u)
}

func (u Dao) Clone() db.DaoMod {
	return Dao{
		u.Dao.Clone(),
	}
}

func (u Dao) FindOneByUsername(ctx context.Context, username string) (*models.User, error) {
	return models.Users(qm.Where("username = ?", username), qm.Load(models.UserRels.UserRole)).One(ctx, u.Db)
}

func (u Dao) FindOneByEmail(ctx context.Context, email string) (*models.User, error) {
	return models.Users(qm.Where("email = ?", email), qm.Load(models.UserRels.UserRole)).One(ctx, u.Db)
}

func (u Dao) FindOneById(ctx context.Context, id int64) (*models.User, error) {
	return models.Users(
		u.GetMods(
			models.UserWhere.ID.EQ(id),
		)...,
	).One(ctx, u.Db)
}

func (u Dao) GetUserRole(ctx context.Context, user *models.User) (*models.UserRole, error) {
	return user.UserRole().One(ctx, u.Db)
}

func (u Dao) GetDefaultLanguage(ctx context.Context, user models.User) (*models.Language, error) {
	return user.DefaultLanguage().One(ctx, u.Db)
}

func (u Dao) FindUserRoleByCode(ctx context.Context, roleCode RoleSet) (*models.UserRole, error) {
	return models.UserRoles(qm.Where("role_code = ?", string(roleCode))).One(ctx, u.Db)
}

func (u Dao) FindUserRoleById(ctx context.Context, id int64) (*models.UserRole, error) {
	return models.UserRoles(
		u.GetMods(
			models.UserRoleWhere.ID.EQ(id),
		)...,
	).One(ctx, u.Db)
}

func (u Dao) GetUserRoleLanguage(ctx context.Context, role *models.UserRole, language models.Language) (*models.UserRoleLanguage, error) {
	l, err := role.UserRoleLanguages(qm.Where(models.UserRoleLanguageColumns.LanguageID+"= ?", language.ID)).One(ctx, u.Db)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (u Dao) GetUserBilling(ctx context.Context, user *models.User) (*models.UserBilling, error) {
	return models.UserBillings(
		u.GetMods(
			models.UserBillingWhere.UserID.EQ(user.ID),
			qm.OrderBy(models.UserBillingColumns.CreatedAt+" DESC"),
		)...,
	).One(ctx, u.Db)
}

func (d Dao) FindUserBillingById(ctx context.Context, id int64) (*models.UserBilling, error) {
	return models.UserBillings(
		d.GetMods(
			models.UserBillingWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d Dao) FindAll(ctx context.Context, email *string, name *string, confirmed *bool) (models.UserSlice, error) {
	var mods []qm.QueryMod

	if email != nil {
		mods = append(mods, qm.Where("email LIKE ?", fmt.Sprintf("%%%s%%", *email)))
	}

	if name != nil {
		mods = append(mods, qm.Where("( name || ' ' || surname ILIKE ? ) ", fmt.Sprintf("%%%s%%", *name)))
	}

	if confirmed != nil {
		mods = append(mods, models.UserWhere.Confirmed.EQ(*confirmed))
	}

	mods = append(mods, qm.OrderBy(models.UserColumns.ID+" DESC"))

	return models.Users(
		d.GetMods(
			mods...,
		)...,
	).All(ctx, d.Db)
}

func (d Dao) TotalUsers(ctx context.Context) (int64, error) {
	return models.Users(
		d.GetMods(
			models.UserWhere.Confirmed.EQ(true),
			qm.WhereIn(models.UserColumns.UserRoleID+" IN ( SELECT id FROM public.user_roles WHERE role_code = ? )", "USER"),
		)...,
	).Count(ctx, d.Db)
}

func (d Dao) BestUserBilling(ctx context2.Context, from time.Time, to time.Time) (*models.UserBilling, error) {
	return models.UserBillings(
		d.GetMods(
			qm.LeftOuterJoin("orders on orders.user_billing_id = user_billings.id"),
			models.OrderWhere.CreatedAt.GTE(from),
			models.OrderWhere.CreatedAt.LTE(to),
			models.OrderWhere.Status.IN(model.OrderProcessedStatusCollection),
			qm.GroupBy("user_billings.id"),
			qm.OrderBy("sum( orders.price_amount ) DESC"),
		)...,
	).One(ctx, d.Db)
}

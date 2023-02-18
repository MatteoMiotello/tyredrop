package user

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/net/context"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type RoleSet string

const (
	ADMIN_ROLE    RoleSet = "ADMIN"
	SUPPLIER_ROLE RoleSet = "SUPPLIER"
	USER_ROLE     RoleSet = "USER"
)

type Dao struct {
	db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		Dao: db.DaoFromExecutor(executor),
	}
}

func (u Dao) FindOneByUsername(ctx context.Context, username string) (*models.User, error) {
	return models.Users(qm.Where("username = ?", username), qm.Load(models.UserRels.UserRole)).One(ctx, u.Db)
}

func (u Dao) FindOneByEmail(ctx context.Context, email string) (*models.User, error) {
	return models.Users(qm.Where("email = ?", email), qm.Load(models.UserRels.UserRole)).One(ctx, u.Db)
}

func (u Dao) FindOneById(ctx context.Context, id int64) (*models.User, error) {
	return models.FindUser(ctx, u.Db, id)
}

func (u Dao) Insert(ctx context.Context, user *models.User) error {
	return user.Insert(ctx, u.Db, boil.Infer())
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

func (u Dao) GetUserRoleLanguage(ctx context.Context, role *models.UserRole, language models.Language) (*models.UserRoleLanguage, error) {
	l, err := role.UserRoleLanguages(qm.Where(models.UserRoleLanguageColumns.LanguageID+"= ?", language.ID)).One(ctx, u.Db)
	if err != nil {
		return nil, err
	}

	return l, nil
}

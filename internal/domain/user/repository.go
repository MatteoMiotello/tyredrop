package user

import (
	"database/sql"
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

type repo db.Repo

func NewUserRepo(db *sql.DB) *repo {
	return &repo{
		Db: db,
	}
}

func (u repo) FindOneByUsername(ctx context.Context, username string) (*models.User, error) {
	return models.Users(qm.Where("username = ?", username), qm.Load(models.UserRels.UserRole)).One(ctx, u.Db)
}

func (u repo) FindOneByEmail(ctx context.Context, email string) (*models.User, error) {
	return models.Users(qm.Where("email = ?", email), qm.Load(models.UserRels.UserRole)).One(ctx, u.Db)
}

func (u repo) FindOneById(ctx context.Context, id int64) (*models.User, error) {
	return models.FindUser(ctx, u.Db, id)
}

func (u repo) Insert(ctx context.Context, user *models.User) error {
	return user.Insert(ctx, u.Db, boil.Infer())
}

func (u repo) GetUserRole(ctx context.Context, user *models.User) (*models.UserRole, error) {
	return user.UserRole().One(ctx, u.Db)
}

func (u repo) GetDefaultLanguage(ctx context.Context, user models.User) (*models.Language, error) {
	return user.DefaultLanguage().One(ctx, u.Db)
}

func (u repo) FindUserRoleByCode(ctx context.Context, roleCode RoleSet) (*models.UserRole, error) {
	return models.UserRoles(qm.Where("role_code = ?", string(roleCode))).One(ctx, u.Db)
}

func (u repo) GetUserRoleLanguage(ctx context.Context, role *models.UserRole, language models.Language) (*models.UserRoleLanguage, error) {
	l, err := role.UserRoleLanguages(qm.Where(models.UserRoleLanguageColumns.LanguageID+"= ?", language.ID)).One(ctx, u.Db)
	if err != nil {
		return nil, err
	}

	return l, nil
}

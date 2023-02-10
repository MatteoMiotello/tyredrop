package repositories

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type UserRoleSet string

const (
	ADMIN_ROLE    UserRoleSet = "ADMIN"
	SUPPLIER_ROLE UserRoleSet = "SUPPLIER"
	USER_ROLE     UserRoleSet = "USER"
)

type UserRoleRepo DbRepo

func NewUserRoleRepoFromCtx(ctx context.Context) *UserRoleRepo {
	return &UserRoleRepo{
		context: ctx,
	}
}

func (u UserRoleRepo) FindByRoleCode(roleCode UserRoleSet) (*models.UserRole, error) {
	return models.UserRoles(qm.Where("role_code = ?", string(roleCode))).One(u.context, db.DB)
}

func (u UserRoleRepo) GetLanguage(role *models.UserRole, language models.Language) (*models.UserRoleLanguage, error) {
	l, err := role.UserRoleLanguages(qm.Where(models.UserRoleLanguageColumns.LanguageID+"= ?", language.ID)).One(u.context, db.DB)
	if err != nil {
		return nil, err
	}

	return l, nil
}

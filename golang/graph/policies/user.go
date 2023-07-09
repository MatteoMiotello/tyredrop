package policies

import (
	"context"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
)

type userPolicy struct {
	model    *models.User
	orderDao *user.Dao
}

func NewUserPolicy(uModel *models.User, dao *user.Dao) *userPolicy {
	return &userPolicy{
		model:    uModel,
		orderDao: dao,
	}
}

func (u *userPolicy) CanRead(ctx context.Context) bool {
	currentUser, err := auth.CurrentUser(ctx)

	if err != nil {
		return false
	}

	if currentUser.R.UserRole.Admin {
		return true
	}

	return currentUser.ID == u.model.ID
}

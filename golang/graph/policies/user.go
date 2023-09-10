package policies

import (
	"context"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
)

type userPolicy struct {
	currentUser *models.User
	orderDao    *user.Dao
}

func NewUserPolicy(dao *user.Dao) *userPolicy {
	return &userPolicy{
		orderDao: dao,
	}
}

func (u *userPolicy) CanRead(ctx context.Context, user *models.User) bool {
	currentUser, err := auth.CurrentUser(ctx)

	if err != nil {
		return false
	}

	if currentUser.R.UserRole.Admin {
		return true
	}

	return currentUser.ID == user.ID
}

func (u *userPolicy) CanUpdateAvatar(ctx context.Context, user *models.User) bool {
	currentUser, err := auth.CurrentUser(ctx)

	if err != nil {
		return false
	}

	if currentUser.R.UserRole.Admin {
		return true
	}

	if currentUser.ID != user.ID {
		return false
	}

	return true
}

func (u *userPolicy) CanUpdateBilling(ctx context.Context, billing *models.UserBilling) bool {
	cu, err := auth.CurrentUser(ctx)

	if err != nil {
		return false
	}

	if cu.R.UserRole.Admin {
		return true
	}

	if billing.UserID != cu.ID {
		return false
	}

	return true
}

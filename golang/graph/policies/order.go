package policies

import (
	"context"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/domain/order"
	"pillowww/titw/models"
)

type orderPolicy struct {
	model    *models.Order
	orderDao *order.Dao
}

func NewOrderPolicy(order *models.Order, dao *order.Dao) *orderPolicy {
	return &orderPolicy{
		model:    order,
		orderDao: dao,
	}
}

func (p orderPolicy) CanRead(ctx context.Context) bool {
	user, err := auth2.CurrentUser(ctx)

	if err != nil {
		return false
	}

	if user.R.UserRole.Admin {
		return true
	}

	billing, err := p.orderDao.GetUserBilling(ctx, p.model)
	if err != nil {
		return false
	}

	return billing.UserID == user.ID
}

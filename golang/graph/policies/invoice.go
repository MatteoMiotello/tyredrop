package policies

import (
	"context"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/domain/user"
)

type InvoicePolicy struct {
	uDao *user.Dao
}

func NewInvoicePolicy(uDao *user.Dao) *InvoicePolicy {
	return &InvoicePolicy{
		uDao: uDao,
	}
}

func (p InvoicePolicy) CanViewAll(ctx context.Context, billingId *int64) bool {
	currentUser, err := auth.CurrentUser(ctx)

	if err != nil {
		return false
	}

	if currentUser.R.UserRole.Admin {
		return true
	}

	if billingId == nil {
		return false
	}

	b, err := p.uDao.FindUserBillingById(ctx, *billingId)

	if err != nil {
		return false
	}

	if b.UserID != currentUser.ID {
		return false
	}

	return true
}

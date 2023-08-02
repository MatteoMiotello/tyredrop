package payment

import (
	"context"
	"github.com/friendsofgo/errors"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/domain/order"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
)

type Service struct {
	paymentDao *Dao
	orderDao   *order.Dao
}

func NewService(pDao *Dao, oDao *order.Dao) *Service {
	return &Service{
		pDao,
		oDao,
	}
}

func (s Service) CreatePayment(ctx context.Context, order *models.Order, method *models.PaymentMethod) (*models.Payment, error) {
	if order.Status != model.OrderStatusNotCompleted.String() {
		return nil, errors.New("order cannot be payed in this status: " + order.Status)
	}

	if !order.PaymentID.IsZero() {
		return nil, errors.New("order is already payed")
	}

	billing, err := s.orderDao.GetUserBilling(ctx, order)

	if err != nil {
		return nil, err
	}

	uPm, _ := s.paymentDao.FindUserPaymentMethodByUserAndMethodId(ctx, billing.UserID, method.ID)

	if uPm == nil {
		if method.Code == constants.PAYMENT_METHOD_BANK_TRANSFER {
			uPm = &models.UserPaymentMethod{
				UserID:          billing.UserID,
				Name:            billing.Name + " " + billing.Surname.String,
				PaymentMethodID: method.ID,
			}

			err := s.paymentDao.Save(ctx, uPm)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Payment method not found for user")
		}
	}

	payment := &models.Payment{
		UserPaymentMethodID: uPm.ID,
		CurrencyID:          order.CurrencyID,
		Amount:              order.PriceAmountTotal,
	}

	err = s.paymentDao.Save(ctx, payment)

	if err != nil {
		return nil, err
	}

	return payment, nil
}

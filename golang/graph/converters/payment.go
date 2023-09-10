package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
)

func PaymentToGraphQL(payment *models.Payment) (*model.Payment, error) {
	curr := payment.R.Currency

	if curr == nil {
		panic("currency not loaded on payment")
	}

	floatVal, err := currency.ToFloat(payment.Amount, curr.IsoCode)

	if err != nil {
		return nil, err
	}

	return &model.Payment{
		ID:                  payment.ID,
		CurrencyID:          payment.CurrencyID,
		UserPaymentMethodID: payment.UserPaymentMethodID,
		Amount:              floatVal,
	}, nil
}

func PaymentMethodToGraphQL(method *models.PaymentMethod) *model.PaymentMethod {
	return &model.PaymentMethod{
		ID:       method.ID,
		Code:     method.Code,
		Name:     method.Name,
		Receiver: method.Receiver.Ptr(),
		BankName: method.BankName.Ptr(),
		Iban:     method.Iban.Ptr(),
	}
}

func UserPaymentMethodToGraphQL(m *models.UserPaymentMethod) *model.UserPaymentMethod {
	return &model.UserPaymentMethod{
		ID:              m.ID,
		PaymentMethodID: m.PaymentMethodID,
		Name:            m.Name,
		Value:           m.Value.Ptr(),
		TypePrimary:     m.TypePrimary,
	}
}

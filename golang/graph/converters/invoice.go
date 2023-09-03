package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func InvoiceToGraphQL(invoice *models.Invoice) *model.Invoice {
	status := model.InvoiceStatusPayed

	if invoice.Status.IsZero() {
		status = model.InvoiceStatusToPay
	}

	if invoice.Status.String == model.InvoiceStatusToPay.String() {
		status = model.InvoiceStatusToPay
	}

	return &model.Invoice{
		ID:            invoice.ID,
		Number:        invoice.Number,
		FilePath:      invoice.FilePath,
		UserBillingId: invoice.UserBillingID,
		CreatedAt:     invoice.CreatedAt,
		Status:        status,
	}
}

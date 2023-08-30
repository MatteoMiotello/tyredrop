package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func InvoiceToGraphQL(invoice *models.Invoice) *model.Invoice {
	return &model.Invoice{
		ID:            invoice.ID,
		Number:        invoice.Number,
		FilePath:      invoice.FilePath,
		UserBillingId: invoice.UserBillingID,
		CreatedAt:     invoice.CreatedAt,
	}
}

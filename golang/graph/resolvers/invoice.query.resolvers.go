package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
)

// AllInvoices is the resolver for the allInvoices field.
func (r *queryResolver) AllInvoices(ctx context.Context, pagination model.PaginationInput, input model.InvoiceFilter) (*model.InvoicePaginator, error) {
	allInvoices, err := r.InvoiceDao.FindAll(ctx, input.UserBillingID, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	invoices, err := r.InvoiceDao.
		Paginate(pagination.Limit, pagination.Offset).
		FindAll(ctx, input.UserBillingID, input.From, input.To, input.Number)

	if err != nil {
		return nil, err
	}

	var graphModels []*model.Invoice

	for _, i := range invoices {
		graphModels = append(graphModels, converters.InvoiceToGraphQL(i))
	}

	return &model.InvoicePaginator{
		Pagination: converters.PaginationToGraphql(&pagination, len(allInvoices)),
		Data:       graphModels,
	}, nil
}

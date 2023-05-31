package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"database/sql"
	"fmt"
	"pillowww/titw/graph"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/graphErrors"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/cart"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	null "github.com/volatiletech/null/v8"
)

// CreateAdminUser is the resolver for the createAdminUser field.
func (r *mutationResolver) CreateAdminUser(ctx context.Context, userInput model.CreateAdminUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateAdminUser - createAdminUser"))
}

// CreateUserBilling is the resolver for the createUserBilling field.
func (r *mutationResolver) CreateUserBilling(ctx context.Context, billingInput model.CreateUserBilling) (*model.UserBilling, error) {
	a := auth.FromCtx(ctx)
	currentUser, err := a.GetUser(ctx)

	if err != nil {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "User not found",
			Extensions: map[string]interface{}{
				"code": "4004",
			},
		}
	}

	var userBillingModel *models.UserBilling = new(models.UserBilling)
	var paymentMethod *models.UserPaymentMethod = new(models.UserPaymentMethod)

	err = db.WithTx(ctx, func(tx *sql.Tx) error {
		userDao := user.NewDao(tx)
		userBillingModel.UserID = currentUser.ID
		userBillingModel.LegalEntityTypeID = billingInput.LegalEntityTypeID
		userBillingModel.Name = billingInput.Name

		if billingInput.Surname != nil {
			userBillingModel.Surname = null.StringFrom(*billingInput.Surname)
		}

		userBillingModel.VatNumber = billingInput.VatNumber
		userBillingModel.AddressLine1 = billingInput.AddressLine1

		if billingInput.AddressLine2 != nil {
			userBillingModel.AddressLine2 = null.StringFrom(*billingInput.AddressLine2)
		}

		userBillingModel.City = billingInput.City
		userBillingModel.Province = billingInput.Province
		userBillingModel.Cap = billingInput.Cap
		userBillingModel.Country = billingInput.Country

		if billingInput.FiscalCode == nil {
			userBillingModel.FiscalCode = billingInput.VatNumber
		} else {
			userBillingModel.FiscalCode = *billingInput.FiscalCode
		}

		err = userDao.Insert(ctx, userBillingModel)
		if err != nil {
			return graphErrors.NewGraphError(ctx, err, "5001")
		}

		paymentMethod.UserID = currentUser.ID
		paymentMethod.Name = billingInput.Name

		if billingInput.Surname != nil {
			paymentMethod.Name += " " + *billingInput.Surname
		}

		paymentMethod.Type = constants.PAYMENT_METHOD_SEPA
		paymentMethod.Value = billingInput.Iban
		paymentMethod.TypePrimary = true

		err = userDao.Insert(ctx, paymentMethod)

		if err != nil {
			return graphErrors.NewGraphError(ctx, err, "5001")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return converters.UserBillingToGraphQL(userBillingModel), nil
}

// AddItemToCart is the resolver for the addItemToCart field.
func (r *mutationResolver) AddItemToCart(ctx context.Context, itemID int64, quantity *int) ([]*model.Cart, error) {
	a := auth.FromCtx(ctx)
	u, err := a.GetUser(ctx)

	if err != nil {
		return nil, gqlerror.Errorf("User not found in context")
	}

	s := cart.NewCartService(r.CartDao)

	_, err = s.AddOrUpdateCart(ctx, u, itemID, quantity)

	if err != nil {
		return nil, err
	}

	cartModels, err := r.CartDao.FindAllByUserId(ctx, u.ID)

	var graphModels []*model.Cart

	for _, cart := range cartModels {
		graphModels = append(graphModels, converters.CartToGraphQL(cart))
	}

	return graphModels, err
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

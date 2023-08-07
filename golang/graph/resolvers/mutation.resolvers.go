package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pillowww/titw/graph"
	"pillowww/titw/graph/aggregators"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/graphErrors"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/cart"
	"pillowww/titw/internal/domain/order"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/internal/domain/user"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"

	"github.com/vektah/gqlparser/v2/gqlerror"
	null "github.com/volatiletech/null/v8"
)

// CreateAdminUser is the resolver for the createAdminUser field.
func (r *mutationResolver) CreateAdminUser(ctx context.Context, userInput model.CreateAdminUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateAdminUser - createAdminUser"))
}

// CreateUserBilling is the resolver for the createUserBilling field.
func (r *mutationResolver) CreateUserBilling(ctx context.Context, billingInput model.CreateUserBilling) (*model.UserBilling, error) {
	currentUser, err := auth.CurrentUser(ctx)

	if err != nil {
		return nil, graphErrors.NewUserNotFoundError(ctx, err)
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
		userBillingModel.SdiCode = null.StringFromPtr(billingInput.SdiCode)
		userBillingModel.SdiPec = null.StringFromPtr(billingInput.SdiPec)

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

		method, err := r.PaymentDao.FindPaymentMethodByCode(ctx, constants.PAYMENT_METHOD_SEPA)

		if err != nil {
			return err
		}

		paymentMethod.PaymentMethodID = method.ID
		paymentMethod.Value = null.StringFrom(billingInput.Iban)
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
func (r *mutationResolver) AddItemToCart(ctx context.Context, itemID int64, quantity *int) (*model.CartResponse, error) {
	a := auth.FromCtx(ctx)
	u, err := a.GetUser(ctx)

	if err != nil {
		return nil, gqlerror.Errorf("User not found in context")
	}

	s := cart.NewCartService(r.CartDao, r.ProductItemPriceDao)

	_, err = s.AddOrUpdateCart(ctx, u, itemID, quantity)

	if err != nil {
		return nil, err
	}

	return aggregators.GetAllCartsByUserId(ctx, r.CartDao, u.ID)
}

// EditCart is the resolver for the editCart field.
func (r *mutationResolver) EditCart(ctx context.Context, cartID int64, quantity int) (*model.CartResponse, error) {
	c, err := r.CartDao.FindOneById(ctx, cartID)

	if err != nil {
		return nil, err
	}

	if quantity < 0 {
		return nil, graphErrors.NewGraphError(ctx, errors.New("Quantity must be grater or equal than 0"), "6000")
	}

	if quantity == 0 {
		err := r.CartDao.Delete(ctx, c)
		if err != nil {
			return nil, err
		}

		return aggregators.GetAllCartsByUserId(ctx, r.CartDao, c.UserID)
	}

	c.Quantity = quantity
	err = r.CartDao.Update(ctx, c)
	if err != nil {
		return nil, err
	}

	return aggregators.GetAllCartsByUserId(ctx, r.CartDao, c.UserID)
}

// EmptyCart is the resolver for the emptyCart field.
func (r *mutationResolver) EmptyCart(ctx context.Context) (*model.CartResponse, error) {
	u, err := auth.CurrentUser(ctx)

	if err != nil {
		return nil, err
	}

	carts, err := r.CartDao.FindAllByUserId(ctx, u.ID)

	if err != nil {
		return nil, err
	}

	for _, c := range carts {
		err := r.CartDao.Delete(ctx, c)

		if err != nil {
			return nil, err
		}
	}

	return aggregators.GetAllCartsByUserId(ctx, r.CartDao, u.ID)
}

// CreateUserAddress is the resolver for the createUserAddress field.
func (r *mutationResolver) CreateUserAddress(ctx context.Context, userAddress model.UserAddressInput) ([]*model.UserAddress, error) {
	u, err := auth.CurrentUser(ctx)

	if err != nil {
		return nil, graphErrors.NewUserNotFoundError(ctx, err)
	}

	uAddressModel := &models.UserAddress{}
	uAddressModel.UserID = u.ID

	converters.GraphQLToUserAddress(userAddress, uAddressModel)

	err = r.UserAddressDao.Insert(ctx, uAddressModel)
	if err != nil {
		return nil, err
	}

	return aggregators.NewUserAggregator(r.UserAddressDao).GetAllAddressesByUser(ctx, u.ID)
}

// EditUserAddress is the resolver for the editUserAddress field.
func (r *mutationResolver) EditUserAddress(ctx context.Context, id int64, userAddress model.UserAddressInput) ([]*model.UserAddress, error) {
	u, err := auth.CurrentUser(ctx)

	if err != nil {
		return nil, graphErrors.NewUserNotFoundError(ctx, err)
	}

	uAddressModel, err := r.UserAddressDao.FindOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	if uAddressModel.UserID != u.ID {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	converters.GraphQLToUserAddress(userAddress, uAddressModel)

	err = r.UserAddressDao.Update(ctx, uAddressModel)
	if err != nil {
		return nil, err
	}

	return aggregators.NewUserAggregator(r.UserAddressDao).GetAllAddressesByUser(ctx, u.ID)
}

// DeleteUserAddress is the resolver for the deleteUserAddress field.
func (r *mutationResolver) DeleteUserAddress(ctx context.Context, id int64) ([]*model.UserAddress, error) {
	u, err := auth.CurrentUser(ctx)

	if err != nil {
		return nil, graphErrors.NewUserNotFoundError(ctx, err)
	}

	uAddressModel, err := r.UserAddressDao.FindOneById(ctx, id)

	if err != nil {
		return nil, err
	}

	if uAddressModel.UserID != u.ID {
		return nil, graphErrors.NewNotAuthorizedError(ctx)
	}

	err = r.ProductItemDao.Delete(ctx, uAddressModel)
	if err != nil {
		return nil, err
	}

	return aggregators.NewUserAggregator(r.UserAddressDao).GetAllAddressesByUser(ctx, u.ID)
}

// NewOrder is the resolver for the newOrder field.
func (r *mutationResolver) NewOrder(ctx context.Context, userID int64, userAddressID int64) (*model.Order, error) {
	uModel, err := r.UserDao.FindOneById(ctx, userID)

	if err != nil {
		return nil, err
	}

	uBilling, err := r.UserDao.GetUserBilling(ctx, uModel)

	if err != nil {
		return nil, err
	}

	carts, err := r.CartDao.FindAllByUserId(ctx, userID)

	if err != nil {
		return nil, err
	}

	uAddressModel, err := r.UserAddressDao.FindOneById(ctx, userAddressID)

	if err != nil {
		return nil, err
	}

	var newOrder *models.Order

	err = db.WithTx(ctx, func(tx *sql.Tx) error {
		oService := order.NewService(
			order.NewDao(tx),
			currency.NewDao(tx),
			product.NewItemDao(tx),
			product.NewItemPriceDao(tx),
		)

		newOrder, err = oService.CreateNewOrder(ctx, uBilling, uAddressModel, carts)

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	newOrder, err = r.OrderDao.Load(models.OrderRels.Currency).FindOneById(ctx, newOrder.ID)

	if err != nil {
		return nil, err
	}

	return converters.OrderToGraphQL(newOrder)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

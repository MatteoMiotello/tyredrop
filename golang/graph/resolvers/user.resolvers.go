package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"pillowww/titw/graph"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/models"
)

// UserRole is the resolver for the userRole field.
func (r *userResolver) UserRole(ctx context.Context, obj *model.User) (*model.UserRole, error) {
	lang := auth2.CurrentLanguage(ctx)

	roleModel, err := r.UserDao.
		Load(
			models.UserRoleRels.UserRoleLanguages,
			models.UserRoleLanguageWhere.LanguageID.EQ(lang.L.ID),
		).
		FindUserRoleById(ctx, obj.UserRoleID)

	if err != nil {
		return nil, err
	}

	return converters.UserRoleToGraphQL(roleModel), nil
}

// UserBilling is the resolver for the userBilling field.
func (r *userResolver) UserBilling(ctx context.Context, obj *model.User) (*model.UserBilling, error) {
	uModel, err := r.UserDao.FindOneById(ctx, obj.ID)

	if err != nil {
		return nil, err
	}

	billing, err := r.UserDao.GetUserBilling(ctx, uModel)

	if billing == nil {
		return nil, nil
	}

	return converters.UserBillingToGraphQL(billing), nil
}

// User is the resolver for the user field.
func (r *userAddressResolver) User(ctx context.Context, obj *model.UserAddress) (*model.User, error) {
	uModel, err := r.UserDao.FindOneById(ctx, obj.ID)

	if err != nil {
		return nil, err
	}

	return converters.UserToGraphQL(uModel), nil
}

// LegalEntityType is the resolver for the legalEntityType field.
func (r *userBillingResolver) LegalEntityType(ctx context.Context, obj *model.UserBilling) (*model.LegalEntityType, error) {
	entityType, err := r.LegalEntityDao.FindOneById(ctx, obj.LegalEntityTypeID)
	if err != nil {
		return nil, err
	}

	return converters.LegalEntityTypeToGraphQL(*entityType), nil
}

// TaxRate is the resolver for the taxRate field.
func (r *userBillingResolver) TaxRate(ctx context.Context, obj *model.UserBilling) (*model.Tax, error) {
	panic(fmt.Errorf("not implemented: TaxRate - taxRate"))
}

// User is the resolver for the user field.
func (r *userBillingResolver) User(ctx context.Context, obj *model.UserBilling) (*model.User, error) {
	user, err := r.UserDao.FindOneById(ctx, obj.UserID)

	if err != nil {
		return nil, err
	}

	return converters.UserToGraphQL(user), nil
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

// UserAddress returns graph.UserAddressResolver implementation.
func (r *Resolver) UserAddress() graph.UserAddressResolver { return &userAddressResolver{r} }

// UserBilling returns graph.UserBillingResolver implementation.
func (r *Resolver) UserBilling() graph.UserBillingResolver { return &userBillingResolver{r} }

type userResolver struct{ *Resolver }
type userAddressResolver struct{ *Resolver }
type userBillingResolver struct{ *Resolver }

package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"pillowww/titw/graph"
	"pillowww/titw/graph/model"
)

// CreateUserBilling is the resolver for the createUserBilling field.
func (r *mutationResolver) CreateUserBilling(ctx context.Context, input *model.CreateUserBilling) (*model.UserBilling, error) {
	panic(fmt.Errorf("not implemented: CreateUserBilling - createUserBilling"))
}

// UserRole is the resolver for the userRole field.
func (r *userResolver) UserRole(ctx context.Context, obj *model.User) (*model.UserRole, error) {
	panic(fmt.Errorf("not implemented: UserRole - userRole"))
}

// UserBilling is the resolver for the userBilling field.
func (r *userResolver) UserBilling(ctx context.Context, obj *model.User) (*model.UserBilling, error) {
	panic(fmt.Errorf("not implemented: UserBilling - userBilling"))
}

// LegalEntityType is the resolver for the legalEntityType field.
func (r *userBillingResolver) LegalEntityType(ctx context.Context, obj *model.UserBilling) (*model.LegalEntityType, error) {
	panic(fmt.Errorf("not implemented: LegalEntityType - legalEntityType"))
}

// TaxRate is the resolver for the taxRate field.
func (r *userBillingResolver) TaxRate(ctx context.Context, obj *model.UserBilling) (*model.TaxRate, error) {
	panic(fmt.Errorf("not implemented: TaxRate - taxRate"))
}

// User is the resolver for the user field.
func (r *userBillingResolver) User(ctx context.Context, obj *model.UserBilling) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

// UserBilling returns graph.UserBillingResolver implementation.
func (r *Resolver) UserBilling() graph.UserBillingResolver { return &userBillingResolver{r} }

type mutationResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userBillingResolver struct{ *Resolver }

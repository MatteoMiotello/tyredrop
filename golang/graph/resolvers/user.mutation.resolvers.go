package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"errors"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/graphErrors"
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

// UpdateUserStatus is the resolver for the updateUserStatus field.
func (r *mutationResolver) UpdateUserStatus(ctx context.Context, userID int64, confirmed *bool, rejected *bool) (*model.User, error) {
	user, err := r.UserDao.
		Load(models.UserRels.UserRole).
		FindOneById(ctx, userID)

	if err != nil {
		return nil, err
	}

	if user.R.UserRole.Admin {
		return nil, graphErrors.NewGraphError(ctx, errors.New("User is admin"), "USER_IS_ADMIN")
	}

	if confirmed != nil {
		user.Confirmed = *confirmed
	}

	if rejected != nil {
		user.Rejected = *rejected
	}

	if user.Confirmed == true && user.Rejected == true {
		return nil, graphErrors.NewGraphError(ctx, errors.New("User cannot be both confirmed and rejected"), "UNABLE_TO_UPDATE_STATUS")
	}

	err = r.UserDao.Save(ctx, user)

	if err != nil {
		return nil, err
	}

	return converters.UserToGraphQL(user), nil
}
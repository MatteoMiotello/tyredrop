package aggregators

import (
	"context"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/domain/user"
)

type UserAggregator struct {
	*user.AddressDao
}

func NewUserAggregator(addressDao *user.AddressDao) *UserAggregator {
	return &UserAggregator{
		addressDao,
	}
}

func (u UserAggregator) GetAllAddressesByUser(ctx context.Context, userId int64) ([]*model.UserAddress, error) {
	dbModels, _ := u.AddressDao.FindAllByUserId(ctx, userId)

	if dbModels == nil {
		return []*model.UserAddress{}, nil
	}

	var graphModels []*model.UserAddress

	for _, a := range dbModels {
		graphModels = append(graphModels, converters.UserAddressToGraphQL(a))
	}

	return graphModels, nil
}

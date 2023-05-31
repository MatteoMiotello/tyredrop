package cart

import (
	"context"
	"pillowww/titw/models"
)

type service struct {
	Dao *Dao
}

func NewCartService(dao *Dao) *service {
	return &service{
		Dao: dao,
	}
}

func (s service) AddOrUpdateCart(ctx context.Context, user *models.User, productItemId int64, quantity *int) (*models.Cart, error) {
	if quantity == nil {
		q := 1
		quantity = &q
	}

	cart, _ := s.Dao.FindOneByUserAndProductItemId(ctx, user.ID, productItemId)

	if cart == nil {
		cart = &models.Cart{
			UserID:        user.ID,
			ProductItemID: productItemId,
			Quantity:      *quantity,
		}
	} else {
		cart.Quantity = cart.Quantity + *quantity
	}

	err := s.Dao.Insert(ctx, cart)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

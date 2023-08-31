package cart

import (
	"context"
	"github.com/friendsofgo/errors"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/models"
)

type service struct {
	Dao          *Dao
	ItemPriceDao *product.ItemPriceDao
}

func NewCartService(dao *Dao, itemPriceDao *product.ItemPriceDao) *service {
	return &service{
		Dao:          dao,
		ItemPriceDao: itemPriceDao,
	}
}

func (s service) AddCart(ctx context.Context, user *models.User, productItemId int64, quantity *int) (*models.Cart, error) {
	if quantity == nil {
		q := 1
		quantity = &q
	}

	currency, err := auth2.CurrentCurrency(ctx)

	if err != nil {
		return nil, errors.WithMessage(err, "Currency not found in language")
	}

	price, _ := s.ItemPriceDao.
		FindOneByProductItemIdAndCurrencyId(ctx, currency.ID, productItemId)

	if err != nil {
		return nil, errors.WithMessage(err, "Product item is not associate with price with currency: "+currency.IsoCode)
	}

	cart, _ := s.Dao.FindOneByUserAndProductItemPriceId(ctx, user.ID, price.ID)

	if cart == nil {
		cart = &models.Cart{
			UserID:             user.ID,
			ProductItemPriceID: price.ID,
			Quantity:           *quantity,
		}
	} else {
		cart.Quantity = cart.Quantity + *quantity
	}

	err = s.Dao.Save(ctx, cart)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (s service) UpdateCart() {

}

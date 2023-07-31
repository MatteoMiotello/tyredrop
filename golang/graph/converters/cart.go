package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func CartToGraphQL(cart *models.Cart) *model.Cart {
	return &model.Cart{
		ID:                 cart.ID,
		UserID:             cart.UserID,
		ProductItemPriceID: cart.ProductItemPriceID,
		Quantity:           cart.Quantity,
	}
}

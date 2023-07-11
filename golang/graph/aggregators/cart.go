package aggregators

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	auth2 "pillowww/titw/internal/auth"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/cart"
	"pillowww/titw/models"
)

func GetAllCartsByUserId(ctx context.Context, cartDao *cart.Dao, userId int64) (*model.CartResponse, error) {
	currentLang := auth2.CurrentLanguage(ctx)

	defaultCur, err := currency.NewDao(db.DB).
		Load(models.CurrencyRels.CurrencyLanguages, models.CurrencyLanguageWhere.LanguageID.EQ(currentLang.L.ID)).
		FindDefault(ctx)

	if err != nil {
		return nil, err
	}

	cartModels, _ := cartDao.
		Load(
			qm.Rels(
				models.CartRels.ProductItem,
				models.ProductItemRels.ProductItemPrices,
			),
			models.ProductItemPriceWhere.CurrencyID.EQ(defaultCur.ID),
		).
		FindAllByUserId(ctx, userId)
	amountTotal := new(float64)

	if cartModels == nil {
		return &model.CartResponse{
			Items:      []*model.Cart{},
			TotalPrice: nil,
		}, nil
	}

	var graphModels []*model.Cart
	for _, c := range cartModels {
		price := c.R.ProductItem.R.ProductItemPrices[0]

		amount, err := currency.ToFloat(price.Price, defaultCur.IsoCode)

		if err != nil {
			return nil, err
		}

		amount = *amountTotal + (amount * float64(c.Quantity))
		amountTotal = &amount
		graphModels = append(graphModels, converters.CartToGraphQL(c))
	}

	return &model.CartResponse{
		Items: graphModels,
		TotalPrice: &model.TotalPrice{
			Value:    *amountTotal,
			Currency: converters.CurrencyToGraphQL(defaultCur),
		},
	}, nil
}

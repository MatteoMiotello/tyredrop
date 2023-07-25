package product

import (
	"context"
	"github.com/friendsofgo/errors"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
	"strconv"
)

type PriceService struct {
	ProductDao     *Dao
	ProductItemDao *ItemDao
	PriceMarkupDao *PriceMarkupDao
	CurrencyDao    *currency.Dao
}

func NewPriceService(productDao *Dao, itemDao *ItemDao, markupDao *PriceMarkupDao, currencyDao *currency.Dao) *PriceService {
	return &PriceService{
		productDao,
		itemDao,
		markupDao,
		currencyDao,
	}
}

func (p PriceService) findPriceMarkup(ctx context.Context, pi *models.ProductItem) (*models.ProductPriceMarkup, error) {
	product, err := p.ProductDao.FindOneById(ctx, pi.ProductID)

	if err != nil {
		return nil, err
	}

	markup, _ := p.PriceMarkupDao.FindPriceMarkupByProductId(ctx, product)

	if markup != nil {
		return markup, nil
	}

	markup, _ = p.PriceMarkupDao.FindPriceMarkupByBrandId(ctx, product.BrandID)

	if markup != nil {
		return markup, nil
	}

	markup, _ = p.PriceMarkupDao.FindPriceMarkupByProductCategoryId(ctx, product.ProductCategoryID)

	if markup != nil {
		return markup, nil
	}

	markup, _ = p.PriceMarkupDao.FindPriceMarkupDefault(ctx)

	if markup == nil {
		return nil, errors.New("markup default not found")
	}

	return markup, nil
}

func (p PriceService) CalculateAndStoreProductPrices(ctx context.Context, pi *models.ProductItem) error {
	defCur, err := p.CurrencyDao.FindDefault(ctx)

	if defCur == nil {
		return errors.WithMessage(err, "currency not found")
	}

	markup, err := p.findPriceMarkup(ctx, pi)

	if err != nil {
		return errors.WithMessage(err, "Price markup not found")
	}

	price, _ := p.ProductDao.FindPriceForProductItemAndCurrency(ctx, pi, defCur)
	charge := pi.SupplierPrice * markup.MarkupPercentage / 100
	priceValue := pi.SupplierPrice + charge

	if price != nil {
		if price.Price == priceValue {
			return nil
		}

		err := p.ProductDao.Delete(ctx, price)

		if err != nil {
			return errors.WithMessage(err, "Error deleting old price")
		}
	}

	newPrice := &models.ProductItemPrice{
		CurrencyID:    defCur.ID,
		ProductItemID: pi.ID,
		Price:         pi.SupplierPrice + charge,
	}

	err = p.ProductDao.Insert(ctx, newPrice)

	if err != nil {
		return errors.WithMessage(err, "Error inserting new price with value: "+strconv.Itoa(newPrice.Price))
	}

	return nil
}

func (p PriceService) UpdatePricesByCategory(ctx context.Context, category *models.ProductCategory) error {
	pItems, _ := p.ProductItemDao.FindByCategory(ctx, category)
	if pItems == nil {
		return nil
	}

	for _, pi := range pItems {
		err := p.CalculateAndStoreProductPrices(ctx, pi)
		if err != nil {
			return err
		}
	}

	return nil
}

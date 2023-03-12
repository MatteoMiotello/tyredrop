package product

import (
	"context"
	"errors"
	"pillowww/titw/internal/currency"
	"pillowww/titw/models"
)

type PriceService struct {
	ProductDao  *Dao
	CurrencyDao *currency.Dao
}

func NewPriceService(productDao *Dao, currencyDao *currency.Dao) *PriceService {
	return &PriceService{
		productDao,
		currencyDao,
	}
}

func (p PriceService) findPriceMarkup(ctx context.Context, pi *models.ProductItem) (*models.ProductPriceMarkup, error) {
	product, err := p.ProductDao.FindOneById(ctx, pi.ProductID)

	if err != nil {
		return nil, err
	}

	markup, _ := p.ProductDao.FindPriceMarkupByProductId(ctx, product)

	if markup != nil {
		return markup, nil
	}

	markup, _ = p.ProductDao.FindPriceMarkupByBrandId(ctx, product.BrandID)

	if markup != nil {
		return markup, nil
	}

	markup, _ = p.ProductDao.FindPriceMarkupByProductCategoryId(ctx, product.ProductCategoryID)

	if markup != nil {
		return markup, nil
	}

	markup, _ = p.ProductDao.FindPriceMarkupDefault(ctx)

	if markup == nil {
		return nil, errors.New("markup default not found")
	}

	return markup, nil
}

func (p PriceService) CalculateAndStoreProductPrices(ctx context.Context, pi *models.ProductItem) error {
	items, _ := p.ProductDao.FindAllPriceForProductItem(ctx, pi)

	if items != nil {
		for _, item := range items {
			err := p.ProductDao.DeleteProductItemPrice(ctx, item)
			if err != nil {
				return err
			}
		}
	}

	defCur, _ := p.CurrencyDao.FindDefault(ctx)

	if defCur == nil {
		return errors.New("currency not found")
	}

	markup, err := p.findPriceMarkup(ctx, pi)

	if err != nil {
		return err
	}

	charge := pi.SupplierPrice * markup.MarkupPercentage / 100

	price := &models.ProductItemPrice{
		CurrencyID:    defCur.ID,
		ProductItemID: pi.ID,
		Price:         pi.SupplierPrice + charge,
	}

	err = p.ProductDao.Insert(ctx, price)

	if err != nil {
		return err
	}

	return nil
}

func (p PriceService) UpdatePricesByCategory(ctx context.Context, category *models.ProductCategory) error {
	pItems, _ := p.ProductDao.FindProductItemsByCategory(ctx, category)
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

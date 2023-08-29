package product

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/models"
	"strconv"
)

type PriceService struct {
	ProductDao          *Dao
	ProductItemDao      *ItemDao
	PriceMarkupDao      *PriceMarkupDao
	CurrencyDao         *currency.Dao
	ProductItemPriceDao *ItemPriceDao
}

func NewPriceService(productDao *Dao, itemDao *ItemDao, markupDao *PriceMarkupDao, currencyDao *currency.Dao, itemPriceDao *ItemPriceDao) *PriceService {
	return &PriceService{
		productDao,
		itemDao,
		markupDao,
		currencyDao,
		itemPriceDao,
	}
}

func (p PriceService) findPriceMarkup(ctx context.Context, pi *models.ProductItem) (*models.ProductPriceMarkup, error) {
	product, err := p.ProductDao.
		Load(models.ProductRels.Brand).
		FindOneById(ctx, pi.ProductID)

	if err != nil {
		return nil, err
	}

	values, err := p.ProductDao.GetAllSpecificationValues(ctx, product)

	if err != nil {
		return nil, err
	}

	for _, v := range values {
		markup, _ := p.PriceMarkupDao.FindByBrandAndSpecificationValue(ctx, product.R.Brand, v)

		if markup != nil {
			fmt.Println("v+b: ", pi.ID)
			return markup, nil
		}

		markup, _ = p.PriceMarkupDao.FindBySpecificationValue(ctx, v)

		if markup != nil {
			fmt.Println("v: ", pi.ID)
			return markup, nil
		}
	}

	markup, _ := p.PriceMarkupDao.FindPriceMarkupByBrandId(ctx, product.BrandID)

	if markup != nil {
		fmt.Println("b: ", pi.ID)
		return markup, nil
	}

	markup, err = p.PriceMarkupDao.FindPriceMarkupDefault(ctx)

	if markup == nil {
		return nil, errors.WithMessage(err, "markup default not found")
	}

	fmt.Println("d: ", pi.ID, markup.MarkupPercentage)

	return markup, nil
}

func (p PriceService) calcPrices(ctx context.Context, pi *models.ProductItem, markup *models.ProductPriceMarkup) error {
	defCur, err := p.CurrencyDao.FindDefault(ctx)

	if defCur == nil {
		return errors.WithMessage(err, "currency not found")
	}

	price, err := p.ProductDao.
		Load(models.ProductItemPriceRels.ProductItemPriceAdditions).
		FindPriceForProductItemAndCurrency(ctx, pi, defCur)

	if err != nil {
		fmt.Println("Error fetching price", err)
	}

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
	} else {
		fmt.Println("price not found")
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

func (p PriceService) CalculateAndStoreProductPrices(ctx context.Context, pi *models.ProductItem) error {
	markup, err := p.findPriceMarkup(ctx, pi)

	if err != nil {
		return errors.WithMessage(err, "Price markup not found")
	}

	return p.calcPrices(ctx, pi, markup)
}

func (p PriceService) CalculatePriceAdditions(ctx context.Context, pi *models.ProductItem, record pdtos.ProductDto) error {
	prices, _ := p.ProductItemDao.ProductItemPrices(ctx, pi)

	if prices == nil {
		return nil
	}

	for _, price := range prices {
		oldAdditions, _ := p.ProductItemPriceDao.FindPriceAdditionsByProductItemPriceID(ctx, price.ID)

		if oldAdditions != nil {
			for _, oldAdd := range oldAdditions {
				err := p.ProductItemPriceDao.HardDelete(ctx, oldAdd)

				if err != nil {
					return err
				}
			}
		}

		for _, aCode := range record.GetPriceAdditionCodes() {
			aType, _ := p.ProductItemPriceDao.FindPriceAdditionTypeByCurrencyAndCode(ctx, price.CurrencyID, aCode)

			if aType == nil {
				continue
			}

			a := &models.ProductItemPriceAddition{
				ProductItemPriceID:  price.ID,
				PriceAdditionTypeID: aType.ID,
				AdditionValue:       aType.AdditionValue,
			}

			err := p.ProductItemPriceDao.Insert(ctx, a)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p PriceService) UpdateMarkup(ctx context.Context, markup *models.ProductPriceMarkup, markupPercentage int) error {
	markup.MarkupPercentage = markupPercentage
	err := p.ProductItemPriceDao.Save(ctx, markup)

	if err != nil {
		return err
	}

	return nil
}

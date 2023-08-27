package product

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/models"
	"pillowww/titw/pkg/log"
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

func (p PriceService) CalculatePriceAdditions(ctx context.Context, pi *models.ProductItem, record pdtos.ProductDto) error {
	prices, err := p.ProductItemDao.ProductItemPrices(ctx, pi)

	if err != nil {
		return err
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
	var products models.ProductSlice
	var err error

	if markup.MarkupPercentage == markupPercentage {
		return nil
	}

	dao := p.ProductDao.Load(models.ProductRels.ProductItems)

	if !markup.ProductID.IsZero() {
		p, err := dao.FindOneById(ctx, markup.ProductID.Int64)

		if err != nil {
			return err
		}

		products = append(products, p)
	}

	if !markup.ProductCategoryID.IsZero() {
		products, err = dao.FindByCategoryId(ctx, markup.ProductCategoryID.Int64)

		if err != nil {
			return err
		}
	}

	if !markup.BrandID.IsZero() {
		products, err = dao.FindByBrandId(ctx, markup.BrandID.Int64)

		if err != nil {
			return err
		}
	}

	if markup.BrandID.IsZero() && markup.ProductID.IsZero() && markup.ProductCategoryID.IsZero() {
		products, err = dao.FindAll(ctx)

		if err != nil {
			return err
		}
	}

	err = db.WithTx(ctx, func(tx *sql.Tx) error {
		p.ProductItemPriceDao.Db = tx

		markup.MarkupPercentage = markupPercentage
		err := p.ProductItemPriceDao.Save(ctx, markup)

		if err != nil {
			return err
		}

		for _, pr := range products {
			for _, pi := range pr.R.ProductItems {
				pi := pi
				calcErr := p.CalculateAndStoreProductPrices(ctx, pi)

				if calcErr != nil {
					log.Error("Price not updated for item: "+strconv.Itoa(int(pi.ID)), calcErr)
					return calcErr
				}
			}
		}

		return err
	})

	if err != nil {
		return err
	}

	return nil
}

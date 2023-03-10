package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/bojanz/currency"
	"github.com/volatiletech/null/v8"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	currency2 "pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
	"strings"
	"time"
)

type Service struct {
	ProductDao  *Dao
	BrandDao    *brand.Dao
	CurrencyDao *currency2.Dao
}

func NewService(dao *Dao, bDao *brand.Dao, cDao *currency2.Dao) *Service {
	return &Service{
		dao,
		bDao,
		cDao,
	}
}

func (s Service) FindOrCreateProduct(ctx context.Context, dto pdtos.ProductDto) (*models.Product, error) {
	code := cases.Upper(language.Und).String(dto.GetBrandName())
	b, err := s.BrandDao.FindOneByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	category, err := s.findCategory(ctx, dto.GetProductCategoryCode())

	if err != nil {
		return nil, err
	}

	p := &models.Product{
		ProductCode:       null.StringFrom(dto.GetProductCode()),
		BrandID:           b.ID,
		ProductCategoryID: category.ID,
	}

	err = s.ProductDao.Upsert(ctx, p, false, []string{models.ProductColumns.ProductCode})

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s Service) UpdateSpecifications(ctx context.Context, product *models.Product, dto pdtos.ProductDto) error {
	t := time.Now()
	for key, value := range dto.GetSpecifications() {
		if value == "" {
			continue
		}

		keyS := string(key)

		pSpec, _ := s.ProductDao.FindOneProductSpecificationByCode(ctx, keyS)

		if pSpec == nil {
			continue
		}

		pValue := &models.ProductSpecificationValue{
			ProductID:              product.ID,
			ProductSpecificationID: pSpec.ID,
			SpecificationValue:     value,
		}

		err := s.ProductDao.Upsert(ctx, pValue, false, []string{
			models.ProductSpecificationValueColumns.ProductSpecificationID,
			models.ProductSpecificationValueColumns.ProductID,
		})

		if err != nil {
			return err
		}
	}
	fmt.Println(time.Since(t))

	return nil
}

func (s Service) CreateProductItem(ctx context.Context, product *models.Product, supplier *models.Supplier, price string, quantity int) (*models.ProductItem, error) {
	err := s.deleteOldItems(ctx, product, supplier)
	if err != nil {
		return nil, err
	}

	price = strings.Replace(price, ",", ".", 1)
	amount, err := currency.NewAmount(price, "EUR")
	if err != nil {
		return nil, err
	}

	priceInt := int(amount.BigInt().Int64())

	i := &models.ProductItem{
		ProductID:        product.ID,
		SupplierID:       supplier.ID,
		SupplierPrice:    priceInt,
		SupplierQuantity: quantity,
	}

	err = s.ProductDao.Insert(ctx, i)
	if err != nil {
		return nil, err
	}

	return i, err
}

func (s Service) CalculateAndStoreProductPrices(ctx context.Context, pi *models.ProductItem) error {
	items, _ := s.ProductDao.FindAllPriceForProductItem(ctx, pi)

	if items != nil {
		for _, item := range items {
			err := s.ProductDao.DeleteProductItemPrice(ctx, item)
			if err != nil {
				return err
			}
		}
	}

	defCur, _ := s.CurrencyDao.FindDefault(ctx)

	if defCur == nil {
		return errors.New("currency not found")
	}

	markup, err := s.findPriceMarkup(ctx, pi)

	if err != nil {
		return err
	}

	charge := pi.SupplierPrice * markup.MarkupPercentage / 100

	price := &models.ProductItemPrice{
		CurrencyID:    defCur.ID,
		ProductItemID: pi.ID,
		Price:         pi.SupplierPrice + charge,
	}

	err = s.ProductDao.Insert(ctx, price)

	if err != nil {
		return err
	}

	return nil
}

func (s Service) findPriceMarkup(ctx context.Context, pi *models.ProductItem) (*models.ProductPriceMarkup, error) {
	product, err := s.ProductDao.FindOneById(ctx, pi.ProductID)

	if err != nil {
		return nil, err
	}

	markup, _ := s.ProductDao.FindPriceMarkupByProductId(ctx, product)

	if markup != nil {
		return markup, nil
	}

	markup, _ = s.ProductDao.FindPriceMarkupByBrandId(ctx, product.BrandID)

	if markup != nil {
		return markup, nil
	}

	markup, _ = s.ProductDao.FindPriceMarkupByProductCategoryId(ctx, product.ProductCategoryID)

	if markup != nil {
		return markup, nil
	}

	markup, _ = s.ProductDao.FindPriceMarkupDefault(ctx)

	if markup == nil {
		return nil, errors.New("markup default not found")
	}

	return markup, nil
}

func (s Service) findCategory(ctx context.Context, code constants.ProductCategoryType) (*models.ProductCategory, error) {
	return s.ProductDao.FindCategoryByCode(ctx, string(code))
}

func (s Service) deleteOldItems(ctx context.Context, product *models.Product, supplier *models.Supplier) error {
	items, _ := s.ProductDao.FindProductItems(ctx, product, supplier)

	if items == nil {
		return nil
	}

	for _, item := range items {
		err := s.ProductDao.DeleteProductItem(ctx, item)
		if err != nil {
			return err
		}
	}

	return nil
}

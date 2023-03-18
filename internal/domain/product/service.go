package product

import (
	"context"
	"fmt"
	"github.com/bojanz/currency"
	"github.com/volatiletech/null/v8"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
	"strings"
)

type Service struct {
	ProductDao       *Dao
	ItemDao          *ItemDao
	SpecificationDao *SpecificationDao
	CategoryDao      *CategoryDao
	BrandDao         *brand.Dao
}

func NewService(dao *Dao, bDao *brand.Dao, cDao *CategoryDao, iDao *ItemDao, sDao *SpecificationDao) *Service {
	return &Service{
		dao,
		iDao,
		sDao,
		cDao,
		bDao,
	}
}

func (s Service) findCategory(ctx context.Context, code constants.ProductCategoryType) (*models.ProductCategory, error) {
	return s.CategoryDao.FindByCode(ctx, string(code))
}

func (s Service) deleteOldItems(ctx context.Context, product *models.Product, supplier *models.Supplier) error {
	items, err := s.ItemDao.FindByProductAndSupplier(ctx, product, supplier)

	fmt.Println(err)

	if items == nil {
		return nil
	}

	for _, item := range items {
		err := s.ProductDao.Delete(ctx, item)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) FindOrCreateProduct(ctx context.Context, dto pdtos.ProductDto) (*models.Product, error) {
	p, _ := s.ProductDao.FindOneByCode(ctx, dto.GetProductCode())

	if p != nil {
		return p, nil
	}

	code := cases.Upper(language.Und).String(dto.GetBrandName())
	b, err := s.BrandDao.FindOneByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	category, err := s.findCategory(ctx, dto.GetProductCategoryCode())

	if err != nil {
		return nil, err
	}

	p = &models.Product{
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
	if product.Completed {
		return nil
	}

	specs, err := s.SpecificationDao.FindByProduct(ctx, product)

	if err != nil {
		return err
	}

	dtoSpecs := dto.GetSpecifications()
	specInserted := 0
	for _, spec := range specs {
		value, found := dtoSpecs[constants.ProductSpecification(spec.SpecificationCode)]
		if !found {
			continue
		}

		if value == "" {
			continue
		}

		pValue := &models.ProductSpecificationValue{
			ProductID:              product.ID,
			ProductSpecificationID: spec.ID,
			SpecificationValue:     value,
		}

		err := s.ProductDao.Upsert(ctx, pValue, false, []string{
			models.ProductSpecificationValueColumns.ProductSpecificationID,
			models.ProductSpecificationValueColumns.ProductID,
		})

		if err != nil {
			return err
		}

		specInserted++
	}

	mandatories, err := s.SpecificationDao.FindMandatoryByProduct(ctx, product)

	if len(mandatories) <= specInserted {
		err = s.setProductComplete(ctx, product)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) setProductComplete(ctx context.Context, product *models.Product) error {
	product.Completed = true
	return s.ProductDao.Update(ctx, product)
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

	priceInt, err := amount.Int64()

	if err != nil {
		return nil, err
	}

	i := &models.ProductItem{
		ProductID:        product.ID,
		SupplierID:       supplier.ID,
		SupplierPrice:    int(priceInt),
		SupplierQuantity: quantity,
	}

	err = s.ProductDao.Insert(ctx, i)
	if err != nil {
		return nil, err
	}

	return i, err
}

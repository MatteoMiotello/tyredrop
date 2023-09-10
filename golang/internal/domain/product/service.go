package product

import (
	"context"
	"github.com/bojanz/currency"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/product/pdtos"
	"pillowww/titw/internal/domain/vehicle"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/log"
	"strconv"
	"strings"
)

type Service struct {
	ProductDao            *Dao
	ItemDao               *ItemDao
	SpecificationDao      *SpecificationDao
	SpecificationValueDao *SpecificationValueDao
	CategoryDao           *CategoryDao
	VehicleDao            *vehicle.Dao
	BrandDao              *brand.Dao
}

func NewService(dao *Dao, bDao *brand.Dao, cDao *CategoryDao, iDao *ItemDao, sDao *SpecificationDao, sVDao *SpecificationValueDao, vDao *vehicle.Dao) *Service {
	return &Service{
		dao,
		iDao,
		sDao,
		sVDao,
		cDao,
		vDao,
		bDao,
	}
}

func (s Service) findCategory(ctx context.Context, code constants.ProductCategoryType) (*models.ProductCategory, error) {
	return s.CategoryDao.FindByCode(ctx, string(code))
}

func (s Service) findOrCreateSpecificationValue(ctx context.Context, specification *models.ProductSpecification, value string) (*models.ProductSpecificationValue, error) {
	specificationValue, err := s.SpecificationValueDao.FindBySpecificationAndValue(ctx, specification, value)

	if specificationValue != nil {
		return specificationValue, nil
	}

	specificationValue = &models.ProductSpecificationValue{
		ProductSpecificationID: specification.ID,
		SpecificationValue:     value,
	}

	err = s.SpecificationValueDao.Insert(ctx, specificationValue)

	if err != nil {
		return nil, errors.WithMessage(err, "Unable to insert specification value with value: "+value)
	}

	return specificationValue, nil
}

func (s Service) FindOrCreateProduct(ctx context.Context, dto pdtos.ProductDto, forceUpdate bool) (*models.Product, error) {
	p, _ := s.ProductDao.FindOneByCode(ctx, dto.GetProductCode())

	if p != nil && !forceUpdate {
		return p, nil
	}

	code := cases.Upper(language.Und).String(dto.GetBrandName())
	b, err := s.BrandDao.FindOneByCode(ctx, code)
	if err != nil {
		return nil, errors.WithMessage(err, "Brand not found with code: "+dto.GetBrandName())
	}

	category, err := s.findCategory(ctx, dto.GetProductCategoryCode())

	if err != nil {
		return nil, errors.WithMessage(err, "Category not found with code: "+string(dto.GetProductCategoryCode()))
	}

	vehicleType, err := s.VehicleDao.FindByCode(ctx, dto.GetVehicleType())

	if err != nil {
		return nil, errors.WithMessage(err, "Vehicle type not found for code: "+string(dto.GetVehicleType()))
	}

	name := strings.ToValidUTF8(dto.GetProductName(), "")

	if p == nil {
		p = &models.Product{
			ProductCode: null.StringFrom(dto.GetProductCode()),
		}
	}

	p.BrandID = b.ID
	p.ProductCategoryID = category.ID
	p.VehicleTypeID = vehicleType.ID
	p.Name = null.StringFrom(name)
	p.EprelProductCode = null.StringFromPtr(dto.GetEprelProductCode())
	p.ImageURL = null.StringFromPtr(dto.GetProductImageUrl())

	err = s.ProductDao.Save(ctx, p)

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
		return errors.WithMessage(err, "No specifications found for product: "+strconv.Itoa(int(product.ID)))
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

		pValue, _ := s.SpecificationValueDao.FindByProductAndCode(ctx, product, spec.SpecificationCode)
		specInserted++

		if pValue != nil {
			continue
		}

		pValue, err = s.findOrCreateSpecificationValue(ctx, spec, strings.TrimSpace(strings.ToValidUTF8(value, "")))

		if err != nil {
			return errors.WithMessage(err, "Error in specification value: "+spec.SpecificationCode)
		}

		relation := &models.ProductProductSpecificationValue{
			ProductSpecificationValueID: pValue.ID,
			ProductID:                   product.ID,
		}

		err = s.SpecificationValueDao.Insert(ctx, relation)

		if err != nil {
			log.Error("Error inserting ProductProductSpecificationValue", err)
			return errors.WithMessage(err, "Error inserting ProductProductSpecificationValue "+spec.SpecificationCode+" "+pValue.SpecificationValue)
		}
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

func (s Service) CreateOrUpdateProductItem(ctx context.Context, product *models.Product, supplier *models.Supplier, price string, quantity int) (*models.ProductItem, error) {
	if quantity == 0 {
		return nil, nil
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

	i, _ := s.ItemDao.FindByProductAndSupplier(ctx, product, supplier)

	if i == nil {
		i = &models.ProductItem{
			ProductID:  product.ID,
			SupplierID: supplier.ID,
		}
	}

	i.SupplierQuantity = quantity
	i.SupplierPrice = int(priceInt)

	err = s.ItemDao.Save(ctx, i)

	if err != nil {
		return nil, errors.WithMessage(err, "Error saving item")
	}

	return i, err
}

package converter

import (
	"context"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/auth"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/models"
	"pillowww/titw/pkg/constants"
	"strconv"
)

type ProductConverter struct {
	ProductDao  *product.Dao
	BrandDao    *brand.Dao
	CurrencyDao *currency.Dao
}

func getUnionType(val string, types constants.ProductSpecification) model.UnionValues {
	switch types {
	case constants.SPEC_TYPE_STRING:
		return model.ProductSpecificationValueString{
			Value: val,
		}
	case constants.SPEC_TYPE_INT:
		intVal, _ := strconv.Atoi(val)

		return model.ProductSpecificationValueInt{
			Value: intVal,
		}
	case constants.SPEC_TYPE_FLOAT:
		floatVal, _ := strconv.ParseFloat(val, 32)

		return model.ProductSpecificationValueFloat{
			Value: floatVal,
		}
	}

	return nil
}

func (p ProductConverter) ProductSpecificationValue(ctx context.Context, value *models.ProductSpecificationValue) (*model.ProductSpecificationValue, error) {
	spec, err := p.ProductDao.ProductSpecification(ctx, value)

	if err != nil {
		return nil, err
	}

	auth := auth.FromCtx(ctx)
	lang := auth.GetLanguage(ctx)

	sLang, err := p.ProductDao.ProductSpecificationLanguage(ctx, spec, lang)

	if err != nil {
		return nil, err
	}

	return &model.ProductSpecificationValue{
		ID:    strconv.Itoa(int(value.ID)),
		Code:  spec.SpecificationCode,
		Name:  sLang.Name,
		Value: getUnionType(value.SpecificationValue, constants.ProductSpecification(spec.Type)),
	}, nil
}

func (p ProductConverter) Product(ctx context.Context, product *models.Product) (*model.Product, error) {
	valueModels, err := p.ProductDao.ProductSpecificationValues(ctx, product)

	if err != nil {
		return nil, err
	}

	var values []*model.ProductSpecificationValue

	for _, valueModel := range valueModels {
		graphModel, err := p.ProductSpecificationValue(ctx, valueModel)

		if err != nil {
			return nil, err
		}

		values = append(values, graphModel)
	}

	brand, err := p.ProductDao.Brand(ctx, product)

	if err != nil {
		return nil, err
	}

	brandConverter := BrandConverter{BrandDao: p.BrandDao}

	return &model.Product{
		ID:             strconv.Itoa(int(product.ID)),
		Code:           product.ProductCode.String,
		Brand:          brandConverter.Brand(ctx, brand),
		Specifications: values,
	}, nil
}

func (p ProductConverter) ProductItem(ctx context.Context, productItem *models.ProductItem) (*model.ProductItem, error) {
	product, err := p.ProductDao.Product(ctx, productItem)

	if err != nil {
		return nil, err
	}

	graphProduct, err := p.Product(ctx, product)

	if err != nil {
		return nil, err
	}

	var prices []*model.Price
	priceModels, err := p.ProductDao.ProductItemPrices(ctx, productItem)

	if err != nil {
		return nil, err
	}

	priceConverter := PriceConverter{
		ProductDao:  p.ProductDao,
		CurrencyDao: p.CurrencyDao,
	}

	for _, pPrice := range priceModels {
		graphPrice, err := priceConverter.ProductItemPrice(ctx, pPrice)

		if err != nil {
			return nil, err
		}

		prices = append(prices, graphPrice)
	}

	return &model.ProductItem{
		ID: productItem.ID,
	}, nil
}

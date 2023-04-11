package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func BrandToGraphQL(brandModel *models.Brand) *model.Brand {
	return &model.Brand{
		ID:        brandModel.ID,
		Name:      brandModel.Name,
		ImageLogo: brandModel.ImageLogo.String,
	}
}

package converter

import (
	"context"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/domain/brand"
	"pillowww/titw/models"
	"strconv"
)

type BrandConverter struct {
	BrandDao *brand.Dao
}

func (c BrandConverter) Brand(ctx context.Context, brandModel *models.Brand) *model.Brand {
	return &model.Brand{
		ID:        strconv.Itoa(int(brandModel.ID)),
		Name:      brandModel.Name,
		ImageLogo: brandModel.ImageLogo.String,
	}
}

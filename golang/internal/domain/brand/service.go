package brand

import (
	"context"
	"github.com/volatiletech/null/v8"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"pillowww/titw/models"
)

type Service struct {
	Dao *Dao
}

func NewBrandService(dao *Dao) *Service {
	return &Service{
		Dao: dao,
	}
}

func (s Service) FindOrCreateBrand(ctx context.Context, name string) (*models.Brand, error) {
	code := cases.Upper(language.Und).String(name)
	b, _ := s.Dao.FindOneByCode(ctx, code)

	if b == nil {
		b = &models.Brand{
			Name:      name,
			BrandCode: code,
			Quality:   null.IntFrom(3),
		}

		err := s.Dao.Insert(ctx, b)

		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

package language

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
)

type Language struct {
	L *models.Language
}

var fallbackLanguage *Language

func SetFallbackLanguage(isoCode string) error {
	language, err := models.Languages(qm.Where("iso_code = ?", isoCode)).One(context.Background(), db.DB)

	if err != nil {
		return err
	}

	fLanguage := Language{
		L: language,
	}

	fallbackLanguage = &fLanguage

	return nil
}

func FallbackLanguage() *Language {
	return fallbackLanguage
}

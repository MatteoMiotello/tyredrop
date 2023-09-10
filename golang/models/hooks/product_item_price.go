package hooks

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"pillowww/titw/models"
)

func RegisterHooks() {
	models.AddProductItemPriceHook(boil.AfterDeleteHook, deleteCarts)
}

func deleteCarts(ctx context.Context, exec boil.ContextExecutor, p *models.ProductItemPrice) error {
	carts, err := p.Carts().All(ctx, exec)
	if err != nil {
		return err
	}

	_, err = carts.DeleteAll(ctx, exec, false)
	if err != nil {
		return err
	}

	return nil
}

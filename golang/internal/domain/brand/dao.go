package brand

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/db"
	"pillowww/titw/models"
	"time"
)

type Dao struct {
	*db.Dao
}

func NewDao(executor boil.ContextExecutor) *Dao {
	return &Dao{
		db.DaoFromExecutor(executor),
	}
}

func (d Dao) SetDao(dao *db.Dao) db.DaoMod {
	d.Dao = dao
	return d
}

func (d Dao) GetDao() *db.Dao {
	return d.Dao
}

func (d Dao) Load(relationship string, mods ...qm.QueryMod) *Dao {
	return db.Load(d, relationship, mods...)
}

func (d Dao) Paginate(first int, offset int) *Dao {
	return db.Paginate(d, first, offset)
}

func (d Dao) Clone() db.DaoMod {
	return Dao{
		d.Dao.Clone(),
	}
}

func (d *Dao) FindOneById(ctx context.Context, id int64) (*models.Brand, error) {
	return models.Brands(
		d.GetMods(
			models.BrandWhere.ID.EQ(id),
		)...,
	).One(ctx, d.Db)
}

func (d *Dao) FindOneByCode(ctx context.Context, code string) (*models.Brand, error) {
	return models.Brands(
		d.GetMods(
			models.BrandWhere.BrandCode.EQ(code),
		)...,
	).One(ctx, d.Db)
}

func (d *Dao) FindByName(ctx context.Context, name string) (models.BrandSlice, error) {
	return models.Brands(
		d.GetMods(
			qm.Where(models.BrandColumns.Name+" ILIKE ?", fmt.Sprintf("%%%s%%", name)),
		)...,
	).All(ctx, d.Db)
}

func (d *Dao) FindAll(ctx context.Context) (models.BrandSlice, error) {
	return models.Brands(
		d.GetMods(
			qm.OrderBy(models.BrandColumns.Quality+" DESC"),
		)...,
	).All(ctx, d.Db)
}

func (d *Dao) BestBrand(ctx context.Context, from time.Time, to time.Time) (*models.Brand, error) {
	return models.Brands(
		d.GetMods(
			qm.LeftOuterJoin("products on products.brand_id = brands.id"),
			qm.LeftOuterJoin("product_items on product_items.product_id = products.id"),
			qm.LeftOuterJoin("product_item_prices on product_item_prices.product_item_id = product_items.id"),
			qm.LeftOuterJoin("order_rows on order_rows.product_item_price_id = product_item_prices.id"),
			qm.LeftOuterJoin("orders on order_rows.order_id = orders.id"),
			models.OrderWhere.CreatedAt.GTE(from),
			models.OrderWhere.CreatedAt.LTE(to),
			models.OrderWhere.Status.IN(model.OrderProcessedStatusCollection),
			qm.GroupBy("brands.id"),
			qm.OrderBy("count( order_rows.id ) DESC"),
		)...,
	).One(ctx, d.Db)
}

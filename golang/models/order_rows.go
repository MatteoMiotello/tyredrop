// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OrderRow is an object representing the database table.
type OrderRow struct {
	ID                 int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrderID            int64       `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	Amount             int         `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Quantity           int         `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	TrackingNumber     null.String `boil:"tracking_number" json:"tracking_number,omitempty" toml:"tracking_number" yaml:"tracking_number,omitempty"`
	DeliveredAt        null.Time   `boil:"delivered_at" json:"delivered_at,omitempty" toml:"delivered_at" yaml:"delivered_at,omitempty"`
	UpdatedAt          time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt          time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ProductItemPriceID int64       `boil:"product_item_price_id" json:"product_item_price_id" toml:"product_item_price_id" yaml:"product_item_price_id"`
	AdditionsAmount    int         `boil:"additions_amount" json:"additions_amount" toml:"additions_amount" yaml:"additions_amount"`

	R *orderRowR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderRowL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderRowColumns = struct {
	ID                 string
	OrderID            string
	Amount             string
	Quantity           string
	TrackingNumber     string
	DeliveredAt        string
	UpdatedAt          string
	CreatedAt          string
	ProductItemPriceID string
	AdditionsAmount    string
}{
	ID:                 "id",
	OrderID:            "order_id",
	Amount:             "amount",
	Quantity:           "quantity",
	TrackingNumber:     "tracking_number",
	DeliveredAt:        "delivered_at",
	UpdatedAt:          "updated_at",
	CreatedAt:          "created_at",
	ProductItemPriceID: "product_item_price_id",
	AdditionsAmount:    "additions_amount",
}

var OrderRowTableColumns = struct {
	ID                 string
	OrderID            string
	Amount             string
	Quantity           string
	TrackingNumber     string
	DeliveredAt        string
	UpdatedAt          string
	CreatedAt          string
	ProductItemPriceID string
	AdditionsAmount    string
}{
	ID:                 "order_rows.id",
	OrderID:            "order_rows.order_id",
	Amount:             "order_rows.amount",
	Quantity:           "order_rows.quantity",
	TrackingNumber:     "order_rows.tracking_number",
	DeliveredAt:        "order_rows.delivered_at",
	UpdatedAt:          "order_rows.updated_at",
	CreatedAt:          "order_rows.created_at",
	ProductItemPriceID: "order_rows.product_item_price_id",
	AdditionsAmount:    "order_rows.additions_amount",
}

// Generated where

var OrderRowWhere = struct {
	ID                 whereHelperint64
	OrderID            whereHelperint64
	Amount             whereHelperint
	Quantity           whereHelperint
	TrackingNumber     whereHelpernull_String
	DeliveredAt        whereHelpernull_Time
	UpdatedAt          whereHelpertime_Time
	CreatedAt          whereHelpertime_Time
	ProductItemPriceID whereHelperint64
	AdditionsAmount    whereHelperint
}{
	ID:                 whereHelperint64{field: "\"order_rows\".\"id\""},
	OrderID:            whereHelperint64{field: "\"order_rows\".\"order_id\""},
	Amount:             whereHelperint{field: "\"order_rows\".\"amount\""},
	Quantity:           whereHelperint{field: "\"order_rows\".\"quantity\""},
	TrackingNumber:     whereHelpernull_String{field: "\"order_rows\".\"tracking_number\""},
	DeliveredAt:        whereHelpernull_Time{field: "\"order_rows\".\"delivered_at\""},
	UpdatedAt:          whereHelpertime_Time{field: "\"order_rows\".\"updated_at\""},
	CreatedAt:          whereHelpertime_Time{field: "\"order_rows\".\"created_at\""},
	ProductItemPriceID: whereHelperint64{field: "\"order_rows\".\"product_item_price_id\""},
	AdditionsAmount:    whereHelperint{field: "\"order_rows\".\"additions_amount\""},
}

// OrderRowRels is where relationship names are stored.
var OrderRowRels = struct {
	Order            string
	ProductItemPrice string
}{
	Order:            "Order",
	ProductItemPrice: "ProductItemPrice",
}

// orderRowR is where relationships are stored.
type orderRowR struct {
	Order            *Order            `boil:"Order" json:"Order" toml:"Order" yaml:"Order"`
	ProductItemPrice *ProductItemPrice `boil:"ProductItemPrice" json:"ProductItemPrice" toml:"ProductItemPrice" yaml:"ProductItemPrice"`
}

// NewStruct creates a new relationship struct
func (*orderRowR) NewStruct() *orderRowR {
	return &orderRowR{}
}

func (r *orderRowR) GetOrder() *Order {
	if r == nil {
		return nil
	}
	return r.Order
}

func (r *orderRowR) GetProductItemPrice() *ProductItemPrice {
	if r == nil {
		return nil
	}
	return r.ProductItemPrice
}

// orderRowL is where Load methods for each relationship are stored.
type orderRowL struct{}

var (
	orderRowAllColumns            = []string{"id", "order_id", "amount", "quantity", "tracking_number", "delivered_at", "updated_at", "created_at", "product_item_price_id", "additions_amount"}
	orderRowColumnsWithoutDefault = []string{"order_id", "amount", "quantity", "product_item_price_id"}
	orderRowColumnsWithDefault    = []string{"id", "tracking_number", "delivered_at", "updated_at", "created_at", "additions_amount"}
	orderRowPrimaryKeyColumns     = []string{"id"}
	orderRowGeneratedColumns      = []string{}
)

type (
	// OrderRowSlice is an alias for a slice of pointers to OrderRow.
	// This should almost always be used instead of []OrderRow.
	OrderRowSlice []*OrderRow
	// OrderRowHook is the signature for custom OrderRow hook methods
	OrderRowHook func(context.Context, boil.ContextExecutor, *OrderRow) error

	orderRowQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderRowType                 = reflect.TypeOf(&OrderRow{})
	orderRowMapping              = queries.MakeStructMapping(orderRowType)
	orderRowPrimaryKeyMapping, _ = queries.BindMapping(orderRowType, orderRowMapping, orderRowPrimaryKeyColumns)
	orderRowInsertCacheMut       sync.RWMutex
	orderRowInsertCache          = make(map[string]insertCache)
	orderRowUpdateCacheMut       sync.RWMutex
	orderRowUpdateCache          = make(map[string]updateCache)
	orderRowUpsertCacheMut       sync.RWMutex
	orderRowUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderRowAfterSelectHooks []OrderRowHook

var orderRowBeforeInsertHooks []OrderRowHook
var orderRowAfterInsertHooks []OrderRowHook

var orderRowBeforeUpdateHooks []OrderRowHook
var orderRowAfterUpdateHooks []OrderRowHook

var orderRowBeforeDeleteHooks []OrderRowHook
var orderRowAfterDeleteHooks []OrderRowHook

var orderRowBeforeUpsertHooks []OrderRowHook
var orderRowAfterUpsertHooks []OrderRowHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrderRow) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrderRow) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrderRow) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrderRow) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrderRow) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrderRow) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrderRow) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrderRow) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrderRow) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderRowAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderRowHook registers your hook function for all future operations.
func AddOrderRowHook(hookPoint boil.HookPoint, orderRowHook OrderRowHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderRowAfterSelectHooks = append(orderRowAfterSelectHooks, orderRowHook)
	case boil.BeforeInsertHook:
		orderRowBeforeInsertHooks = append(orderRowBeforeInsertHooks, orderRowHook)
	case boil.AfterInsertHook:
		orderRowAfterInsertHooks = append(orderRowAfterInsertHooks, orderRowHook)
	case boil.BeforeUpdateHook:
		orderRowBeforeUpdateHooks = append(orderRowBeforeUpdateHooks, orderRowHook)
	case boil.AfterUpdateHook:
		orderRowAfterUpdateHooks = append(orderRowAfterUpdateHooks, orderRowHook)
	case boil.BeforeDeleteHook:
		orderRowBeforeDeleteHooks = append(orderRowBeforeDeleteHooks, orderRowHook)
	case boil.AfterDeleteHook:
		orderRowAfterDeleteHooks = append(orderRowAfterDeleteHooks, orderRowHook)
	case boil.BeforeUpsertHook:
		orderRowBeforeUpsertHooks = append(orderRowBeforeUpsertHooks, orderRowHook)
	case boil.AfterUpsertHook:
		orderRowAfterUpsertHooks = append(orderRowAfterUpsertHooks, orderRowHook)
	}
}

// One returns a single orderRow record from the query.
func (q orderRowQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderRow, error) {
	o := &OrderRow{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for order_rows")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrderRow records from the query.
func (q orderRowQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderRowSlice, error) {
	var o []*OrderRow

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrderRow slice")
	}

	if len(orderRowAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrderRow records in the query.
func (q orderRowQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count order_rows rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orderRowQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if order_rows exists")
	}

	return count > 0, nil
}

// Order pointed to by the foreign key.
func (o *OrderRow) Order(mods ...qm.QueryMod) orderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrderID),
	}

	queryMods = append(queryMods, mods...)

	return Orders(queryMods...)
}

// ProductItemPrice pointed to by the foreign key.
func (o *OrderRow) ProductItemPrice(mods ...qm.QueryMod) productItemPriceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProductItemPriceID),
	}

	queryMods = append(queryMods, mods...)

	return ProductItemPrices(queryMods...)
}

// LoadOrder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (orderRowL) LoadOrder(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrderRow interface{}, mods queries.Applicator) error {
	var slice []*OrderRow
	var object *OrderRow

	if singular {
		var ok bool
		object, ok = maybeOrderRow.(*OrderRow)
		if !ok {
			object = new(OrderRow)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOrderRow)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOrderRow))
			}
		}
	} else {
		s, ok := maybeOrderRow.(*[]*OrderRow)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOrderRow)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOrderRow))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orderRowR{}
		}
		args = append(args, object.OrderID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orderRowR{}
			}

			for _, a := range args {
				if a == obj.OrderID {
					continue Outer
				}
			}

			args = append(args, obj.OrderID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`orders`),
		qm.WhereIn(`orders.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Order")
	}

	var resultSlice []*Order
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Order")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for orders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orders")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Order = foreign
		if foreign.R == nil {
			foreign.R = &orderR{}
		}
		foreign.R.OrderRows = append(foreign.R.OrderRows, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrderID == foreign.ID {
				local.R.Order = foreign
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.OrderRows = append(foreign.R.OrderRows, local)
				break
			}
		}
	}

	return nil
}

// LoadProductItemPrice allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (orderRowL) LoadProductItemPrice(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrderRow interface{}, mods queries.Applicator) error {
	var slice []*OrderRow
	var object *OrderRow

	if singular {
		var ok bool
		object, ok = maybeOrderRow.(*OrderRow)
		if !ok {
			object = new(OrderRow)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOrderRow)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOrderRow))
			}
		}
	} else {
		s, ok := maybeOrderRow.(*[]*OrderRow)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOrderRow)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOrderRow))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orderRowR{}
		}
		args = append(args, object.ProductItemPriceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orderRowR{}
			}

			for _, a := range args {
				if a == obj.ProductItemPriceID {
					continue Outer
				}
			}

			args = append(args, obj.ProductItemPriceID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`product_item_prices`),
		qm.WhereIn(`product_item_prices.id in ?`, args...),
		qmhelper.WhereIsNull(`product_item_prices.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ProductItemPrice")
	}

	var resultSlice []*ProductItemPrice
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ProductItemPrice")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for product_item_prices")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for product_item_prices")
	}

	if len(productItemPriceAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ProductItemPrice = foreign
		if foreign.R == nil {
			foreign.R = &productItemPriceR{}
		}
		foreign.R.OrderRows = append(foreign.R.OrderRows, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductItemPriceID == foreign.ID {
				local.R.ProductItemPrice = foreign
				if foreign.R == nil {
					foreign.R = &productItemPriceR{}
				}
				foreign.R.OrderRows = append(foreign.R.OrderRows, local)
				break
			}
		}
	}

	return nil
}

// SetOrder of the orderRow to the related item.
// Sets o.R.Order to related.
// Adds o to related.R.OrderRows.
func (o *OrderRow) SetOrder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Order) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"order_rows\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"order_id"}),
		strmangle.WhereClause("\"", "\"", 2, orderRowPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrderID = related.ID
	if o.R == nil {
		o.R = &orderRowR{
			Order: related,
		}
	} else {
		o.R.Order = related
	}

	if related.R == nil {
		related.R = &orderR{
			OrderRows: OrderRowSlice{o},
		}
	} else {
		related.R.OrderRows = append(related.R.OrderRows, o)
	}

	return nil
}

// SetProductItemPrice of the orderRow to the related item.
// Sets o.R.ProductItemPrice to related.
// Adds o to related.R.OrderRows.
func (o *OrderRow) SetProductItemPrice(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ProductItemPrice) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"order_rows\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"product_item_price_id"}),
		strmangle.WhereClause("\"", "\"", 2, orderRowPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProductItemPriceID = related.ID
	if o.R == nil {
		o.R = &orderRowR{
			ProductItemPrice: related,
		}
	} else {
		o.R.ProductItemPrice = related
	}

	if related.R == nil {
		related.R = &productItemPriceR{
			OrderRows: OrderRowSlice{o},
		}
	} else {
		related.R.OrderRows = append(related.R.OrderRows, o)
	}

	return nil
}

// OrderRows retrieves all the records using an executor.
func OrderRows(mods ...qm.QueryMod) orderRowQuery {
	mods = append(mods, qm.From("\"order_rows\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"order_rows\".*"})
	}

	return orderRowQuery{q}
}

// FindOrderRow retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrderRow(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*OrderRow, error) {
	orderRowObj := &OrderRow{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"order_rows\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderRowObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from order_rows")
	}

	if err = orderRowObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderRowObj, err
	}

	return orderRowObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderRow) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no order_rows provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderRowColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderRowInsertCacheMut.RLock()
	cache, cached := orderRowInsertCache[key]
	orderRowInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderRowAllColumns,
			orderRowColumnsWithDefault,
			orderRowColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderRowType, orderRowMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderRowType, orderRowMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"order_rows\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"order_rows\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into order_rows")
	}

	if !cached {
		orderRowInsertCacheMut.Lock()
		orderRowInsertCache[key] = cache
		orderRowInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OrderRow.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrderRow) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderRowUpdateCacheMut.RLock()
	cache, cached := orderRowUpdateCache[key]
	orderRowUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderRowAllColumns,
			orderRowPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update order_rows, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"order_rows\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orderRowPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderRowType, orderRowMapping, append(wl, orderRowPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update order_rows row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for order_rows")
	}

	if !cached {
		orderRowUpdateCacheMut.Lock()
		orderRowUpdateCache[key] = cache
		orderRowUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q orderRowQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for order_rows")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for order_rows")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderRowSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderRowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"order_rows\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderRowPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in orderRow slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all orderRow")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrderRow) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no order_rows provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderRowColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	orderRowUpsertCacheMut.RLock()
	cache, cached := orderRowUpsertCache[key]
	orderRowUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orderRowAllColumns,
			orderRowColumnsWithDefault,
			orderRowColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			orderRowAllColumns,
			orderRowPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert order_rows, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orderRowPrimaryKeyColumns))
			copy(conflict, orderRowPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"order_rows\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orderRowType, orderRowMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderRowType, orderRowMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert order_rows")
	}

	if !cached {
		orderRowUpsertCacheMut.Lock()
		orderRowUpsertCache[key] = cache
		orderRowUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OrderRow record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrderRow) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrderRow provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderRowPrimaryKeyMapping)
	sql := "DELETE FROM \"order_rows\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from order_rows")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for order_rows")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orderRowQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no orderRowQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from order_rows")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for order_rows")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderRowSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderRowBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderRowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"order_rows\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderRowPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orderRow slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for order_rows")
	}

	if len(orderRowAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrderRow) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrderRow(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderRowSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderRowSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderRowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"order_rows\".* FROM \"order_rows\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderRowPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrderRowSlice")
	}

	*o = slice

	return nil
}

// OrderRowExists checks if the OrderRow row exists.
func OrderRowExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"order_rows\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if order_rows exists")
	}

	return exists, nil
}

// Exists checks if the OrderRow row exists.
func (o *OrderRow) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrderRowExists(ctx, exec, o.ID)
}

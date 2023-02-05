// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TaxRate is an object representing the database table.
type TaxRate struct {
	ID               int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	MarkupPercentage int       `boil:"markup_percentage" json:"markup_percentage" toml:"markup_percentage" yaml:"markup_percentage"`
	Name             string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *taxRateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taxRateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaxRateColumns = struct {
	ID               string
	MarkupPercentage string
	Name             string
	CreatedAt        string
}{
	ID:               "id",
	MarkupPercentage: "markup_percentage",
	Name:             "name",
	CreatedAt:        "created_at",
}

var TaxRateTableColumns = struct {
	ID               string
	MarkupPercentage string
	Name             string
	CreatedAt        string
}{
	ID:               "tax_rates.id",
	MarkupPercentage: "tax_rates.markup_percentage",
	Name:             "tax_rates.name",
	CreatedAt:        "tax_rates.created_at",
}

// Generated where

var TaxRateWhere = struct {
	ID               whereHelperint64
	MarkupPercentage whereHelperint
	Name             whereHelperstring
	CreatedAt        whereHelpertime_Time
}{
	ID:               whereHelperint64{field: "\"tax_rates\".\"id\""},
	MarkupPercentage: whereHelperint{field: "\"tax_rates\".\"markup_percentage\""},
	Name:             whereHelperstring{field: "\"tax_rates\".\"name\""},
	CreatedAt:        whereHelpertime_Time{field: "\"tax_rates\".\"created_at\""},
}

// TaxRateRels is where relationship names are stored.
var TaxRateRels = struct {
	Orders                     string
	DefaultTaxRateUserBillings string
}{
	Orders:                     "Orders",
	DefaultTaxRateUserBillings: "DefaultTaxRateUserBillings",
}

// taxRateR is where relationships are stored.
type taxRateR struct {
	Orders                     OrderSlice       `boil:"Orders" json:"Orders" toml:"Orders" yaml:"Orders"`
	DefaultTaxRateUserBillings UserBillingSlice `boil:"DefaultTaxRateUserBillings" json:"DefaultTaxRateUserBillings" toml:"DefaultTaxRateUserBillings" yaml:"DefaultTaxRateUserBillings"`
}

// NewStruct creates a new relationship struct
func (*taxRateR) NewStruct() *taxRateR {
	return &taxRateR{}
}

func (r *taxRateR) GetOrders() OrderSlice {
	if r == nil {
		return nil
	}
	return r.Orders
}

func (r *taxRateR) GetDefaultTaxRateUserBillings() UserBillingSlice {
	if r == nil {
		return nil
	}
	return r.DefaultTaxRateUserBillings
}

// taxRateL is where Load methods for each relationship are stored.
type taxRateL struct{}

var (
	taxRateAllColumns            = []string{"id", "markup_percentage", "name", "created_at"}
	taxRateColumnsWithoutDefault = []string{"markup_percentage", "name"}
	taxRateColumnsWithDefault    = []string{"id", "created_at"}
	taxRatePrimaryKeyColumns     = []string{"id"}
	taxRateGeneratedColumns      = []string{}
)

type (
	// TaxRateSlice is an alias for a slice of pointers to TaxRate.
	// This should almost always be used instead of []TaxRate.
	TaxRateSlice []*TaxRate
	// TaxRateHook is the signature for custom TaxRate hook methods
	TaxRateHook func(context.Context, boil.ContextExecutor, *TaxRate) error

	taxRateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taxRateType                 = reflect.TypeOf(&TaxRate{})
	taxRateMapping              = queries.MakeStructMapping(taxRateType)
	taxRatePrimaryKeyMapping, _ = queries.BindMapping(taxRateType, taxRateMapping, taxRatePrimaryKeyColumns)
	taxRateInsertCacheMut       sync.RWMutex
	taxRateInsertCache          = make(map[string]insertCache)
	taxRateUpdateCacheMut       sync.RWMutex
	taxRateUpdateCache          = make(map[string]updateCache)
	taxRateUpsertCacheMut       sync.RWMutex
	taxRateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var taxRateAfterSelectHooks []TaxRateHook

var taxRateBeforeInsertHooks []TaxRateHook
var taxRateAfterInsertHooks []TaxRateHook

var taxRateBeforeUpdateHooks []TaxRateHook
var taxRateAfterUpdateHooks []TaxRateHook

var taxRateBeforeDeleteHooks []TaxRateHook
var taxRateAfterDeleteHooks []TaxRateHook

var taxRateBeforeUpsertHooks []TaxRateHook
var taxRateAfterUpsertHooks []TaxRateHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TaxRate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TaxRate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TaxRate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TaxRate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TaxRate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TaxRate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TaxRate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TaxRate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TaxRate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxRateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTaxRateHook registers your hook function for all future operations.
func AddTaxRateHook(hookPoint boil.HookPoint, taxRateHook TaxRateHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		taxRateAfterSelectHooks = append(taxRateAfterSelectHooks, taxRateHook)
	case boil.BeforeInsertHook:
		taxRateBeforeInsertHooks = append(taxRateBeforeInsertHooks, taxRateHook)
	case boil.AfterInsertHook:
		taxRateAfterInsertHooks = append(taxRateAfterInsertHooks, taxRateHook)
	case boil.BeforeUpdateHook:
		taxRateBeforeUpdateHooks = append(taxRateBeforeUpdateHooks, taxRateHook)
	case boil.AfterUpdateHook:
		taxRateAfterUpdateHooks = append(taxRateAfterUpdateHooks, taxRateHook)
	case boil.BeforeDeleteHook:
		taxRateBeforeDeleteHooks = append(taxRateBeforeDeleteHooks, taxRateHook)
	case boil.AfterDeleteHook:
		taxRateAfterDeleteHooks = append(taxRateAfterDeleteHooks, taxRateHook)
	case boil.BeforeUpsertHook:
		taxRateBeforeUpsertHooks = append(taxRateBeforeUpsertHooks, taxRateHook)
	case boil.AfterUpsertHook:
		taxRateAfterUpsertHooks = append(taxRateAfterUpsertHooks, taxRateHook)
	}
}

// One returns a single taxRate record from the query.
func (q taxRateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TaxRate, error) {
	o := &TaxRate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tax_rates")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TaxRate records from the query.
func (q taxRateQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaxRateSlice, error) {
	var o []*TaxRate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TaxRate slice")
	}

	if len(taxRateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TaxRate records in the query.
func (q taxRateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tax_rates rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q taxRateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tax_rates exists")
	}

	return count > 0, nil
}

// Orders retrieves all the order's Orders with an executor.
func (o *TaxRate) Orders(mods ...qm.QueryMod) orderQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"orders\".\"tax_rate_id\"=?", o.ID),
	)

	return Orders(queryMods...)
}

// DefaultTaxRateUserBillings retrieves all the user_billing's UserBillings with an executor via default_tax_rate_id column.
func (o *TaxRate) DefaultTaxRateUserBillings(mods ...qm.QueryMod) userBillingQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_billings\".\"default_tax_rate_id\"=?", o.ID),
	)

	return UserBillings(queryMods...)
}

// LoadOrders allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (taxRateL) LoadOrders(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTaxRate interface{}, mods queries.Applicator) error {
	var slice []*TaxRate
	var object *TaxRate

	if singular {
		var ok bool
		object, ok = maybeTaxRate.(*TaxRate)
		if !ok {
			object = new(TaxRate)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTaxRate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTaxRate))
			}
		}
	} else {
		s, ok := maybeTaxRate.(*[]*TaxRate)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTaxRate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTaxRate))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taxRateR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taxRateR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`orders`),
		qm.WhereIn(`orders.tax_rate_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load orders")
	}

	var resultSlice []*Order
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice orders")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on orders")
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
	if singular {
		object.R.Orders = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderR{}
			}
			foreign.R.TaxRate = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TaxRateID {
				local.R.Orders = append(local.R.Orders, foreign)
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.TaxRate = local
				break
			}
		}
	}

	return nil
}

// LoadDefaultTaxRateUserBillings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (taxRateL) LoadDefaultTaxRateUserBillings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTaxRate interface{}, mods queries.Applicator) error {
	var slice []*TaxRate
	var object *TaxRate

	if singular {
		var ok bool
		object, ok = maybeTaxRate.(*TaxRate)
		if !ok {
			object = new(TaxRate)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTaxRate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTaxRate))
			}
		}
	} else {
		s, ok := maybeTaxRate.(*[]*TaxRate)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTaxRate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTaxRate))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taxRateR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taxRateR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_billings`),
		qm.WhereIn(`user_billings.default_tax_rate_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_billings")
	}

	var resultSlice []*UserBilling
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_billings")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_billings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_billings")
	}

	if len(userBillingAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.DefaultTaxRateUserBillings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userBillingR{}
			}
			foreign.R.DefaultTaxRate = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.DefaultTaxRateID {
				local.R.DefaultTaxRateUserBillings = append(local.R.DefaultTaxRateUserBillings, foreign)
				if foreign.R == nil {
					foreign.R = &userBillingR{}
				}
				foreign.R.DefaultTaxRate = local
				break
			}
		}
	}

	return nil
}

// AddOrders adds the given related objects to the existing relationships
// of the tax_rate, optionally inserting them as new records.
// Appends related to o.R.Orders.
// Sets related.R.TaxRate appropriately.
func (o *TaxRate) AddOrders(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Order) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TaxRateID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"orders\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"tax_rate_id"}),
				strmangle.WhereClause("\"", "\"", 2, orderPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TaxRateID = o.ID
		}
	}

	if o.R == nil {
		o.R = &taxRateR{
			Orders: related,
		}
	} else {
		o.R.Orders = append(o.R.Orders, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderR{
				TaxRate: o,
			}
		} else {
			rel.R.TaxRate = o
		}
	}
	return nil
}

// AddDefaultTaxRateUserBillings adds the given related objects to the existing relationships
// of the tax_rate, optionally inserting them as new records.
// Appends related to o.R.DefaultTaxRateUserBillings.
// Sets related.R.DefaultTaxRate appropriately.
func (o *TaxRate) AddDefaultTaxRateUserBillings(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserBilling) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DefaultTaxRateID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_billings\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"default_tax_rate_id"}),
				strmangle.WhereClause("\"", "\"", 2, userBillingPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DefaultTaxRateID = o.ID
		}
	}

	if o.R == nil {
		o.R = &taxRateR{
			DefaultTaxRateUserBillings: related,
		}
	} else {
		o.R.DefaultTaxRateUserBillings = append(o.R.DefaultTaxRateUserBillings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userBillingR{
				DefaultTaxRate: o,
			}
		} else {
			rel.R.DefaultTaxRate = o
		}
	}
	return nil
}

// TaxRates retrieves all the records using an executor.
func TaxRates(mods ...qm.QueryMod) taxRateQuery {
	mods = append(mods, qm.From("\"tax_rates\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"tax_rates\".*"})
	}

	return taxRateQuery{q}
}

// FindTaxRate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTaxRate(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TaxRate, error) {
	taxRateObj := &TaxRate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tax_rates\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, taxRateObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tax_rates")
	}

	if err = taxRateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return taxRateObj, err
	}

	return taxRateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TaxRate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tax_rates provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taxRateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taxRateInsertCacheMut.RLock()
	cache, cached := taxRateInsertCache[key]
	taxRateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taxRateAllColumns,
			taxRateColumnsWithDefault,
			taxRateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(taxRateType, taxRateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taxRateType, taxRateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tax_rates\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tax_rates\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into tax_rates")
	}

	if !cached {
		taxRateInsertCacheMut.Lock()
		taxRateInsertCache[key] = cache
		taxRateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TaxRate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TaxRate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	taxRateUpdateCacheMut.RLock()
	cache, cached := taxRateUpdateCache[key]
	taxRateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taxRateAllColumns,
			taxRatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tax_rates, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tax_rates\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, taxRatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taxRateType, taxRateMapping, append(wl, taxRatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update tax_rates row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tax_rates")
	}

	if !cached {
		taxRateUpdateCacheMut.Lock()
		taxRateUpdateCache[key] = cache
		taxRateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q taxRateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tax_rates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tax_rates")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaxRateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxRatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tax_rates\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, taxRatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in taxRate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all taxRate")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TaxRate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tax_rates provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taxRateColumnsWithDefault, o)

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

	taxRateUpsertCacheMut.RLock()
	cache, cached := taxRateUpsertCache[key]
	taxRateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taxRateAllColumns,
			taxRateColumnsWithDefault,
			taxRateColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			taxRateAllColumns,
			taxRatePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert tax_rates, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(taxRatePrimaryKeyColumns))
			copy(conflict, taxRatePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tax_rates\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(taxRateType, taxRateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taxRateType, taxRateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert tax_rates")
	}

	if !cached {
		taxRateUpsertCacheMut.Lock()
		taxRateUpsertCache[key] = cache
		taxRateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TaxRate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TaxRate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TaxRate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taxRatePrimaryKeyMapping)
	sql := "DELETE FROM \"tax_rates\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tax_rates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tax_rates")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q taxRateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no taxRateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tax_rates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tax_rates")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaxRateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(taxRateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxRatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tax_rates\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taxRatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from taxRate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tax_rates")
	}

	if len(taxRateAfterDeleteHooks) != 0 {
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
func (o *TaxRate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTaxRate(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaxRateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaxRateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxRatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tax_rates\".* FROM \"tax_rates\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taxRatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TaxRateSlice")
	}

	*o = slice

	return nil
}

// TaxRateExists checks if the TaxRate row exists.
func TaxRateExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tax_rates\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tax_rates exists")
	}

	return exists, nil
}

// Exists checks if the TaxRate row exists.
func (o *TaxRate) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TaxRateExists(ctx, exec, o.ID)
}

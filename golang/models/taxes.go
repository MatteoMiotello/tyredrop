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

// Taxis is an object representing the database table.
type Taxis struct {
	ID               int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	MarkupPercentage int       `boil:"markup_percentage" json:"markup_percentage" toml:"markup_percentage" yaml:"markup_percentage"`
	Name             string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *taxisR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taxisL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaxisColumns = struct {
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

var TaxisTableColumns = struct {
	ID               string
	MarkupPercentage string
	Name             string
	CreatedAt        string
}{
	ID:               "taxes.id",
	MarkupPercentage: "taxes.markup_percentage",
	Name:             "taxes.name",
	CreatedAt:        "taxes.created_at",
}

// Generated where

var TaxisWhere = struct {
	ID               whereHelperint64
	MarkupPercentage whereHelperint
	Name             whereHelperstring
	CreatedAt        whereHelpertime_Time
}{
	ID:               whereHelperint64{field: "\"taxes\".\"id\""},
	MarkupPercentage: whereHelperint{field: "\"taxes\".\"markup_percentage\""},
	Name:             whereHelperstring{field: "\"taxes\".\"name\""},
	CreatedAt:        whereHelpertime_Time{field: "\"taxes\".\"created_at\""},
}

// TaxisRels is where relationship names are stored.
var TaxisRels = struct {
	TaxOrders string
}{
	TaxOrders: "TaxOrders",
}

// taxisR is where relationships are stored.
type taxisR struct {
	TaxOrders OrderSlice `boil:"TaxOrders" json:"TaxOrders" toml:"TaxOrders" yaml:"TaxOrders"`
}

// NewStruct creates a new relationship struct
func (*taxisR) NewStruct() *taxisR {
	return &taxisR{}
}

func (r *taxisR) GetTaxOrders() OrderSlice {
	if r == nil {
		return nil
	}
	return r.TaxOrders
}

// taxisL is where Load methods for each relationship are stored.
type taxisL struct{}

var (
	taxisAllColumns            = []string{"id", "markup_percentage", "name", "created_at"}
	taxisColumnsWithoutDefault = []string{"markup_percentage", "name"}
	taxisColumnsWithDefault    = []string{"id", "created_at"}
	taxisPrimaryKeyColumns     = []string{"id"}
	taxisGeneratedColumns      = []string{}
)

type (
	// TaxisSlice is an alias for a slice of pointers to Taxis.
	// This should almost always be used instead of []Taxis.
	TaxisSlice []*Taxis
	// TaxisHook is the signature for custom Taxis hook methods
	TaxisHook func(context.Context, boil.ContextExecutor, *Taxis) error

	taxisQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taxisType                 = reflect.TypeOf(&Taxis{})
	taxisMapping              = queries.MakeStructMapping(taxisType)
	taxisPrimaryKeyMapping, _ = queries.BindMapping(taxisType, taxisMapping, taxisPrimaryKeyColumns)
	taxisInsertCacheMut       sync.RWMutex
	taxisInsertCache          = make(map[string]insertCache)
	taxisUpdateCacheMut       sync.RWMutex
	taxisUpdateCache          = make(map[string]updateCache)
	taxisUpsertCacheMut       sync.RWMutex
	taxisUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var taxisAfterSelectHooks []TaxisHook

var taxisBeforeInsertHooks []TaxisHook
var taxisAfterInsertHooks []TaxisHook

var taxisBeforeUpdateHooks []TaxisHook
var taxisAfterUpdateHooks []TaxisHook

var taxisBeforeDeleteHooks []TaxisHook
var taxisAfterDeleteHooks []TaxisHook

var taxisBeforeUpsertHooks []TaxisHook
var taxisAfterUpsertHooks []TaxisHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Taxis) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Taxis) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Taxis) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Taxis) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Taxis) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Taxis) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Taxis) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Taxis) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Taxis) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxisAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTaxisHook registers your hook function for all future operations.
func AddTaxisHook(hookPoint boil.HookPoint, taxisHook TaxisHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		taxisAfterSelectHooks = append(taxisAfterSelectHooks, taxisHook)
	case boil.BeforeInsertHook:
		taxisBeforeInsertHooks = append(taxisBeforeInsertHooks, taxisHook)
	case boil.AfterInsertHook:
		taxisAfterInsertHooks = append(taxisAfterInsertHooks, taxisHook)
	case boil.BeforeUpdateHook:
		taxisBeforeUpdateHooks = append(taxisBeforeUpdateHooks, taxisHook)
	case boil.AfterUpdateHook:
		taxisAfterUpdateHooks = append(taxisAfterUpdateHooks, taxisHook)
	case boil.BeforeDeleteHook:
		taxisBeforeDeleteHooks = append(taxisBeforeDeleteHooks, taxisHook)
	case boil.AfterDeleteHook:
		taxisAfterDeleteHooks = append(taxisAfterDeleteHooks, taxisHook)
	case boil.BeforeUpsertHook:
		taxisBeforeUpsertHooks = append(taxisBeforeUpsertHooks, taxisHook)
	case boil.AfterUpsertHook:
		taxisAfterUpsertHooks = append(taxisAfterUpsertHooks, taxisHook)
	}
}

// One returns a single taxis record from the query.
func (q taxisQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Taxis, error) {
	o := &Taxis{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for taxes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Taxis records from the query.
func (q taxisQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaxisSlice, error) {
	var o []*Taxis

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Taxis slice")
	}

	if len(taxisAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Taxis records in the query.
func (q taxisQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count taxes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q taxisQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if taxes exists")
	}

	return count > 0, nil
}

// TaxOrders retrieves all the order's Orders with an executor via tax_id column.
func (o *Taxis) TaxOrders(mods ...qm.QueryMod) orderQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"orders\".\"tax_id\"=?", o.ID),
	)

	return Orders(queryMods...)
}

// LoadTaxOrders allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (taxisL) LoadTaxOrders(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTaxis interface{}, mods queries.Applicator) error {
	var slice []*Taxis
	var object *Taxis

	if singular {
		var ok bool
		object, ok = maybeTaxis.(*Taxis)
		if !ok {
			object = new(Taxis)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTaxis)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTaxis))
			}
		}
	} else {
		s, ok := maybeTaxis.(*[]*Taxis)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTaxis)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTaxis))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taxisR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taxisR{}
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
		qm.WhereIn(`orders.tax_id in ?`, args...),
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
		object.R.TaxOrders = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderR{}
			}
			foreign.R.Tax = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TaxID {
				local.R.TaxOrders = append(local.R.TaxOrders, foreign)
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.Tax = local
				break
			}
		}
	}

	return nil
}

// AddTaxOrders adds the given related objects to the existing relationships
// of the taxis, optionally inserting them as new records.
// Appends related to o.R.TaxOrders.
// Sets related.R.Tax appropriately.
func (o *Taxis) AddTaxOrders(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Order) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TaxID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"orders\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"tax_id"}),
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

			rel.TaxID = o.ID
		}
	}

	if o.R == nil {
		o.R = &taxisR{
			TaxOrders: related,
		}
	} else {
		o.R.TaxOrders = append(o.R.TaxOrders, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderR{
				Tax: o,
			}
		} else {
			rel.R.Tax = o
		}
	}
	return nil
}

// Taxes retrieves all the records using an executor.
func Taxes(mods ...qm.QueryMod) taxisQuery {
	mods = append(mods, qm.From("\"taxes\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"taxes\".*"})
	}

	return taxisQuery{q}
}

// FindTaxis retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTaxis(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Taxis, error) {
	taxisObj := &Taxis{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"taxes\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, taxisObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from taxes")
	}

	if err = taxisObj.doAfterSelectHooks(ctx, exec); err != nil {
		return taxisObj, err
	}

	return taxisObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Taxis) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no taxes provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(taxisColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taxisInsertCacheMut.RLock()
	cache, cached := taxisInsertCache[key]
	taxisInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taxisAllColumns,
			taxisColumnsWithDefault,
			taxisColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(taxisType, taxisMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taxisType, taxisMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"taxes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"taxes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into taxes")
	}

	if !cached {
		taxisInsertCacheMut.Lock()
		taxisInsertCache[key] = cache
		taxisInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Taxis.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Taxis) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	taxisUpdateCacheMut.RLock()
	cache, cached := taxisUpdateCache[key]
	taxisUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taxisAllColumns,
			taxisPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update taxes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"taxes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, taxisPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taxisType, taxisMapping, append(wl, taxisPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update taxes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for taxes")
	}

	if !cached {
		taxisUpdateCacheMut.Lock()
		taxisUpdateCache[key] = cache
		taxisUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q taxisQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for taxes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for taxes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaxisSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxisPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"taxes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, taxisPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in taxis slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all taxis")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Taxis) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no taxes provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(taxisColumnsWithDefault, o)

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

	taxisUpsertCacheMut.RLock()
	cache, cached := taxisUpsertCache[key]
	taxisUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taxisAllColumns,
			taxisColumnsWithDefault,
			taxisColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			taxisAllColumns,
			taxisPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert taxes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(taxisPrimaryKeyColumns))
			copy(conflict, taxisPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"taxes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(taxisType, taxisMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taxisType, taxisMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert taxes")
	}

	if !cached {
		taxisUpsertCacheMut.Lock()
		taxisUpsertCache[key] = cache
		taxisUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Taxis record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Taxis) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Taxis provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taxisPrimaryKeyMapping)
	sql := "DELETE FROM \"taxes\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from taxes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for taxes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q taxisQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no taxisQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from taxes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for taxes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaxisSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(taxisBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxisPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"taxes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taxisPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from taxis slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for taxes")
	}

	if len(taxisAfterDeleteHooks) != 0 {
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
func (o *Taxis) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTaxis(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaxisSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaxisSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxisPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"taxes\".* FROM \"taxes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taxisPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TaxisSlice")
	}

	*o = slice

	return nil
}

// TaxisExists checks if the Taxis row exists.
func TaxisExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"taxes\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if taxes exists")
	}

	return exists, nil
}

// Exists checks if the Taxis row exists.
func (o *Taxis) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TaxisExists(ctx, exec, o.ID)
}

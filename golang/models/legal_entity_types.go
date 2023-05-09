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

// LegalEntityType is an object representing the database table.
type LegalEntityType struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsPerson  bool      `boil:"is_person" json:"is_person" toml:"is_person" yaml:"is_person"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *legalEntityTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L legalEntityTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LegalEntityTypeColumns = struct {
	ID        string
	Name      string
	IsPerson  string
	CreatedAt string
}{
	ID:        "id",
	Name:      "name",
	IsPerson:  "is_person",
	CreatedAt: "created_at",
}

var LegalEntityTypeTableColumns = struct {
	ID        string
	Name      string
	IsPerson  string
	CreatedAt string
}{
	ID:        "legal_entity_types.id",
	Name:      "legal_entity_types.name",
	IsPerson:  "legal_entity_types.is_person",
	CreatedAt: "legal_entity_types.created_at",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var LegalEntityTypeWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	IsPerson  whereHelperbool
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"legal_entity_types\".\"id\""},
	Name:      whereHelperstring{field: "\"legal_entity_types\".\"name\""},
	IsPerson:  whereHelperbool{field: "\"legal_entity_types\".\"is_person\""},
	CreatedAt: whereHelpertime_Time{field: "\"legal_entity_types\".\"created_at\""},
}

// LegalEntityTypeRels is where relationship names are stored.
var LegalEntityTypeRels = struct {
	UserBillings string
}{
	UserBillings: "UserBillings",
}

// legalEntityTypeR is where relationships are stored.
type legalEntityTypeR struct {
	UserBillings UserBillingSlice `boil:"UserBillings" json:"UserBillings" toml:"UserBillings" yaml:"UserBillings"`
}

// NewStruct creates a new relationship struct
func (*legalEntityTypeR) NewStruct() *legalEntityTypeR {
	return &legalEntityTypeR{}
}

func (r *legalEntityTypeR) GetUserBillings() UserBillingSlice {
	if r == nil {
		return nil
	}
	return r.UserBillings
}

// legalEntityTypeL is where Load methods for each relationship are stored.
type legalEntityTypeL struct{}

var (
	legalEntityTypeAllColumns            = []string{"id", "name", "is_person", "created_at"}
	legalEntityTypeColumnsWithoutDefault = []string{"name", "is_person"}
	legalEntityTypeColumnsWithDefault    = []string{"id", "created_at"}
	legalEntityTypePrimaryKeyColumns     = []string{"id"}
	legalEntityTypeGeneratedColumns      = []string{}
)

type (
	// LegalEntityTypeSlice is an alias for a slice of pointers to LegalEntityType.
	// This should almost always be used instead of []LegalEntityType.
	LegalEntityTypeSlice []*LegalEntityType
	// LegalEntityTypeHook is the signature for custom LegalEntityType hook methods
	LegalEntityTypeHook func(context.Context, boil.ContextExecutor, *LegalEntityType) error

	legalEntityTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	legalEntityTypeType                 = reflect.TypeOf(&LegalEntityType{})
	legalEntityTypeMapping              = queries.MakeStructMapping(legalEntityTypeType)
	legalEntityTypePrimaryKeyMapping, _ = queries.BindMapping(legalEntityTypeType, legalEntityTypeMapping, legalEntityTypePrimaryKeyColumns)
	legalEntityTypeInsertCacheMut       sync.RWMutex
	legalEntityTypeInsertCache          = make(map[string]insertCache)
	legalEntityTypeUpdateCacheMut       sync.RWMutex
	legalEntityTypeUpdateCache          = make(map[string]updateCache)
	legalEntityTypeUpsertCacheMut       sync.RWMutex
	legalEntityTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var legalEntityTypeAfterSelectHooks []LegalEntityTypeHook

var legalEntityTypeBeforeInsertHooks []LegalEntityTypeHook
var legalEntityTypeAfterInsertHooks []LegalEntityTypeHook

var legalEntityTypeBeforeUpdateHooks []LegalEntityTypeHook
var legalEntityTypeAfterUpdateHooks []LegalEntityTypeHook

var legalEntityTypeBeforeDeleteHooks []LegalEntityTypeHook
var legalEntityTypeAfterDeleteHooks []LegalEntityTypeHook

var legalEntityTypeBeforeUpsertHooks []LegalEntityTypeHook
var legalEntityTypeAfterUpsertHooks []LegalEntityTypeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LegalEntityType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LegalEntityType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LegalEntityType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LegalEntityType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LegalEntityType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LegalEntityType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LegalEntityType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LegalEntityType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LegalEntityType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range legalEntityTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLegalEntityTypeHook registers your hook function for all future operations.
func AddLegalEntityTypeHook(hookPoint boil.HookPoint, legalEntityTypeHook LegalEntityTypeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		legalEntityTypeAfterSelectHooks = append(legalEntityTypeAfterSelectHooks, legalEntityTypeHook)
	case boil.BeforeInsertHook:
		legalEntityTypeBeforeInsertHooks = append(legalEntityTypeBeforeInsertHooks, legalEntityTypeHook)
	case boil.AfterInsertHook:
		legalEntityTypeAfterInsertHooks = append(legalEntityTypeAfterInsertHooks, legalEntityTypeHook)
	case boil.BeforeUpdateHook:
		legalEntityTypeBeforeUpdateHooks = append(legalEntityTypeBeforeUpdateHooks, legalEntityTypeHook)
	case boil.AfterUpdateHook:
		legalEntityTypeAfterUpdateHooks = append(legalEntityTypeAfterUpdateHooks, legalEntityTypeHook)
	case boil.BeforeDeleteHook:
		legalEntityTypeBeforeDeleteHooks = append(legalEntityTypeBeforeDeleteHooks, legalEntityTypeHook)
	case boil.AfterDeleteHook:
		legalEntityTypeAfterDeleteHooks = append(legalEntityTypeAfterDeleteHooks, legalEntityTypeHook)
	case boil.BeforeUpsertHook:
		legalEntityTypeBeforeUpsertHooks = append(legalEntityTypeBeforeUpsertHooks, legalEntityTypeHook)
	case boil.AfterUpsertHook:
		legalEntityTypeAfterUpsertHooks = append(legalEntityTypeAfterUpsertHooks, legalEntityTypeHook)
	}
}

// One returns a single legalEntityType record from the query.
func (q legalEntityTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LegalEntityType, error) {
	o := &LegalEntityType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for legal_entity_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LegalEntityType records from the query.
func (q legalEntityTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (LegalEntityTypeSlice, error) {
	var o []*LegalEntityType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LegalEntityType slice")
	}

	if len(legalEntityTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LegalEntityType records in the query.
func (q legalEntityTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count legal_entity_types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q legalEntityTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if legal_entity_types exists")
	}

	return count > 0, nil
}

// UserBillings retrieves all the user_billing's UserBillings with an executor.
func (o *LegalEntityType) UserBillings(mods ...qm.QueryMod) userBillingQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_billings\".\"legal_entity_type_id\"=?", o.ID),
	)

	return UserBillings(queryMods...)
}

// LoadUserBillings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (legalEntityTypeL) LoadUserBillings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLegalEntityType interface{}, mods queries.Applicator) error {
	var slice []*LegalEntityType
	var object *LegalEntityType

	if singular {
		var ok bool
		object, ok = maybeLegalEntityType.(*LegalEntityType)
		if !ok {
			object = new(LegalEntityType)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeLegalEntityType)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeLegalEntityType))
			}
		}
	} else {
		s, ok := maybeLegalEntityType.(*[]*LegalEntityType)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeLegalEntityType)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeLegalEntityType))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &legalEntityTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &legalEntityTypeR{}
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
		qm.WhereIn(`user_billings.legal_entity_type_id in ?`, args...),
		qmhelper.WhereIsNull(`user_billings.deleted_at`),
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
		object.R.UserBillings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userBillingR{}
			}
			foreign.R.LegalEntityType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.LegalEntityTypeID {
				local.R.UserBillings = append(local.R.UserBillings, foreign)
				if foreign.R == nil {
					foreign.R = &userBillingR{}
				}
				foreign.R.LegalEntityType = local
				break
			}
		}
	}

	return nil
}

// AddUserBillings adds the given related objects to the existing relationships
// of the legal_entity_type, optionally inserting them as new records.
// Appends related to o.R.UserBillings.
// Sets related.R.LegalEntityType appropriately.
func (o *LegalEntityType) AddUserBillings(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserBilling) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.LegalEntityTypeID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_billings\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"legal_entity_type_id"}),
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

			rel.LegalEntityTypeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &legalEntityTypeR{
			UserBillings: related,
		}
	} else {
		o.R.UserBillings = append(o.R.UserBillings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userBillingR{
				LegalEntityType: o,
			}
		} else {
			rel.R.LegalEntityType = o
		}
	}
	return nil
}

// LegalEntityTypes retrieves all the records using an executor.
func LegalEntityTypes(mods ...qm.QueryMod) legalEntityTypeQuery {
	mods = append(mods, qm.From("\"legal_entity_types\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"legal_entity_types\".*"})
	}

	return legalEntityTypeQuery{q}
}

// FindLegalEntityType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLegalEntityType(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*LegalEntityType, error) {
	legalEntityTypeObj := &LegalEntityType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"legal_entity_types\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, legalEntityTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from legal_entity_types")
	}

	if err = legalEntityTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return legalEntityTypeObj, err
	}

	return legalEntityTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LegalEntityType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no legal_entity_types provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(legalEntityTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	legalEntityTypeInsertCacheMut.RLock()
	cache, cached := legalEntityTypeInsertCache[key]
	legalEntityTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			legalEntityTypeAllColumns,
			legalEntityTypeColumnsWithDefault,
			legalEntityTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(legalEntityTypeType, legalEntityTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(legalEntityTypeType, legalEntityTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"legal_entity_types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"legal_entity_types\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into legal_entity_types")
	}

	if !cached {
		legalEntityTypeInsertCacheMut.Lock()
		legalEntityTypeInsertCache[key] = cache
		legalEntityTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the LegalEntityType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LegalEntityType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	legalEntityTypeUpdateCacheMut.RLock()
	cache, cached := legalEntityTypeUpdateCache[key]
	legalEntityTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			legalEntityTypeAllColumns,
			legalEntityTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update legal_entity_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"legal_entity_types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, legalEntityTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(legalEntityTypeType, legalEntityTypeMapping, append(wl, legalEntityTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update legal_entity_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for legal_entity_types")
	}

	if !cached {
		legalEntityTypeUpdateCacheMut.Lock()
		legalEntityTypeUpdateCache[key] = cache
		legalEntityTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q legalEntityTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for legal_entity_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for legal_entity_types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LegalEntityTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), legalEntityTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"legal_entity_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, legalEntityTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in legalEntityType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all legalEntityType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LegalEntityType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no legal_entity_types provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(legalEntityTypeColumnsWithDefault, o)

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

	legalEntityTypeUpsertCacheMut.RLock()
	cache, cached := legalEntityTypeUpsertCache[key]
	legalEntityTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			legalEntityTypeAllColumns,
			legalEntityTypeColumnsWithDefault,
			legalEntityTypeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			legalEntityTypeAllColumns,
			legalEntityTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert legal_entity_types, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(legalEntityTypePrimaryKeyColumns))
			copy(conflict, legalEntityTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"legal_entity_types\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(legalEntityTypeType, legalEntityTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(legalEntityTypeType, legalEntityTypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert legal_entity_types")
	}

	if !cached {
		legalEntityTypeUpsertCacheMut.Lock()
		legalEntityTypeUpsertCache[key] = cache
		legalEntityTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single LegalEntityType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LegalEntityType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LegalEntityType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), legalEntityTypePrimaryKeyMapping)
	sql := "DELETE FROM \"legal_entity_types\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from legal_entity_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for legal_entity_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q legalEntityTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no legalEntityTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from legal_entity_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for legal_entity_types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LegalEntityTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(legalEntityTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), legalEntityTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"legal_entity_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, legalEntityTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from legalEntityType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for legal_entity_types")
	}

	if len(legalEntityTypeAfterDeleteHooks) != 0 {
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
func (o *LegalEntityType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLegalEntityType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LegalEntityTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LegalEntityTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), legalEntityTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"legal_entity_types\".* FROM \"legal_entity_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, legalEntityTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LegalEntityTypeSlice")
	}

	*o = slice

	return nil
}

// LegalEntityTypeExists checks if the LegalEntityType row exists.
func LegalEntityTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"legal_entity_types\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if legal_entity_types exists")
	}

	return exists, nil
}

// Exists checks if the LegalEntityType row exists.
func (o *LegalEntityType) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LegalEntityTypeExists(ctx, exec, o.ID)
}

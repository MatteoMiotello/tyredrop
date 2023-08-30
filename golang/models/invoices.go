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

// Invoice is an object representing the database table.
type Invoice struct {
	ID            int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserBillingID int64     `boil:"user_billing_id" json:"user_billing_id" toml:"user_billing_id" yaml:"user_billing_id"`
	Number        string    `boil:"number" json:"number" toml:"number" yaml:"number"`
	FilePath      string    `boil:"file_path" json:"file_path" toml:"file_path" yaml:"file_path"`
	DeletedAt     null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	UpdatedAt     time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *invoiceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L invoiceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InvoiceColumns = struct {
	ID            string
	UserBillingID string
	Number        string
	FilePath      string
	DeletedAt     string
	UpdatedAt     string
	CreatedAt     string
}{
	ID:            "id",
	UserBillingID: "user_billing_id",
	Number:        "number",
	FilePath:      "file_path",
	DeletedAt:     "deleted_at",
	UpdatedAt:     "updated_at",
	CreatedAt:     "created_at",
}

var InvoiceTableColumns = struct {
	ID            string
	UserBillingID string
	Number        string
	FilePath      string
	DeletedAt     string
	UpdatedAt     string
	CreatedAt     string
}{
	ID:            "invoices.id",
	UserBillingID: "invoices.user_billing_id",
	Number:        "invoices.number",
	FilePath:      "invoices.file_path",
	DeletedAt:     "invoices.deleted_at",
	UpdatedAt:     "invoices.updated_at",
	CreatedAt:     "invoices.created_at",
}

// Generated where

var InvoiceWhere = struct {
	ID            whereHelperint64
	UserBillingID whereHelperint64
	Number        whereHelperstring
	FilePath      whereHelperstring
	DeletedAt     whereHelpernull_Time
	UpdatedAt     whereHelpertime_Time
	CreatedAt     whereHelpertime_Time
}{
	ID:            whereHelperint64{field: "\"invoices\".\"id\""},
	UserBillingID: whereHelperint64{field: "\"invoices\".\"user_billing_id\""},
	Number:        whereHelperstring{field: "\"invoices\".\"number\""},
	FilePath:      whereHelperstring{field: "\"invoices\".\"file_path\""},
	DeletedAt:     whereHelpernull_Time{field: "\"invoices\".\"deleted_at\""},
	UpdatedAt:     whereHelpertime_Time{field: "\"invoices\".\"updated_at\""},
	CreatedAt:     whereHelpertime_Time{field: "\"invoices\".\"created_at\""},
}

// InvoiceRels is where relationship names are stored.
var InvoiceRels = struct {
	UserBilling string
}{
	UserBilling: "UserBilling",
}

// invoiceR is where relationships are stored.
type invoiceR struct {
	UserBilling *UserBilling `boil:"UserBilling" json:"UserBilling" toml:"UserBilling" yaml:"UserBilling"`
}

// NewStruct creates a new relationship struct
func (*invoiceR) NewStruct() *invoiceR {
	return &invoiceR{}
}

func (r *invoiceR) GetUserBilling() *UserBilling {
	if r == nil {
		return nil
	}
	return r.UserBilling
}

// invoiceL is where Load methods for each relationship are stored.
type invoiceL struct{}

var (
	invoiceAllColumns            = []string{"id", "user_billing_id", "number", "file_path", "deleted_at", "updated_at", "created_at"}
	invoiceColumnsWithoutDefault = []string{"user_billing_id", "number", "file_path"}
	invoiceColumnsWithDefault    = []string{"id", "deleted_at", "updated_at", "created_at"}
	invoicePrimaryKeyColumns     = []string{"id"}
	invoiceGeneratedColumns      = []string{}
)

type (
	// InvoiceSlice is an alias for a slice of pointers to Invoice.
	// This should almost always be used instead of []Invoice.
	InvoiceSlice []*Invoice
	// InvoiceHook is the signature for custom Invoice hook methods
	InvoiceHook func(context.Context, boil.ContextExecutor, *Invoice) error

	invoiceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	invoiceType                 = reflect.TypeOf(&Invoice{})
	invoiceMapping              = queries.MakeStructMapping(invoiceType)
	invoicePrimaryKeyMapping, _ = queries.BindMapping(invoiceType, invoiceMapping, invoicePrimaryKeyColumns)
	invoiceInsertCacheMut       sync.RWMutex
	invoiceInsertCache          = make(map[string]insertCache)
	invoiceUpdateCacheMut       sync.RWMutex
	invoiceUpdateCache          = make(map[string]updateCache)
	invoiceUpsertCacheMut       sync.RWMutex
	invoiceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var invoiceAfterSelectHooks []InvoiceHook

var invoiceBeforeInsertHooks []InvoiceHook
var invoiceAfterInsertHooks []InvoiceHook

var invoiceBeforeUpdateHooks []InvoiceHook
var invoiceAfterUpdateHooks []InvoiceHook

var invoiceBeforeDeleteHooks []InvoiceHook
var invoiceAfterDeleteHooks []InvoiceHook

var invoiceBeforeUpsertHooks []InvoiceHook
var invoiceAfterUpsertHooks []InvoiceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Invoice) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Invoice) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Invoice) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Invoice) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Invoice) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Invoice) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Invoice) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Invoice) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Invoice) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInvoiceHook registers your hook function for all future operations.
func AddInvoiceHook(hookPoint boil.HookPoint, invoiceHook InvoiceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		invoiceAfterSelectHooks = append(invoiceAfterSelectHooks, invoiceHook)
	case boil.BeforeInsertHook:
		invoiceBeforeInsertHooks = append(invoiceBeforeInsertHooks, invoiceHook)
	case boil.AfterInsertHook:
		invoiceAfterInsertHooks = append(invoiceAfterInsertHooks, invoiceHook)
	case boil.BeforeUpdateHook:
		invoiceBeforeUpdateHooks = append(invoiceBeforeUpdateHooks, invoiceHook)
	case boil.AfterUpdateHook:
		invoiceAfterUpdateHooks = append(invoiceAfterUpdateHooks, invoiceHook)
	case boil.BeforeDeleteHook:
		invoiceBeforeDeleteHooks = append(invoiceBeforeDeleteHooks, invoiceHook)
	case boil.AfterDeleteHook:
		invoiceAfterDeleteHooks = append(invoiceAfterDeleteHooks, invoiceHook)
	case boil.BeforeUpsertHook:
		invoiceBeforeUpsertHooks = append(invoiceBeforeUpsertHooks, invoiceHook)
	case boil.AfterUpsertHook:
		invoiceAfterUpsertHooks = append(invoiceAfterUpsertHooks, invoiceHook)
	}
}

// One returns a single invoice record from the query.
func (q invoiceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Invoice, error) {
	o := &Invoice{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for invoices")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Invoice records from the query.
func (q invoiceQuery) All(ctx context.Context, exec boil.ContextExecutor) (InvoiceSlice, error) {
	var o []*Invoice

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Invoice slice")
	}

	if len(invoiceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Invoice records in the query.
func (q invoiceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count invoices rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q invoiceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if invoices exists")
	}

	return count > 0, nil
}

// UserBilling pointed to by the foreign key.
func (o *Invoice) UserBilling(mods ...qm.QueryMod) userBillingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserBillingID),
	}

	queryMods = append(queryMods, mods...)

	return UserBillings(queryMods...)
}

// LoadUserBilling allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (invoiceL) LoadUserBilling(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvoice interface{}, mods queries.Applicator) error {
	var slice []*Invoice
	var object *Invoice

	if singular {
		var ok bool
		object, ok = maybeInvoice.(*Invoice)
		if !ok {
			object = new(Invoice)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeInvoice)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeInvoice))
			}
		}
	} else {
		s, ok := maybeInvoice.(*[]*Invoice)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeInvoice)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeInvoice))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &invoiceR{}
		}
		args = append(args, object.UserBillingID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &invoiceR{}
			}

			for _, a := range args {
				if a == obj.UserBillingID {
					continue Outer
				}
			}

			args = append(args, obj.UserBillingID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_billings`),
		qm.WhereIn(`user_billings.id in ?`, args...),
		qmhelper.WhereIsNull(`user_billings.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserBilling")
	}

	var resultSlice []*UserBilling
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserBilling")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_billings")
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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.UserBilling = foreign
		if foreign.R == nil {
			foreign.R = &userBillingR{}
		}
		foreign.R.Invoices = append(foreign.R.Invoices, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserBillingID == foreign.ID {
				local.R.UserBilling = foreign
				if foreign.R == nil {
					foreign.R = &userBillingR{}
				}
				foreign.R.Invoices = append(foreign.R.Invoices, local)
				break
			}
		}
	}

	return nil
}

// SetUserBilling of the invoice to the related item.
// Sets o.R.UserBilling to related.
// Adds o to related.R.Invoices.
func (o *Invoice) SetUserBilling(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserBilling) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"invoices\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_billing_id"}),
		strmangle.WhereClause("\"", "\"", 2, invoicePrimaryKeyColumns),
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

	o.UserBillingID = related.ID
	if o.R == nil {
		o.R = &invoiceR{
			UserBilling: related,
		}
	} else {
		o.R.UserBilling = related
	}

	if related.R == nil {
		related.R = &userBillingR{
			Invoices: InvoiceSlice{o},
		}
	} else {
		related.R.Invoices = append(related.R.Invoices, o)
	}

	return nil
}

// Invoices retrieves all the records using an executor.
func Invoices(mods ...qm.QueryMod) invoiceQuery {
	mods = append(mods, qm.From("\"invoices\""), qmhelper.WhereIsNull("\"invoices\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"invoices\".*"})
	}

	return invoiceQuery{q}
}

// FindInvoice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvoice(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Invoice, error) {
	invoiceObj := &Invoice{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"invoices\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, invoiceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from invoices")
	}

	if err = invoiceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return invoiceObj, err
	}

	return invoiceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Invoice) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no invoices provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(invoiceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	invoiceInsertCacheMut.RLock()
	cache, cached := invoiceInsertCache[key]
	invoiceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			invoiceAllColumns,
			invoiceColumnsWithDefault,
			invoiceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(invoiceType, invoiceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(invoiceType, invoiceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invoices\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invoices\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into invoices")
	}

	if !cached {
		invoiceInsertCacheMut.Lock()
		invoiceInsertCache[key] = cache
		invoiceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Invoice.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Invoice) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	invoiceUpdateCacheMut.RLock()
	cache, cached := invoiceUpdateCache[key]
	invoiceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			invoiceAllColumns,
			invoicePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update invoices, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"invoices\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, invoicePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(invoiceType, invoiceMapping, append(wl, invoicePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update invoices row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for invoices")
	}

	if !cached {
		invoiceUpdateCacheMut.Lock()
		invoiceUpdateCache[key] = cache
		invoiceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q invoiceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for invoices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for invoices")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InvoiceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"invoices\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, invoicePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in invoice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all invoice")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Invoice) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no invoices provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(invoiceColumnsWithDefault, o)

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

	invoiceUpsertCacheMut.RLock()
	cache, cached := invoiceUpsertCache[key]
	invoiceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			invoiceAllColumns,
			invoiceColumnsWithDefault,
			invoiceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			invoiceAllColumns,
			invoicePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert invoices, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(invoicePrimaryKeyColumns))
			copy(conflict, invoicePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"invoices\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(invoiceType, invoiceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(invoiceType, invoiceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert invoices")
	}

	if !cached {
		invoiceUpsertCacheMut.Lock()
		invoiceUpsertCache[key] = cache
		invoiceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Invoice record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Invoice) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Invoice provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), invoicePrimaryKeyMapping)
		sql = "DELETE FROM \"invoices\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"invoices\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(invoiceType, invoiceMapping, append(wl, invoicePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from invoices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for invoices")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q invoiceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no invoiceQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invoices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invoices")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InvoiceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(invoiceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoicePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"invoices\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invoicePrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoicePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"invoices\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, invoicePrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invoice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invoices")
	}

	if len(invoiceAfterDeleteHooks) != 0 {
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
func (o *Invoice) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvoice(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InvoiceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InvoiceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"invoices\".* FROM \"invoices\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invoicePrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InvoiceSlice")
	}

	*o = slice

	return nil
}

// InvoiceExists checks if the Invoice row exists.
func InvoiceExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"invoices\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if invoices exists")
	}

	return exists, nil
}

// Exists checks if the Invoice row exists.
func (o *Invoice) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return InvoiceExists(ctx, exec, o.ID)
}

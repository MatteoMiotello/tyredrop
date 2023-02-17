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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ProductLanguage is an object representing the database table.
type ProductLanguage struct {
	ID          int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProductID   int64       `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	LanguageID  int64       `boil:"language_id" json:"language_id" toml:"language_id" yaml:"language_id"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	UpdatedAt   time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *productLanguageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productLanguageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductLanguageColumns = struct {
	ID          string
	ProductID   string
	LanguageID  string
	Name        string
	Description string
	UpdatedAt   string
	CreatedAt   string
}{
	ID:          "id",
	ProductID:   "product_id",
	LanguageID:  "language_id",
	Name:        "name",
	Description: "description",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
}

var ProductLanguageTableColumns = struct {
	ID          string
	ProductID   string
	LanguageID  string
	Name        string
	Description string
	UpdatedAt   string
	CreatedAt   string
}{
	ID:          "product_languages.id",
	ProductID:   "product_languages.product_id",
	LanguageID:  "product_languages.language_id",
	Name:        "product_languages.name",
	Description: "product_languages.description",
	UpdatedAt:   "product_languages.updated_at",
	CreatedAt:   "product_languages.created_at",
}

// Generated where

var ProductLanguageWhere = struct {
	ID          whereHelperint64
	ProductID   whereHelperint64
	LanguageID  whereHelperint64
	Name        whereHelperstring
	Description whereHelpernull_String
	UpdatedAt   whereHelpertime_Time
	CreatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint64{field: "\"product_languages\".\"id\""},
	ProductID:   whereHelperint64{field: "\"product_languages\".\"product_id\""},
	LanguageID:  whereHelperint64{field: "\"product_languages\".\"language_id\""},
	Name:        whereHelperstring{field: "\"product_languages\".\"name\""},
	Description: whereHelpernull_String{field: "\"product_languages\".\"description\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"product_languages\".\"updated_at\""},
	CreatedAt:   whereHelpertime_Time{field: "\"product_languages\".\"created_at\""},
}

// ProductLanguageRels is where relationship names are stored.
var ProductLanguageRels = struct {
	Language string
	Product  string
}{
	Language: "Language",
	Product:  "Product",
}

// productLanguageR is where relationships are stored.
type productLanguageR struct {
	Language *Language `boil:"Language" json:"Language" toml:"Language" yaml:"Language"`
	Product  *Product  `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
}

// NewStruct creates a new relationship struct
func (*productLanguageR) NewStruct() *productLanguageR {
	return &productLanguageR{}
}

func (r *productLanguageR) GetLanguage() *Language {
	if r == nil {
		return nil
	}
	return r.Language
}

func (r *productLanguageR) GetProduct() *Product {
	if r == nil {
		return nil
	}
	return r.Product
}

// productLanguageL is where Load methods for each relationship are stored.
type productLanguageL struct{}

var (
	productLanguageAllColumns            = []string{"id", "product_id", "language_id", "name", "description", "updated_at", "created_at"}
	productLanguageColumnsWithoutDefault = []string{"name"}
	productLanguageColumnsWithDefault    = []string{"id", "product_id", "language_id", "description", "updated_at", "created_at"}
	productLanguagePrimaryKeyColumns     = []string{"id"}
	productLanguageGeneratedColumns      = []string{}
)

type (
	// ProductLanguageSlice is an alias for a slice of pointers to ProductLanguage.
	// This should almost always be used instead of []ProductLanguage.
	ProductLanguageSlice []*ProductLanguage
	// ProductLanguageHook is the signature for custom ProductLanguage hook methods
	ProductLanguageHook func(context.Context, boil.ContextExecutor, *ProductLanguage) error

	productLanguageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productLanguageType                 = reflect.TypeOf(&ProductLanguage{})
	productLanguageMapping              = queries.MakeStructMapping(productLanguageType)
	productLanguagePrimaryKeyMapping, _ = queries.BindMapping(productLanguageType, productLanguageMapping, productLanguagePrimaryKeyColumns)
	productLanguageInsertCacheMut       sync.RWMutex
	productLanguageInsertCache          = make(map[string]insertCache)
	productLanguageUpdateCacheMut       sync.RWMutex
	productLanguageUpdateCache          = make(map[string]updateCache)
	productLanguageUpsertCacheMut       sync.RWMutex
	productLanguageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productLanguageAfterSelectHooks []ProductLanguageHook

var productLanguageBeforeInsertHooks []ProductLanguageHook
var productLanguageAfterInsertHooks []ProductLanguageHook

var productLanguageBeforeUpdateHooks []ProductLanguageHook
var productLanguageAfterUpdateHooks []ProductLanguageHook

var productLanguageBeforeDeleteHooks []ProductLanguageHook
var productLanguageAfterDeleteHooks []ProductLanguageHook

var productLanguageBeforeUpsertHooks []ProductLanguageHook
var productLanguageAfterUpsertHooks []ProductLanguageHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductLanguage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductLanguage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductLanguage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductLanguage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductLanguage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductLanguage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductLanguage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductLanguage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductLanguage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productLanguageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductLanguageHook registers your hook function for all future operations.
func AddProductLanguageHook(hookPoint boil.HookPoint, productLanguageHook ProductLanguageHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productLanguageAfterSelectHooks = append(productLanguageAfterSelectHooks, productLanguageHook)
	case boil.BeforeInsertHook:
		productLanguageBeforeInsertHooks = append(productLanguageBeforeInsertHooks, productLanguageHook)
	case boil.AfterInsertHook:
		productLanguageAfterInsertHooks = append(productLanguageAfterInsertHooks, productLanguageHook)
	case boil.BeforeUpdateHook:
		productLanguageBeforeUpdateHooks = append(productLanguageBeforeUpdateHooks, productLanguageHook)
	case boil.AfterUpdateHook:
		productLanguageAfterUpdateHooks = append(productLanguageAfterUpdateHooks, productLanguageHook)
	case boil.BeforeDeleteHook:
		productLanguageBeforeDeleteHooks = append(productLanguageBeforeDeleteHooks, productLanguageHook)
	case boil.AfterDeleteHook:
		productLanguageAfterDeleteHooks = append(productLanguageAfterDeleteHooks, productLanguageHook)
	case boil.BeforeUpsertHook:
		productLanguageBeforeUpsertHooks = append(productLanguageBeforeUpsertHooks, productLanguageHook)
	case boil.AfterUpsertHook:
		productLanguageAfterUpsertHooks = append(productLanguageAfterUpsertHooks, productLanguageHook)
	}
}

// One returns a single productLanguage record from the query.
func (q productLanguageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProductLanguage, error) {
	o := &ProductLanguage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for product_languages")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProductLanguage records from the query.
func (q productLanguageQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductLanguageSlice, error) {
	var o []*ProductLanguage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProductLanguage slice")
	}

	if len(productLanguageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProductLanguage records in the query.
func (q productLanguageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count product_languages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q productLanguageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if product_languages exists")
	}

	return count > 0, nil
}

// Language pointed to by the foreign key.
func (o *ProductLanguage) Language(mods ...qm.QueryMod) languageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LanguageID),
	}

	queryMods = append(queryMods, mods...)

	return Languages(queryMods...)
}

// Product pointed to by the foreign key.
func (o *ProductLanguage) Product(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	return Products(queryMods...)
}

// LoadLanguage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (productLanguageL) LoadLanguage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProductLanguage interface{}, mods queries.Applicator) error {
	var slice []*ProductLanguage
	var object *ProductLanguage

	if singular {
		var ok bool
		object, ok = maybeProductLanguage.(*ProductLanguage)
		if !ok {
			object = new(ProductLanguage)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProductLanguage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProductLanguage))
			}
		}
	} else {
		s, ok := maybeProductLanguage.(*[]*ProductLanguage)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProductLanguage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProductLanguage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &productLanguageR{}
		}
		args = append(args, object.LanguageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &productLanguageR{}
			}

			for _, a := range args {
				if a == obj.LanguageID {
					continue Outer
				}
			}

			args = append(args, obj.LanguageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`languages`),
		qm.WhereIn(`languages.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Language")
	}

	var resultSlice []*Language
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Language")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for languages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for languages")
	}

	if len(languageAfterSelectHooks) != 0 {
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
		object.R.Language = foreign
		if foreign.R == nil {
			foreign.R = &languageR{}
		}
		foreign.R.ProductLanguages = append(foreign.R.ProductLanguages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LanguageID == foreign.ID {
				local.R.Language = foreign
				if foreign.R == nil {
					foreign.R = &languageR{}
				}
				foreign.R.ProductLanguages = append(foreign.R.ProductLanguages, local)
				break
			}
		}
	}

	return nil
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (productLanguageL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProductLanguage interface{}, mods queries.Applicator) error {
	var slice []*ProductLanguage
	var object *ProductLanguage

	if singular {
		var ok bool
		object, ok = maybeProductLanguage.(*ProductLanguage)
		if !ok {
			object = new(ProductLanguage)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProductLanguage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProductLanguage))
			}
		}
	} else {
		s, ok := maybeProductLanguage.(*[]*ProductLanguage)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProductLanguage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProductLanguage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &productLanguageR{}
		}
		args = append(args, object.ProductID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &productLanguageR{}
			}

			for _, a := range args {
				if a == obj.ProductID {
					continue Outer
				}
			}

			args = append(args, obj.ProductID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`products`),
		qm.WhereIn(`products.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Product")
	}

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Product")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for products")
	}

	if len(productAfterSelectHooks) != 0 {
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
		object.R.Product = foreign
		if foreign.R == nil {
			foreign.R = &productR{}
		}
		foreign.R.ProductLanguages = append(foreign.R.ProductLanguages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductID == foreign.ID {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.ProductLanguages = append(foreign.R.ProductLanguages, local)
				break
			}
		}
	}

	return nil
}

// SetLanguage of the productLanguage to the related item.
// Sets o.R.Language to related.
// Adds o to related.R.ProductLanguages.
func (o *ProductLanguage) SetLanguage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Language) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"product_languages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"language_id"}),
		strmangle.WhereClause("\"", "\"", 2, productLanguagePrimaryKeyColumns),
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

	o.LanguageID = related.ID
	if o.R == nil {
		o.R = &productLanguageR{
			Language: related,
		}
	} else {
		o.R.Language = related
	}

	if related.R == nil {
		related.R = &languageR{
			ProductLanguages: ProductLanguageSlice{o},
		}
	} else {
		related.R.ProductLanguages = append(related.R.ProductLanguages, o)
	}

	return nil
}

// SetProduct of the productLanguage to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductLanguages.
func (o *ProductLanguage) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"product_languages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"product_id"}),
		strmangle.WhereClause("\"", "\"", 2, productLanguagePrimaryKeyColumns),
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

	o.ProductID = related.ID
	if o.R == nil {
		o.R = &productLanguageR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			ProductLanguages: ProductLanguageSlice{o},
		}
	} else {
		related.R.ProductLanguages = append(related.R.ProductLanguages, o)
	}

	return nil
}

// ProductLanguages retrieves all the records using an executor.
func ProductLanguages(mods ...qm.QueryMod) productLanguageQuery {
	mods = append(mods, qm.From("\"product_languages\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"product_languages\".*"})
	}

	return productLanguageQuery{q}
}

// FindProductLanguage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductLanguage(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ProductLanguage, error) {
	productLanguageObj := &ProductLanguage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"product_languages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productLanguageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from product_languages")
	}

	if err = productLanguageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productLanguageObj, err
	}

	return productLanguageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductLanguage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no product_languages provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(productLanguageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productLanguageInsertCacheMut.RLock()
	cache, cached := productLanguageInsertCache[key]
	productLanguageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productLanguageAllColumns,
			productLanguageColumnsWithDefault,
			productLanguageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(productLanguageType, productLanguageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productLanguageType, productLanguageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"product_languages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"product_languages\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into product_languages")
	}

	if !cached {
		productLanguageInsertCacheMut.Lock()
		productLanguageInsertCache[key] = cache
		productLanguageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProductLanguage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductLanguage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productLanguageUpdateCacheMut.RLock()
	cache, cached := productLanguageUpdateCache[key]
	productLanguageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productLanguageAllColumns,
			productLanguagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update product_languages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"product_languages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, productLanguagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productLanguageType, productLanguageMapping, append(wl, productLanguagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update product_languages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for product_languages")
	}

	if !cached {
		productLanguageUpdateCacheMut.Lock()
		productLanguageUpdateCache[key] = cache
		productLanguageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q productLanguageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for product_languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for product_languages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductLanguageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productLanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"product_languages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productLanguagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in productLanguage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all productLanguage")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProductLanguage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no product_languages provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(productLanguageColumnsWithDefault, o)

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

	productLanguageUpsertCacheMut.RLock()
	cache, cached := productLanguageUpsertCache[key]
	productLanguageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			productLanguageAllColumns,
			productLanguageColumnsWithDefault,
			productLanguageColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			productLanguageAllColumns,
			productLanguagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert product_languages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(productLanguagePrimaryKeyColumns))
			copy(conflict, productLanguagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"product_languages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(productLanguageType, productLanguageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productLanguageType, productLanguageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert product_languages")
	}

	if !cached {
		productLanguageUpsertCacheMut.Lock()
		productLanguageUpsertCache[key] = cache
		productLanguageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProductLanguage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductLanguage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProductLanguage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productLanguagePrimaryKeyMapping)
	sql := "DELETE FROM \"product_languages\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from product_languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for product_languages")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q productLanguageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no productLanguageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from product_languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for product_languages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductLanguageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productLanguageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productLanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"product_languages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productLanguagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from productLanguage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for product_languages")
	}

	if len(productLanguageAfterDeleteHooks) != 0 {
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
func (o *ProductLanguage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductLanguage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductLanguageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductLanguageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productLanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"product_languages\".* FROM \"product_languages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productLanguagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProductLanguageSlice")
	}

	*o = slice

	return nil
}

// ProductLanguageExists checks if the ProductLanguage row exists.
func ProductLanguageExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"product_languages\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if product_languages exists")
	}

	return exists, nil
}

// Exists checks if the ProductLanguage row exists.
func (o *ProductLanguage) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductLanguageExists(ctx, exec, o.ID)
}

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

// UserAddress is an object representing the database table.
type UserAddress struct {
	ID           int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID       int64       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	AddressName  string      `boil:"address_name" json:"address_name" toml:"address_name" yaml:"address_name"`
	AddressLine1 string      `boil:"address_line_1" json:"address_line_1" toml:"address_line_1" yaml:"address_line_1"`
	AddressLine2 null.String `boil:"address_line_2" json:"address_line_2,omitempty" toml:"address_line_2" yaml:"address_line_2,omitempty"`
	City         string      `boil:"city" json:"city" toml:"city" yaml:"city"`
	Province     string      `boil:"province" json:"province" toml:"province" yaml:"province"`
	PostalCode   string      `boil:"postal_code" json:"postal_code" toml:"postal_code" yaml:"postal_code"`
	Country      string      `boil:"country" json:"country" toml:"country" yaml:"country"`
	IsDefault    bool        `boil:"is_default" json:"is_default" toml:"is_default" yaml:"is_default"`
	DeletedAt    null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	UpdatedAt    time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *userAddressR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userAddressL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserAddressColumns = struct {
	ID           string
	UserID       string
	AddressName  string
	AddressLine1 string
	AddressLine2 string
	City         string
	Province     string
	PostalCode   string
	Country      string
	IsDefault    string
	DeletedAt    string
	UpdatedAt    string
	CreatedAt    string
}{
	ID:           "id",
	UserID:       "user_id",
	AddressName:  "address_name",
	AddressLine1: "address_line_1",
	AddressLine2: "address_line_2",
	City:         "city",
	Province:     "province",
	PostalCode:   "postal_code",
	Country:      "country",
	IsDefault:    "is_default",
	DeletedAt:    "deleted_at",
	UpdatedAt:    "updated_at",
	CreatedAt:    "created_at",
}

var UserAddressTableColumns = struct {
	ID           string
	UserID       string
	AddressName  string
	AddressLine1 string
	AddressLine2 string
	City         string
	Province     string
	PostalCode   string
	Country      string
	IsDefault    string
	DeletedAt    string
	UpdatedAt    string
	CreatedAt    string
}{
	ID:           "user_address.id",
	UserID:       "user_address.user_id",
	AddressName:  "user_address.address_name",
	AddressLine1: "user_address.address_line_1",
	AddressLine2: "user_address.address_line_2",
	City:         "user_address.city",
	Province:     "user_address.province",
	PostalCode:   "user_address.postal_code",
	Country:      "user_address.country",
	IsDefault:    "user_address.is_default",
	DeletedAt:    "user_address.deleted_at",
	UpdatedAt:    "user_address.updated_at",
	CreatedAt:    "user_address.created_at",
}

// Generated where

var UserAddressWhere = struct {
	ID           whereHelperint64
	UserID       whereHelperint64
	AddressName  whereHelperstring
	AddressLine1 whereHelperstring
	AddressLine2 whereHelpernull_String
	City         whereHelperstring
	Province     whereHelperstring
	PostalCode   whereHelperstring
	Country      whereHelperstring
	IsDefault    whereHelperbool
	DeletedAt    whereHelpernull_Time
	UpdatedAt    whereHelpertime_Time
	CreatedAt    whereHelpertime_Time
}{
	ID:           whereHelperint64{field: "\"user_address\".\"id\""},
	UserID:       whereHelperint64{field: "\"user_address\".\"user_id\""},
	AddressName:  whereHelperstring{field: "\"user_address\".\"address_name\""},
	AddressLine1: whereHelperstring{field: "\"user_address\".\"address_line_1\""},
	AddressLine2: whereHelpernull_String{field: "\"user_address\".\"address_line_2\""},
	City:         whereHelperstring{field: "\"user_address\".\"city\""},
	Province:     whereHelperstring{field: "\"user_address\".\"province\""},
	PostalCode:   whereHelperstring{field: "\"user_address\".\"postal_code\""},
	Country:      whereHelperstring{field: "\"user_address\".\"country\""},
	IsDefault:    whereHelperbool{field: "\"user_address\".\"is_default\""},
	DeletedAt:    whereHelpernull_Time{field: "\"user_address\".\"deleted_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"user_address\".\"updated_at\""},
	CreatedAt:    whereHelpertime_Time{field: "\"user_address\".\"created_at\""},
}

// UserAddressRels is where relationship names are stored.
var UserAddressRels = struct {
	User string
}{
	User: "User",
}

// userAddressR is where relationships are stored.
type userAddressR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userAddressR) NewStruct() *userAddressR {
	return &userAddressR{}
}

func (r *userAddressR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userAddressL is where Load methods for each relationship are stored.
type userAddressL struct{}

var (
	userAddressAllColumns            = []string{"id", "user_id", "address_name", "address_line_1", "address_line_2", "city", "province", "postal_code", "country", "is_default", "deleted_at", "updated_at", "created_at"}
	userAddressColumnsWithoutDefault = []string{"user_id", "address_name", "address_line_1", "city", "province", "postal_code", "country", "is_default"}
	userAddressColumnsWithDefault    = []string{"id", "address_line_2", "deleted_at", "updated_at", "created_at"}
	userAddressPrimaryKeyColumns     = []string{"id"}
	userAddressGeneratedColumns      = []string{}
)

type (
	// UserAddressSlice is an alias for a slice of pointers to UserAddress.
	// This should almost always be used instead of []UserAddress.
	UserAddressSlice []*UserAddress
	// UserAddressHook is the signature for custom UserAddress hook methods
	UserAddressHook func(context.Context, boil.ContextExecutor, *UserAddress) error

	userAddressQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userAddressType                 = reflect.TypeOf(&UserAddress{})
	userAddressMapping              = queries.MakeStructMapping(userAddressType)
	userAddressPrimaryKeyMapping, _ = queries.BindMapping(userAddressType, userAddressMapping, userAddressPrimaryKeyColumns)
	userAddressInsertCacheMut       sync.RWMutex
	userAddressInsertCache          = make(map[string]insertCache)
	userAddressUpdateCacheMut       sync.RWMutex
	userAddressUpdateCache          = make(map[string]updateCache)
	userAddressUpsertCacheMut       sync.RWMutex
	userAddressUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userAddressAfterSelectHooks []UserAddressHook

var userAddressBeforeInsertHooks []UserAddressHook
var userAddressAfterInsertHooks []UserAddressHook

var userAddressBeforeUpdateHooks []UserAddressHook
var userAddressAfterUpdateHooks []UserAddressHook

var userAddressBeforeDeleteHooks []UserAddressHook
var userAddressAfterDeleteHooks []UserAddressHook

var userAddressBeforeUpsertHooks []UserAddressHook
var userAddressAfterUpsertHooks []UserAddressHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserAddress) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserAddress) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserAddress) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserAddress) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserAddress) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserAddress) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserAddress) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserAddress) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserAddress) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserAddressHook registers your hook function for all future operations.
func AddUserAddressHook(hookPoint boil.HookPoint, userAddressHook UserAddressHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userAddressAfterSelectHooks = append(userAddressAfterSelectHooks, userAddressHook)
	case boil.BeforeInsertHook:
		userAddressBeforeInsertHooks = append(userAddressBeforeInsertHooks, userAddressHook)
	case boil.AfterInsertHook:
		userAddressAfterInsertHooks = append(userAddressAfterInsertHooks, userAddressHook)
	case boil.BeforeUpdateHook:
		userAddressBeforeUpdateHooks = append(userAddressBeforeUpdateHooks, userAddressHook)
	case boil.AfterUpdateHook:
		userAddressAfterUpdateHooks = append(userAddressAfterUpdateHooks, userAddressHook)
	case boil.BeforeDeleteHook:
		userAddressBeforeDeleteHooks = append(userAddressBeforeDeleteHooks, userAddressHook)
	case boil.AfterDeleteHook:
		userAddressAfterDeleteHooks = append(userAddressAfterDeleteHooks, userAddressHook)
	case boil.BeforeUpsertHook:
		userAddressBeforeUpsertHooks = append(userAddressBeforeUpsertHooks, userAddressHook)
	case boil.AfterUpsertHook:
		userAddressAfterUpsertHooks = append(userAddressAfterUpsertHooks, userAddressHook)
	}
}

// One returns a single userAddress record from the query.
func (q userAddressQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserAddress, error) {
	o := &UserAddress{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_address")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserAddress records from the query.
func (q userAddressQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserAddressSlice, error) {
	var o []*UserAddress

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserAddress slice")
	}

	if len(userAddressAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserAddress records in the query.
func (q userAddressQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_address rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userAddressQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_address exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserAddress) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userAddressL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserAddress interface{}, mods queries.Applicator) error {
	var slice []*UserAddress
	var object *UserAddress

	if singular {
		var ok bool
		object, ok = maybeUserAddress.(*UserAddress)
		if !ok {
			object = new(UserAddress)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserAddress)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserAddress))
			}
		}
	} else {
		s, ok := maybeUserAddress.(*[]*UserAddress)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserAddress)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserAddress))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userAddressR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userAddressR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
		qmhelper.WhereIsNull(`users.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserAddresses = append(foreign.R.UserAddresses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserAddresses = append(foreign.R.UserAddresses, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the userAddress to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAddresses.
func (o *UserAddress) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_address\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userAddressPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userAddressR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserAddresses: UserAddressSlice{o},
		}
	} else {
		related.R.UserAddresses = append(related.R.UserAddresses, o)
	}

	return nil
}

// UserAddresses retrieves all the records using an executor.
func UserAddresses(mods ...qm.QueryMod) userAddressQuery {
	mods = append(mods, qm.From("\"user_address\""), qmhelper.WhereIsNull("\"user_address\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_address\".*"})
	}

	return userAddressQuery{q}
}

// FindUserAddress retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserAddress(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*UserAddress, error) {
	userAddressObj := &UserAddress{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_address\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userAddressObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_address")
	}

	if err = userAddressObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userAddressObj, err
	}

	return userAddressObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserAddress) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_address provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userAddressColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userAddressInsertCacheMut.RLock()
	cache, cached := userAddressInsertCache[key]
	userAddressInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userAddressAllColumns,
			userAddressColumnsWithDefault,
			userAddressColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userAddressType, userAddressMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userAddressType, userAddressMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_address\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_address\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_address")
	}

	if !cached {
		userAddressInsertCacheMut.Lock()
		userAddressInsertCache[key] = cache
		userAddressInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserAddress.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserAddress) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userAddressUpdateCacheMut.RLock()
	cache, cached := userAddressUpdateCache[key]
	userAddressUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userAddressAllColumns,
			userAddressPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_address, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_address\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userAddressPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userAddressType, userAddressMapping, append(wl, userAddressPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_address row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_address")
	}

	if !cached {
		userAddressUpdateCacheMut.Lock()
		userAddressUpdateCache[key] = cache
		userAddressUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userAddressQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_address")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserAddressSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_address\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userAddressPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userAddress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userAddress")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserAddress) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_address provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(userAddressColumnsWithDefault, o)

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

	userAddressUpsertCacheMut.RLock()
	cache, cached := userAddressUpsertCache[key]
	userAddressUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userAddressAllColumns,
			userAddressColumnsWithDefault,
			userAddressColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userAddressAllColumns,
			userAddressPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_address, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userAddressPrimaryKeyColumns))
			copy(conflict, userAddressPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_address\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userAddressType, userAddressMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userAddressType, userAddressMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_address")
	}

	if !cached {
		userAddressUpsertCacheMut.Lock()
		userAddressUpsertCache[key] = cache
		userAddressUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserAddress record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserAddress) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserAddress provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userAddressPrimaryKeyMapping)
		sql = "DELETE FROM \"user_address\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"user_address\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(userAddressType, userAddressMapping, append(wl, userAddressPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to delete from user_address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_address")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userAddressQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userAddressQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_address")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserAddressSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userAddressBeforeDeleteHooks) != 0 {
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
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAddressPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"user_address\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userAddressPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAddressPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"user_address\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, userAddressPrimaryKeyColumns, len(o)),
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
		return 0, errors.Wrap(err, "models: unable to delete all from userAddress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_address")
	}

	if len(userAddressAfterDeleteHooks) != 0 {
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
func (o *UserAddress) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserAddress(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserAddressSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserAddressSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_address\".* FROM \"user_address\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userAddressPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserAddressSlice")
	}

	*o = slice

	return nil
}

// UserAddressExists checks if the UserAddress row exists.
func UserAddressExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_address\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_address exists")
	}

	return exists, nil
}

// Exists checks if the UserAddress row exists.
func (o *UserAddress) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserAddressExists(ctx, exec, o.ID)
}

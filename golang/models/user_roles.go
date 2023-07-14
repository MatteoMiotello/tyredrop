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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserRole is an object representing the database table.
type UserRole struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	RoleCode  string    `boil:"role_code" json:"role_code" toml:"role_code" yaml:"role_code"`
	Admin     bool      `boil:"admin" json:"admin" toml:"admin" yaml:"admin"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *userRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserRoleColumns = struct {
	ID        string
	RoleCode  string
	Admin     string
	CreatedAt string
}{
	ID:        "id",
	RoleCode:  "role_code",
	Admin:     "admin",
	CreatedAt: "created_at",
}

var UserRoleTableColumns = struct {
	ID        string
	RoleCode  string
	Admin     string
	CreatedAt string
}{
	ID:        "user_roles.id",
	RoleCode:  "user_roles.role_code",
	Admin:     "user_roles.admin",
	CreatedAt: "user_roles.created_at",
}

// Generated where

var UserRoleWhere = struct {
	ID        whereHelperint64
	RoleCode  whereHelperstring
	Admin     whereHelperbool
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"user_roles\".\"id\""},
	RoleCode:  whereHelperstring{field: "\"user_roles\".\"role_code\""},
	Admin:     whereHelperbool{field: "\"user_roles\".\"admin\""},
	CreatedAt: whereHelpertime_Time{field: "\"user_roles\".\"created_at\""},
}

// UserRoleRels is where relationship names are stored.
var UserRoleRels = struct {
	UserRoleLanguages string
	Users             string
}{
	UserRoleLanguages: "UserRoleLanguages",
	Users:             "Users",
}

// userRoleR is where relationships are stored.
type userRoleR struct {
	UserRoleLanguages UserRoleLanguageSlice `boil:"UserRoleLanguages" json:"UserRoleLanguages" toml:"UserRoleLanguages" yaml:"UserRoleLanguages"`
	Users             UserSlice             `boil:"Users" json:"Users" toml:"Users" yaml:"Users"`
}

// NewStruct creates a new relationship struct
func (*userRoleR) NewStruct() *userRoleR {
	return &userRoleR{}
}

func (r *userRoleR) GetUserRoleLanguages() UserRoleLanguageSlice {
	if r == nil {
		return nil
	}
	return r.UserRoleLanguages
}

func (r *userRoleR) GetUsers() UserSlice {
	if r == nil {
		return nil
	}
	return r.Users
}

// userRoleL is where Load methods for each relationship are stored.
type userRoleL struct{}

var (
	userRoleAllColumns            = []string{"id", "role_code", "admin", "created_at"}
	userRoleColumnsWithoutDefault = []string{"role_code", "admin"}
	userRoleColumnsWithDefault    = []string{"id", "created_at"}
	userRolePrimaryKeyColumns     = []string{"id"}
	userRoleGeneratedColumns      = []string{}
)

type (
	// UserRoleSlice is an alias for a slice of pointers to UserRole.
	// This should almost always be used instead of []UserRole.
	UserRoleSlice []*UserRole
	// UserRoleHook is the signature for custom UserRole hook methods
	UserRoleHook func(context.Context, boil.ContextExecutor, *UserRole) error

	userRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userRoleType                 = reflect.TypeOf(&UserRole{})
	userRoleMapping              = queries.MakeStructMapping(userRoleType)
	userRolePrimaryKeyMapping, _ = queries.BindMapping(userRoleType, userRoleMapping, userRolePrimaryKeyColumns)
	userRoleInsertCacheMut       sync.RWMutex
	userRoleInsertCache          = make(map[string]insertCache)
	userRoleUpdateCacheMut       sync.RWMutex
	userRoleUpdateCache          = make(map[string]updateCache)
	userRoleUpsertCacheMut       sync.RWMutex
	userRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userRoleAfterSelectHooks []UserRoleHook

var userRoleBeforeInsertHooks []UserRoleHook
var userRoleAfterInsertHooks []UserRoleHook

var userRoleBeforeUpdateHooks []UserRoleHook
var userRoleAfterUpdateHooks []UserRoleHook

var userRoleBeforeDeleteHooks []UserRoleHook
var userRoleAfterDeleteHooks []UserRoleHook

var userRoleBeforeUpsertHooks []UserRoleHook
var userRoleAfterUpsertHooks []UserRoleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserRole) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserRole) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserRole) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserRole) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserRole) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserRole) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserRole) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserRole) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserRole) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRoleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserRoleHook registers your hook function for all future operations.
func AddUserRoleHook(hookPoint boil.HookPoint, userRoleHook UserRoleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userRoleAfterSelectHooks = append(userRoleAfterSelectHooks, userRoleHook)
	case boil.BeforeInsertHook:
		userRoleBeforeInsertHooks = append(userRoleBeforeInsertHooks, userRoleHook)
	case boil.AfterInsertHook:
		userRoleAfterInsertHooks = append(userRoleAfterInsertHooks, userRoleHook)
	case boil.BeforeUpdateHook:
		userRoleBeforeUpdateHooks = append(userRoleBeforeUpdateHooks, userRoleHook)
	case boil.AfterUpdateHook:
		userRoleAfterUpdateHooks = append(userRoleAfterUpdateHooks, userRoleHook)
	case boil.BeforeDeleteHook:
		userRoleBeforeDeleteHooks = append(userRoleBeforeDeleteHooks, userRoleHook)
	case boil.AfterDeleteHook:
		userRoleAfterDeleteHooks = append(userRoleAfterDeleteHooks, userRoleHook)
	case boil.BeforeUpsertHook:
		userRoleBeforeUpsertHooks = append(userRoleBeforeUpsertHooks, userRoleHook)
	case boil.AfterUpsertHook:
		userRoleAfterUpsertHooks = append(userRoleAfterUpsertHooks, userRoleHook)
	}
}

// One returns a single userRole record from the query.
func (q userRoleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserRole, error) {
	o := &UserRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_roles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserRole records from the query.
func (q userRoleQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserRoleSlice, error) {
	var o []*UserRole

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserRole slice")
	}

	if len(userRoleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserRole records in the query.
func (q userRoleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_roles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userRoleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_roles exists")
	}

	return count > 0, nil
}

// UserRoleLanguages retrieves all the user_role_language's UserRoleLanguages with an executor.
func (o *UserRole) UserRoleLanguages(mods ...qm.QueryMod) userRoleLanguageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_role_languages\".\"user_role_id\"=?", o.ID),
	)

	return UserRoleLanguages(queryMods...)
}

// Users retrieves all the user's Users with an executor.
func (o *UserRole) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"users\".\"user_role_id\"=?", o.ID),
	)

	return Users(queryMods...)
}

// LoadUserRoleLanguages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (userRoleL) LoadUserRoleLanguages(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserRole interface{}, mods queries.Applicator) error {
	var slice []*UserRole
	var object *UserRole

	if singular {
		var ok bool
		object, ok = maybeUserRole.(*UserRole)
		if !ok {
			object = new(UserRole)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserRole)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserRole))
			}
		}
	} else {
		s, ok := maybeUserRole.(*[]*UserRole)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserRole)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserRole))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userRoleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userRoleR{}
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
		qm.From(`user_role_languages`),
		qm.WhereIn(`user_role_languages.user_role_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_role_languages")
	}

	var resultSlice []*UserRoleLanguage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_role_languages")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_role_languages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_role_languages")
	}

	if len(userRoleLanguageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserRoleLanguages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userRoleLanguageR{}
			}
			foreign.R.UserRole = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.UserRoleID {
				local.R.UserRoleLanguages = append(local.R.UserRoleLanguages, foreign)
				if foreign.R == nil {
					foreign.R = &userRoleLanguageR{}
				}
				foreign.R.UserRole = local
				break
			}
		}
	}

	return nil
}

// LoadUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (userRoleL) LoadUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserRole interface{}, mods queries.Applicator) error {
	var slice []*UserRole
	var object *UserRole

	if singular {
		var ok bool
		object, ok = maybeUserRole.(*UserRole)
		if !ok {
			object = new(UserRole)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserRole)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserRole))
			}
		}
	} else {
		s, ok := maybeUserRole.(*[]*UserRole)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserRole)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserRole))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userRoleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userRoleR{}
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
		qm.From(`users`),
		qm.WhereIn(`users.user_role_id in ?`, args...),
		qmhelper.WhereIsNull(`users.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users")
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
	if singular {
		object.R.Users = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userR{}
			}
			foreign.R.UserRole = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.UserRoleID {
				local.R.Users = append(local.R.Users, foreign)
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserRole = local
				break
			}
		}
	}

	return nil
}

// AddUserRoleLanguages adds the given related objects to the existing relationships
// of the user_role, optionally inserting them as new records.
// Appends related to o.R.UserRoleLanguages.
// Sets related.R.UserRole appropriately.
func (o *UserRole) AddUserRoleLanguages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserRoleLanguage) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserRoleID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_role_languages\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"user_role_id"}),
				strmangle.WhereClause("\"", "\"", 2, userRoleLanguagePrimaryKeyColumns),
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

			rel.UserRoleID = o.ID
		}
	}

	if o.R == nil {
		o.R = &userRoleR{
			UserRoleLanguages: related,
		}
	} else {
		o.R.UserRoleLanguages = append(o.R.UserRoleLanguages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userRoleLanguageR{
				UserRole: o,
			}
		} else {
			rel.R.UserRole = o
		}
	}
	return nil
}

// AddUsers adds the given related objects to the existing relationships
// of the user_role, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.UserRole appropriately.
func (o *UserRole) AddUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserRoleID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"users\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"user_role_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPrimaryKeyColumns),
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

			rel.UserRoleID = o.ID
		}
	}

	if o.R == nil {
		o.R = &userRoleR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				UserRole: o,
			}
		} else {
			rel.R.UserRole = o
		}
	}
	return nil
}

// UserRoles retrieves all the records using an executor.
func UserRoles(mods ...qm.QueryMod) userRoleQuery {
	mods = append(mods, qm.From("\"user_roles\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_roles\".*"})
	}

	return userRoleQuery{q}
}

// FindUserRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserRole(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*UserRole, error) {
	userRoleObj := &UserRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_roles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userRoleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_roles")
	}

	if err = userRoleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userRoleObj, err
	}

	return userRoleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserRole) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_roles provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userRoleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userRoleInsertCacheMut.RLock()
	cache, cached := userRoleInsertCache[key]
	userRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userRoleAllColumns,
			userRoleColumnsWithDefault,
			userRoleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userRoleType, userRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userRoleType, userRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_roles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_roles\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_roles")
	}

	if !cached {
		userRoleInsertCacheMut.Lock()
		userRoleInsertCache[key] = cache
		userRoleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserRole.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserRole) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userRoleUpdateCacheMut.RLock()
	cache, cached := userRoleUpdateCache[key]
	userRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userRoleAllColumns,
			userRolePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_roles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userRoleType, userRoleMapping, append(wl, userRolePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_roles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_roles")
	}

	if !cached {
		userRoleUpdateCacheMut.Lock()
		userRoleUpdateCache[key] = cache
		userRoleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userRoleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_roles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserRoleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userRolePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userRole slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userRole")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserRole) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_roles provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(userRoleColumnsWithDefault, o)

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

	userRoleUpsertCacheMut.RLock()
	cache, cached := userRoleUpsertCache[key]
	userRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userRoleAllColumns,
			userRoleColumnsWithDefault,
			userRoleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userRoleAllColumns,
			userRolePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_roles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userRolePrimaryKeyColumns))
			copy(conflict, userRolePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_roles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userRoleType, userRoleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userRoleType, userRoleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_roles")
	}

	if !cached {
		userRoleUpsertCacheMut.Lock()
		userRoleUpsertCache[key] = cache
		userRoleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserRole) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserRole provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userRolePrimaryKeyMapping)
	sql := "DELETE FROM \"user_roles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_roles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userRoleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_roles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserRoleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userRoleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userRolePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userRole slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_roles")
	}

	if len(userRoleAfterDeleteHooks) != 0 {
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
func (o *UserRole) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserRole(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserRoleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_roles\".* FROM \"user_roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userRolePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserRoleSlice")
	}

	*o = slice

	return nil
}

// UserRoleExists checks if the UserRole row exists.
func UserRoleExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_roles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_roles exists")
	}

	return exists, nil
}

// Exists checks if the UserRole row exists.
func (o *UserRole) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserRoleExists(ctx, exec, o.ID)
}

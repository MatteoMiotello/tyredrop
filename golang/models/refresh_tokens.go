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

// RefreshToken is an object representing the database table.
type RefreshToken struct {
	ID           int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID       int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	RefreshToken string    `boil:"refresh_token" json:"refresh_token" toml:"refresh_token" yaml:"refresh_token"`
	ExpiresAt    time.Time `boil:"expires_at" json:"expires_at" toml:"expires_at" yaml:"expires_at"`
	TimeLastUse  null.Time `boil:"time_last_use" json:"time_last_use,omitempty" toml:"time_last_use" yaml:"time_last_use,omitempty"`
	DeletedAt    null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *refreshTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L refreshTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RefreshTokenColumns = struct {
	ID           string
	UserID       string
	RefreshToken string
	ExpiresAt    string
	TimeLastUse  string
	DeletedAt    string
	CreatedAt    string
}{
	ID:           "id",
	UserID:       "user_id",
	RefreshToken: "refresh_token",
	ExpiresAt:    "expires_at",
	TimeLastUse:  "time_last_use",
	DeletedAt:    "deleted_at",
	CreatedAt:    "created_at",
}

var RefreshTokenTableColumns = struct {
	ID           string
	UserID       string
	RefreshToken string
	ExpiresAt    string
	TimeLastUse  string
	DeletedAt    string
	CreatedAt    string
}{
	ID:           "refresh_tokens.id",
	UserID:       "refresh_tokens.user_id",
	RefreshToken: "refresh_tokens.refresh_token",
	ExpiresAt:    "refresh_tokens.expires_at",
	TimeLastUse:  "refresh_tokens.time_last_use",
	DeletedAt:    "refresh_tokens.deleted_at",
	CreatedAt:    "refresh_tokens.created_at",
}

// Generated where

var RefreshTokenWhere = struct {
	ID           whereHelperint64
	UserID       whereHelperint64
	RefreshToken whereHelperstring
	ExpiresAt    whereHelpertime_Time
	TimeLastUse  whereHelpernull_Time
	DeletedAt    whereHelpernull_Time
	CreatedAt    whereHelpertime_Time
}{
	ID:           whereHelperint64{field: "\"refresh_tokens\".\"id\""},
	UserID:       whereHelperint64{field: "\"refresh_tokens\".\"user_id\""},
	RefreshToken: whereHelperstring{field: "\"refresh_tokens\".\"refresh_token\""},
	ExpiresAt:    whereHelpertime_Time{field: "\"refresh_tokens\".\"expires_at\""},
	TimeLastUse:  whereHelpernull_Time{field: "\"refresh_tokens\".\"time_last_use\""},
	DeletedAt:    whereHelpernull_Time{field: "\"refresh_tokens\".\"deleted_at\""},
	CreatedAt:    whereHelpertime_Time{field: "\"refresh_tokens\".\"created_at\""},
}

// RefreshTokenRels is where relationship names are stored.
var RefreshTokenRels = struct {
	User string
}{
	User: "User",
}

// refreshTokenR is where relationships are stored.
type refreshTokenR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*refreshTokenR) NewStruct() *refreshTokenR {
	return &refreshTokenR{}
}

func (r *refreshTokenR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// refreshTokenL is where Load methods for each relationship are stored.
type refreshTokenL struct{}

var (
	refreshTokenAllColumns            = []string{"id", "user_id", "refresh_token", "expires_at", "time_last_use", "deleted_at", "created_at"}
	refreshTokenColumnsWithoutDefault = []string{"user_id", "refresh_token", "expires_at"}
	refreshTokenColumnsWithDefault    = []string{"id", "time_last_use", "deleted_at", "created_at"}
	refreshTokenPrimaryKeyColumns     = []string{"id"}
	refreshTokenGeneratedColumns      = []string{}
)

type (
	// RefreshTokenSlice is an alias for a slice of pointers to RefreshToken.
	// This should almost always be used instead of []RefreshToken.
	RefreshTokenSlice []*RefreshToken
	// RefreshTokenHook is the signature for custom RefreshToken hook methods
	RefreshTokenHook func(context.Context, boil.ContextExecutor, *RefreshToken) error

	refreshTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	refreshTokenType                 = reflect.TypeOf(&RefreshToken{})
	refreshTokenMapping              = queries.MakeStructMapping(refreshTokenType)
	refreshTokenPrimaryKeyMapping, _ = queries.BindMapping(refreshTokenType, refreshTokenMapping, refreshTokenPrimaryKeyColumns)
	refreshTokenInsertCacheMut       sync.RWMutex
	refreshTokenInsertCache          = make(map[string]insertCache)
	refreshTokenUpdateCacheMut       sync.RWMutex
	refreshTokenUpdateCache          = make(map[string]updateCache)
	refreshTokenUpsertCacheMut       sync.RWMutex
	refreshTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var refreshTokenAfterSelectHooks []RefreshTokenHook

var refreshTokenBeforeInsertHooks []RefreshTokenHook
var refreshTokenAfterInsertHooks []RefreshTokenHook

var refreshTokenBeforeUpdateHooks []RefreshTokenHook
var refreshTokenAfterUpdateHooks []RefreshTokenHook

var refreshTokenBeforeDeleteHooks []RefreshTokenHook
var refreshTokenAfterDeleteHooks []RefreshTokenHook

var refreshTokenBeforeUpsertHooks []RefreshTokenHook
var refreshTokenAfterUpsertHooks []RefreshTokenHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RefreshToken) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RefreshToken) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RefreshToken) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RefreshToken) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RefreshToken) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RefreshToken) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RefreshToken) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RefreshToken) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RefreshToken) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshTokenAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRefreshTokenHook registers your hook function for all future operations.
func AddRefreshTokenHook(hookPoint boil.HookPoint, refreshTokenHook RefreshTokenHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		refreshTokenAfterSelectHooks = append(refreshTokenAfterSelectHooks, refreshTokenHook)
	case boil.BeforeInsertHook:
		refreshTokenBeforeInsertHooks = append(refreshTokenBeforeInsertHooks, refreshTokenHook)
	case boil.AfterInsertHook:
		refreshTokenAfterInsertHooks = append(refreshTokenAfterInsertHooks, refreshTokenHook)
	case boil.BeforeUpdateHook:
		refreshTokenBeforeUpdateHooks = append(refreshTokenBeforeUpdateHooks, refreshTokenHook)
	case boil.AfterUpdateHook:
		refreshTokenAfterUpdateHooks = append(refreshTokenAfterUpdateHooks, refreshTokenHook)
	case boil.BeforeDeleteHook:
		refreshTokenBeforeDeleteHooks = append(refreshTokenBeforeDeleteHooks, refreshTokenHook)
	case boil.AfterDeleteHook:
		refreshTokenAfterDeleteHooks = append(refreshTokenAfterDeleteHooks, refreshTokenHook)
	case boil.BeforeUpsertHook:
		refreshTokenBeforeUpsertHooks = append(refreshTokenBeforeUpsertHooks, refreshTokenHook)
	case boil.AfterUpsertHook:
		refreshTokenAfterUpsertHooks = append(refreshTokenAfterUpsertHooks, refreshTokenHook)
	}
}

// One returns a single refreshToken record from the query.
func (q refreshTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RefreshToken, error) {
	o := &RefreshToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for refresh_tokens")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RefreshToken records from the query.
func (q refreshTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (RefreshTokenSlice, error) {
	var o []*RefreshToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RefreshToken slice")
	}

	if len(refreshTokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RefreshToken records in the query.
func (q refreshTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count refresh_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q refreshTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if refresh_tokens exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *RefreshToken) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (refreshTokenL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRefreshToken interface{}, mods queries.Applicator) error {
	var slice []*RefreshToken
	var object *RefreshToken

	if singular {
		var ok bool
		object, ok = maybeRefreshToken.(*RefreshToken)
		if !ok {
			object = new(RefreshToken)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRefreshToken)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRefreshToken))
			}
		}
	} else {
		s, ok := maybeRefreshToken.(*[]*RefreshToken)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRefreshToken)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRefreshToken))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &refreshTokenR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &refreshTokenR{}
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
		foreign.R.RefreshTokens = append(foreign.R.RefreshTokens, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.RefreshTokens = append(foreign.R.RefreshTokens, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the refreshToken to the related item.
// Sets o.R.User to related.
// Adds o to related.R.RefreshTokens.
func (o *RefreshToken) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"refresh_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, refreshTokenPrimaryKeyColumns),
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
		o.R = &refreshTokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			RefreshTokens: RefreshTokenSlice{o},
		}
	} else {
		related.R.RefreshTokens = append(related.R.RefreshTokens, o)
	}

	return nil
}

// RefreshTokens retrieves all the records using an executor.
func RefreshTokens(mods ...qm.QueryMod) refreshTokenQuery {
	mods = append(mods, qm.From("\"refresh_tokens\""), qmhelper.WhereIsNull("\"refresh_tokens\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"refresh_tokens\".*"})
	}

	return refreshTokenQuery{q}
}

// FindRefreshToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRefreshToken(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*RefreshToken, error) {
	refreshTokenObj := &RefreshToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"refresh_tokens\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, refreshTokenObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from refresh_tokens")
	}

	if err = refreshTokenObj.doAfterSelectHooks(ctx, exec); err != nil {
		return refreshTokenObj, err
	}

	return refreshTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RefreshToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no refresh_tokens provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(refreshTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	refreshTokenInsertCacheMut.RLock()
	cache, cached := refreshTokenInsertCache[key]
	refreshTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			refreshTokenAllColumns,
			refreshTokenColumnsWithDefault,
			refreshTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"refresh_tokens\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"refresh_tokens\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into refresh_tokens")
	}

	if !cached {
		refreshTokenInsertCacheMut.Lock()
		refreshTokenInsertCache[key] = cache
		refreshTokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RefreshToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RefreshToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	refreshTokenUpdateCacheMut.RLock()
	cache, cached := refreshTokenUpdateCache[key]
	refreshTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			refreshTokenAllColumns,
			refreshTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update refresh_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"refresh_tokens\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, refreshTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, append(wl, refreshTokenPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update refresh_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for refresh_tokens")
	}

	if !cached {
		refreshTokenUpdateCacheMut.Lock()
		refreshTokenUpdateCache[key] = cache
		refreshTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q refreshTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for refresh_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for refresh_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RefreshTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"refresh_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, refreshTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in refreshToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all refreshToken")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RefreshToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no refresh_tokens provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(refreshTokenColumnsWithDefault, o)

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

	refreshTokenUpsertCacheMut.RLock()
	cache, cached := refreshTokenUpsertCache[key]
	refreshTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			refreshTokenAllColumns,
			refreshTokenColumnsWithDefault,
			refreshTokenColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			refreshTokenAllColumns,
			refreshTokenPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert refresh_tokens, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(refreshTokenPrimaryKeyColumns))
			copy(conflict, refreshTokenPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"refresh_tokens\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert refresh_tokens")
	}

	if !cached {
		refreshTokenUpsertCacheMut.Lock()
		refreshTokenUpsertCache[key] = cache
		refreshTokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RefreshToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RefreshToken) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RefreshToken provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), refreshTokenPrimaryKeyMapping)
		sql = "DELETE FROM \"refresh_tokens\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"refresh_tokens\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(refreshTokenType, refreshTokenMapping, append(wl, refreshTokenPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to delete from refresh_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for refresh_tokens")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q refreshTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no refreshTokenQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from refresh_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for refresh_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RefreshTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(refreshTokenBeforeDeleteHooks) != 0 {
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
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshTokenPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"refresh_tokens\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, refreshTokenPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshTokenPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"refresh_tokens\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, refreshTokenPrimaryKeyColumns, len(o)),
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
		return 0, errors.Wrap(err, "models: unable to delete all from refreshToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for refresh_tokens")
	}

	if len(refreshTokenAfterDeleteHooks) != 0 {
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
func (o *RefreshToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRefreshToken(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RefreshTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RefreshTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"refresh_tokens\".* FROM \"refresh_tokens\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, refreshTokenPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RefreshTokenSlice")
	}

	*o = slice

	return nil
}

// RefreshTokenExists checks if the RefreshToken row exists.
func RefreshTokenExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"refresh_tokens\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if refresh_tokens exists")
	}

	return exists, nil
}

// Exists checks if the RefreshToken row exists.
func (o *RefreshToken) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RefreshTokenExists(ctx, exec, o.ID)
}

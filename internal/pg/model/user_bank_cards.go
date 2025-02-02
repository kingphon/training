// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelpg

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

// UserBankCard is an object representing the database table.
type UserBankCard struct {
	ID        null.String `boil:"_id" json:"_id" toml:"_id" yaml:"_id"`
	Number    null.String `boil:"number" json:"number" toml:"number" yaml:"number"`
	UserID    null.String `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	IsDefault null.Bool   `boil:"is_default" json:"is_default,omitempty" toml:"is_default" yaml:"is_default,omitempty"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userBankCardR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userBankCardL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserBankCardColumns = struct {
	ID        string
	Number    string
	UserID    string
	IsDefault string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "_id",
	Number:    "number",
	UserID:    "user_id",
	IsDefault: "is_default",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserBankCardTableColumns = struct {
	ID        string
	Number    string
	UserID    string
	IsDefault string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "user_bank_cards._id",
	Number:    "user_bank_cards.number",
	UserID:    "user_bank_cards.user_id",
	IsDefault: "user_bank_cards.is_default",
	CreatedAt: "user_bank_cards.created_at",
	UpdatedAt: "user_bank_cards.updated_at",
}

// Generated where

var UserBankCardWhere = struct {
	ID        whereHelpernull_String
	Number    whereHelpernull_String
	UserID    whereHelpernull_String
	IsDefault whereHelpernull_Bool
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelpernull_String{field: "\"user_bank_cards\".\"_id\""},
	Number:    whereHelpernull_String{field: "\"user_bank_cards\".\"number\""},
	UserID:    whereHelpernull_String{field: "\"user_bank_cards\".\"user_id\""},
	IsDefault: whereHelpernull_Bool{field: "\"user_bank_cards\".\"is_default\""},
	CreatedAt: whereHelpernull_Time{field: "\"user_bank_cards\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"user_bank_cards\".\"updated_at\""},
}

// UserBankCardRels is where relationship names are stored.
var UserBankCardRels = struct {
	User string
}{
	User: "User",
}

// userBankCardR is where relationships are stored.
type userBankCardR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userBankCardR) NewStruct() *userBankCardR {
	return &userBankCardR{}
}

func (r *userBankCardR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userBankCardL is where Load methods for each relationship are stored.
type userBankCardL struct{}

var (
	userBankCardAllColumns            = []string{"_id", "number", "user_id", "is_default", "created_at", "updated_at"}
	userBankCardColumnsWithoutDefault = []string{"_id", "number", "user_id", "created_at", "updated_at"}
	userBankCardColumnsWithDefault    = []string{"is_default"}
	userBankCardPrimaryKeyColumns     = []string{"_id"}
	userBankCardGeneratedColumns      = []string{}
)

type (
	// UserBankCardSlice is an alias for a slice of pointers to UserBankCard.
	// This should almost always be used instead of []UserBankCard.
	UserBankCardSlice []*UserBankCard
	// UserBankCardHook is the signature for custom UserBankCard hook methods
	UserBankCardHook func(context.Context, boil.ContextExecutor, *UserBankCard) error

	userBankCardQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userBankCardType                 = reflect.TypeOf(&UserBankCard{})
	userBankCardMapping              = queries.MakeStructMapping(userBankCardType)
	userBankCardPrimaryKeyMapping, _ = queries.BindMapping(userBankCardType, userBankCardMapping, userBankCardPrimaryKeyColumns)
	userBankCardInsertCacheMut       sync.RWMutex
	userBankCardInsertCache          = make(map[string]insertCache)
	userBankCardUpdateCacheMut       sync.RWMutex
	userBankCardUpdateCache          = make(map[string]updateCache)
	userBankCardUpsertCacheMut       sync.RWMutex
	userBankCardUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userBankCardAfterSelectMu sync.Mutex
var userBankCardAfterSelectHooks []UserBankCardHook

var userBankCardBeforeInsertMu sync.Mutex
var userBankCardBeforeInsertHooks []UserBankCardHook
var userBankCardAfterInsertMu sync.Mutex
var userBankCardAfterInsertHooks []UserBankCardHook

var userBankCardBeforeUpdateMu sync.Mutex
var userBankCardBeforeUpdateHooks []UserBankCardHook
var userBankCardAfterUpdateMu sync.Mutex
var userBankCardAfterUpdateHooks []UserBankCardHook

var userBankCardBeforeDeleteMu sync.Mutex
var userBankCardBeforeDeleteHooks []UserBankCardHook
var userBankCardAfterDeleteMu sync.Mutex
var userBankCardAfterDeleteHooks []UserBankCardHook

var userBankCardBeforeUpsertMu sync.Mutex
var userBankCardBeforeUpsertHooks []UserBankCardHook
var userBankCardAfterUpsertMu sync.Mutex
var userBankCardAfterUpsertHooks []UserBankCardHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserBankCard) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserBankCard) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserBankCard) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserBankCard) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserBankCard) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserBankCard) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserBankCard) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserBankCard) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserBankCard) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBankCardAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserBankCardHook registers your hook function for all future operations.
func AddUserBankCardHook(hookPoint boil.HookPoint, userBankCardHook UserBankCardHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userBankCardAfterSelectMu.Lock()
		userBankCardAfterSelectHooks = append(userBankCardAfterSelectHooks, userBankCardHook)
		userBankCardAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userBankCardBeforeInsertMu.Lock()
		userBankCardBeforeInsertHooks = append(userBankCardBeforeInsertHooks, userBankCardHook)
		userBankCardBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userBankCardAfterInsertMu.Lock()
		userBankCardAfterInsertHooks = append(userBankCardAfterInsertHooks, userBankCardHook)
		userBankCardAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userBankCardBeforeUpdateMu.Lock()
		userBankCardBeforeUpdateHooks = append(userBankCardBeforeUpdateHooks, userBankCardHook)
		userBankCardBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userBankCardAfterUpdateMu.Lock()
		userBankCardAfterUpdateHooks = append(userBankCardAfterUpdateHooks, userBankCardHook)
		userBankCardAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userBankCardBeforeDeleteMu.Lock()
		userBankCardBeforeDeleteHooks = append(userBankCardBeforeDeleteHooks, userBankCardHook)
		userBankCardBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userBankCardAfterDeleteMu.Lock()
		userBankCardAfterDeleteHooks = append(userBankCardAfterDeleteHooks, userBankCardHook)
		userBankCardAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userBankCardBeforeUpsertMu.Lock()
		userBankCardBeforeUpsertHooks = append(userBankCardBeforeUpsertHooks, userBankCardHook)
		userBankCardBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userBankCardAfterUpsertMu.Lock()
		userBankCardAfterUpsertHooks = append(userBankCardAfterUpsertHooks, userBankCardHook)
		userBankCardAfterUpsertMu.Unlock()
	}
}

// One returns a single userBankCard record from the query.
func (q userBankCardQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserBankCard, error) {
	o := &UserBankCard{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "modelpg: failed to execute a one query for user_bank_cards")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserBankCard records from the query.
func (q userBankCardQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserBankCardSlice, error) {
	var o []*UserBankCard

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "modelpg: failed to assign all query results to UserBankCard slice")
	}

	if len(userBankCardAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserBankCard records in the query.
func (q userBankCardQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to count user_bank_cards rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userBankCardQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "modelpg: failed to check if user_bank_cards exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserBankCard) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userBankCardL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserBankCard interface{}, mods queries.Applicator) error {
	var slice []*UserBankCard
	var object *UserBankCard

	if singular {
		var ok bool
		object, ok = maybeUserBankCard.(*UserBankCard)
		if !ok {
			object = new(UserBankCard)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserBankCard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserBankCard))
			}
		}
	} else {
		s, ok := maybeUserBankCard.(*[]*UserBankCard)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserBankCard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserBankCard))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &userBankCardR{}
		}
		if !queries.IsNil(object.UserID) {
			args[object.UserID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userBankCardR{}
			}

			if !queries.IsNil(obj.UserID) {
				args[obj.UserID] = struct{}{}
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users._id in ?`, argsSlice...),
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
		foreign.R.UserBankCards = append(foreign.R.UserBankCards, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserBankCards = append(foreign.R.UserBankCards, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the userBankCard to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserBankCards.
func (o *UserBankCard) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_bank_cards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userBankCardPrimaryKeyColumns),
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

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &userBankCardR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserBankCards: UserBankCardSlice{o},
		}
	} else {
		related.R.UserBankCards = append(related.R.UserBankCards, o)
	}

	return nil
}

// UserBankCards retrieves all the records using an executor.
func UserBankCards(mods ...qm.QueryMod) userBankCardQuery {
	mods = append(mods, qm.From("\"user_bank_cards\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_bank_cards\".*"})
	}

	return userBankCardQuery{q}
}

// FindUserBankCard retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserBankCard(ctx context.Context, exec boil.ContextExecutor, iD null.String, selectCols ...string) (*UserBankCard, error) {
	userBankCardObj := &UserBankCard{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_bank_cards\" where \"_id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userBankCardObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "modelpg: unable to select from user_bank_cards")
	}

	if err = userBankCardObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userBankCardObj, err
	}

	return userBankCardObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserBankCard) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("modelpg: no user_bank_cards provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userBankCardColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userBankCardInsertCacheMut.RLock()
	cache, cached := userBankCardInsertCache[key]
	userBankCardInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userBankCardAllColumns,
			userBankCardColumnsWithDefault,
			userBankCardColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userBankCardType, userBankCardMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userBankCardType, userBankCardMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_bank_cards\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_bank_cards\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "modelpg: unable to insert into user_bank_cards")
	}

	if !cached {
		userBankCardInsertCacheMut.Lock()
		userBankCardInsertCache[key] = cache
		userBankCardInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserBankCard.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserBankCard) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userBankCardUpdateCacheMut.RLock()
	cache, cached := userBankCardUpdateCache[key]
	userBankCardUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userBankCardAllColumns,
			userBankCardPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("modelpg: unable to update user_bank_cards, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_bank_cards\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userBankCardPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userBankCardType, userBankCardMapping, append(wl, userBankCardPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "modelpg: unable to update user_bank_cards row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by update for user_bank_cards")
	}

	if !cached {
		userBankCardUpdateCacheMut.Lock()
		userBankCardUpdateCache[key] = cache
		userBankCardUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userBankCardQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to update all for user_bank_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to retrieve rows affected for user_bank_cards")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserBankCardSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("modelpg: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBankCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_bank_cards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userBankCardPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to update all in userBankCard slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to retrieve rows affected all in update all userBankCard")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserBankCard) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("modelpg: no user_bank_cards provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userBankCardColumnsWithDefault, o)

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

	userBankCardUpsertCacheMut.RLock()
	cache, cached := userBankCardUpsertCache[key]
	userBankCardUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			userBankCardAllColumns,
			userBankCardColumnsWithDefault,
			userBankCardColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userBankCardAllColumns,
			userBankCardPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("modelpg: unable to upsert user_bank_cards, could not build update column list")
		}

		ret := strmangle.SetComplement(userBankCardAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(userBankCardPrimaryKeyColumns) == 0 {
				return errors.New("modelpg: unable to upsert user_bank_cards, could not build conflict column list")
			}

			conflict = make([]string, len(userBankCardPrimaryKeyColumns))
			copy(conflict, userBankCardPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_bank_cards\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(userBankCardType, userBankCardMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userBankCardType, userBankCardMapping, ret)
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
		return errors.Wrap(err, "modelpg: unable to upsert user_bank_cards")
	}

	if !cached {
		userBankCardUpsertCacheMut.Lock()
		userBankCardUpsertCache[key] = cache
		userBankCardUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserBankCard record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserBankCard) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("modelpg: no UserBankCard provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userBankCardPrimaryKeyMapping)
	sql := "DELETE FROM \"user_bank_cards\" WHERE \"_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete from user_bank_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by delete for user_bank_cards")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userBankCardQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("modelpg: no userBankCardQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete all from user_bank_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by deleteall for user_bank_cards")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserBankCardSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userBankCardBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBankCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_bank_cards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userBankCardPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete all from userBankCard slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by deleteall for user_bank_cards")
	}

	if len(userBankCardAfterDeleteHooks) != 0 {
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
func (o *UserBankCard) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserBankCard(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserBankCardSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserBankCardSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBankCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_bank_cards\".* FROM \"user_bank_cards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userBankCardPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "modelpg: unable to reload all in UserBankCardSlice")
	}

	*o = slice

	return nil
}

// UserBankCardExists checks if the UserBankCard row exists.
func UserBankCardExists(ctx context.Context, exec boil.ContextExecutor, iD null.String) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_bank_cards\" where \"_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "modelpg: unable to check if user_bank_cards exists")
	}

	return exists, nil
}

// Exists checks if the UserBankCard row exists.
func (o *UserBankCard) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserBankCardExists(ctx, exec, o.ID)
}

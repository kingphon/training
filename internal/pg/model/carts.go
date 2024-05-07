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

// Cart is an object representing the database table.
type Cart struct {
	ID        null.String `boil:"_id" json:"_id" toml:"_id" yaml:"_id"`
	UserID    null.String `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *cartR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cartL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CartColumns = struct {
	ID        string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "_id",
	UserID:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var CartTableColumns = struct {
	ID        string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "carts._id",
	UserID:    "carts.user_id",
	CreatedAt: "carts.created_at",
	UpdatedAt: "carts.updated_at",
}

// Generated where

var CartWhere = struct {
	ID        whereHelpernull_String
	UserID    whereHelpernull_String
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelpernull_String{field: "\"carts\".\"_id\""},
	UserID:    whereHelpernull_String{field: "\"carts\".\"user_id\""},
	CreatedAt: whereHelpernull_Time{field: "\"carts\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"carts\".\"updated_at\""},
}

// CartRels is where relationship names are stored.
var CartRels = struct {
	User      string
	CartItems string
}{
	User:      "User",
	CartItems: "CartItems",
}

// cartR is where relationships are stored.
type cartR struct {
	User      *User         `boil:"User" json:"User" toml:"User" yaml:"User"`
	CartItems CartItemSlice `boil:"CartItems" json:"CartItems" toml:"CartItems" yaml:"CartItems"`
}

// NewStruct creates a new relationship struct
func (*cartR) NewStruct() *cartR {
	return &cartR{}
}

func (r *cartR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *cartR) GetCartItems() CartItemSlice {
	if r == nil {
		return nil
	}
	return r.CartItems
}

// cartL is where Load methods for each relationship are stored.
type cartL struct{}

var (
	cartAllColumns            = []string{"_id", "user_id", "created_at", "updated_at"}
	cartColumnsWithoutDefault = []string{"_id", "user_id", "created_at", "updated_at"}
	cartColumnsWithDefault    = []string{}
	cartPrimaryKeyColumns     = []string{"_id"}
	cartGeneratedColumns      = []string{}
)

type (
	// CartSlice is an alias for a slice of pointers to Cart.
	// This should almost always be used instead of []Cart.
	CartSlice []*Cart
	// CartHook is the signature for custom Cart hook methods
	CartHook func(context.Context, boil.ContextExecutor, *Cart) error

	cartQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cartType                 = reflect.TypeOf(&Cart{})
	cartMapping              = queries.MakeStructMapping(cartType)
	cartPrimaryKeyMapping, _ = queries.BindMapping(cartType, cartMapping, cartPrimaryKeyColumns)
	cartInsertCacheMut       sync.RWMutex
	cartInsertCache          = make(map[string]insertCache)
	cartUpdateCacheMut       sync.RWMutex
	cartUpdateCache          = make(map[string]updateCache)
	cartUpsertCacheMut       sync.RWMutex
	cartUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cartAfterSelectMu sync.Mutex
var cartAfterSelectHooks []CartHook

var cartBeforeInsertMu sync.Mutex
var cartBeforeInsertHooks []CartHook
var cartAfterInsertMu sync.Mutex
var cartAfterInsertHooks []CartHook

var cartBeforeUpdateMu sync.Mutex
var cartBeforeUpdateHooks []CartHook
var cartAfterUpdateMu sync.Mutex
var cartAfterUpdateHooks []CartHook

var cartBeforeDeleteMu sync.Mutex
var cartBeforeDeleteHooks []CartHook
var cartAfterDeleteMu sync.Mutex
var cartAfterDeleteHooks []CartHook

var cartBeforeUpsertMu sync.Mutex
var cartBeforeUpsertHooks []CartHook
var cartAfterUpsertMu sync.Mutex
var cartAfterUpsertHooks []CartHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cart) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cart) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cart) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cart) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cart) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cart) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cart) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cart) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cart) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cartAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCartHook registers your hook function for all future operations.
func AddCartHook(hookPoint boil.HookPoint, cartHook CartHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		cartAfterSelectMu.Lock()
		cartAfterSelectHooks = append(cartAfterSelectHooks, cartHook)
		cartAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		cartBeforeInsertMu.Lock()
		cartBeforeInsertHooks = append(cartBeforeInsertHooks, cartHook)
		cartBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		cartAfterInsertMu.Lock()
		cartAfterInsertHooks = append(cartAfterInsertHooks, cartHook)
		cartAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		cartBeforeUpdateMu.Lock()
		cartBeforeUpdateHooks = append(cartBeforeUpdateHooks, cartHook)
		cartBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		cartAfterUpdateMu.Lock()
		cartAfterUpdateHooks = append(cartAfterUpdateHooks, cartHook)
		cartAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		cartBeforeDeleteMu.Lock()
		cartBeforeDeleteHooks = append(cartBeforeDeleteHooks, cartHook)
		cartBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		cartAfterDeleteMu.Lock()
		cartAfterDeleteHooks = append(cartAfterDeleteHooks, cartHook)
		cartAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		cartBeforeUpsertMu.Lock()
		cartBeforeUpsertHooks = append(cartBeforeUpsertHooks, cartHook)
		cartBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		cartAfterUpsertMu.Lock()
		cartAfterUpsertHooks = append(cartAfterUpsertHooks, cartHook)
		cartAfterUpsertMu.Unlock()
	}
}

// One returns a single cart record from the query.
func (q cartQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Cart, error) {
	o := &Cart{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "modelpg: failed to execute a one query for carts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Cart records from the query.
func (q cartQuery) All(ctx context.Context, exec boil.ContextExecutor) (CartSlice, error) {
	var o []*Cart

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "modelpg: failed to assign all query results to Cart slice")
	}

	if len(cartAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Cart records in the query.
func (q cartQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to count carts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cartQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "modelpg: failed to check if carts exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Cart) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// CartItems retrieves all the cart_item's CartItems with an executor.
func (o *Cart) CartItems(mods ...qm.QueryMod) cartItemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"cart_items\".\"cart_id\"=?", o.ID),
	)

	return CartItems(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (cartL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCart interface{}, mods queries.Applicator) error {
	var slice []*Cart
	var object *Cart

	if singular {
		var ok bool
		object, ok = maybeCart.(*Cart)
		if !ok {
			object = new(Cart)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCart)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCart))
			}
		}
	} else {
		s, ok := maybeCart.(*[]*Cart)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCart)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCart))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &cartR{}
		}
		if !queries.IsNil(object.UserID) {
			args[object.UserID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cartR{}
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
		foreign.R.Carts = append(foreign.R.Carts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Carts = append(foreign.R.Carts, local)
				break
			}
		}
	}

	return nil
}

// LoadCartItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (cartL) LoadCartItems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCart interface{}, mods queries.Applicator) error {
	var slice []*Cart
	var object *Cart

	if singular {
		var ok bool
		object, ok = maybeCart.(*Cart)
		if !ok {
			object = new(Cart)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCart)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCart))
			}
		}
	} else {
		s, ok := maybeCart.(*[]*Cart)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCart)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCart))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &cartR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cartR{}
			}
			args[obj.ID] = struct{}{}
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
		qm.From(`cart_items`),
		qm.WhereIn(`cart_items.cart_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load cart_items")
	}

	var resultSlice []*CartItem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice cart_items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on cart_items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for cart_items")
	}

	if len(cartItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CartItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &cartItemR{}
			}
			foreign.R.Cart = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.CartID) {
				local.R.CartItems = append(local.R.CartItems, foreign)
				if foreign.R == nil {
					foreign.R = &cartItemR{}
				}
				foreign.R.Cart = local
				break
			}
		}
	}

	return nil
}

// SetUser of the cart to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Carts.
func (o *Cart) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"carts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, cartPrimaryKeyColumns),
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
		o.R = &cartR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Carts: CartSlice{o},
		}
	} else {
		related.R.Carts = append(related.R.Carts, o)
	}

	return nil
}

// AddCartItems adds the given related objects to the existing relationships
// of the cart, optionally inserting them as new records.
// Appends related to o.R.CartItems.
// Sets related.R.Cart appropriately.
func (o *Cart) AddCartItems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CartItem) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.CartID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"cart_items\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"cart_id"}),
				strmangle.WhereClause("\"", "\"", 2, cartItemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ShopID, rel.CartID, rel.ProductID, rel.SkuID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.CartID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &cartR{
			CartItems: related,
		}
	} else {
		o.R.CartItems = append(o.R.CartItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &cartItemR{
				Cart: o,
			}
		} else {
			rel.R.Cart = o
		}
	}
	return nil
}

// Carts retrieves all the records using an executor.
func Carts(mods ...qm.QueryMod) cartQuery {
	mods = append(mods, qm.From("\"carts\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"carts\".*"})
	}

	return cartQuery{q}
}

// FindCart retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCart(ctx context.Context, exec boil.ContextExecutor, iD null.String, selectCols ...string) (*Cart, error) {
	cartObj := &Cart{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"carts\" where \"_id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cartObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "modelpg: unable to select from carts")
	}

	if err = cartObj.doAfterSelectHooks(ctx, exec); err != nil {
		return cartObj, err
	}

	return cartObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Cart) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("modelpg: no carts provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(cartColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cartInsertCacheMut.RLock()
	cache, cached := cartInsertCache[key]
	cartInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cartAllColumns,
			cartColumnsWithDefault,
			cartColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cartType, cartMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cartType, cartMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"carts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"carts\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "modelpg: unable to insert into carts")
	}

	if !cached {
		cartInsertCacheMut.Lock()
		cartInsertCache[key] = cache
		cartInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Cart.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Cart) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cartUpdateCacheMut.RLock()
	cache, cached := cartUpdateCache[key]
	cartUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cartAllColumns,
			cartPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("modelpg: unable to update carts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"carts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cartPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cartType, cartMapping, append(wl, cartPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "modelpg: unable to update carts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by update for carts")
	}

	if !cached {
		cartUpdateCacheMut.Lock()
		cartUpdateCache[key] = cache
		cartUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cartQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to update all for carts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to retrieve rows affected for carts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CartSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"carts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, cartPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to update all in cart slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to retrieve rows affected all in update all cart")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Cart) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("modelpg: no carts provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(cartColumnsWithDefault, o)

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

	cartUpsertCacheMut.RLock()
	cache, cached := cartUpsertCache[key]
	cartUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			cartAllColumns,
			cartColumnsWithDefault,
			cartColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			cartAllColumns,
			cartPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("modelpg: unable to upsert carts, could not build update column list")
		}

		ret := strmangle.SetComplement(cartAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(cartPrimaryKeyColumns) == 0 {
				return errors.New("modelpg: unable to upsert carts, could not build conflict column list")
			}

			conflict = make([]string, len(cartPrimaryKeyColumns))
			copy(conflict, cartPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"carts\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(cartType, cartMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cartType, cartMapping, ret)
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
		return errors.Wrap(err, "modelpg: unable to upsert carts")
	}

	if !cached {
		cartUpsertCacheMut.Lock()
		cartUpsertCache[key] = cache
		cartUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Cart record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cart) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("modelpg: no Cart provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cartPrimaryKeyMapping)
	sql := "DELETE FROM \"carts\" WHERE \"_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete from carts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by delete for carts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cartQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("modelpg: no cartQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete all from carts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by deleteall for carts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CartSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cartBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"carts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cartPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete all from cart slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by deleteall for carts")
	}

	if len(cartAfterDeleteHooks) != 0 {
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
func (o *Cart) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCart(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CartSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CartSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"carts\".* FROM \"carts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cartPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "modelpg: unable to reload all in CartSlice")
	}

	*o = slice

	return nil
}

// CartExists checks if the Cart row exists.
func CartExists(ctx context.Context, exec boil.ContextExecutor, iD null.String) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"carts\" where \"_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "modelpg: unable to check if carts exists")
	}

	return exists, nil
}

// Exists checks if the Cart row exists.
func (o *Cart) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CartExists(ctx, exec, o.ID)
}

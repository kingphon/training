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
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OrderSession is an object representing the database table.
type OrderSession struct {
	ID                null.String     `boil:"_id" json:"_id" toml:"_id" yaml:"_id"`
	OrderItems        null.JSON       `boil:"order_items" json:"order_items,omitempty" toml:"order_items" yaml:"order_items,omitempty"`
	UserID            null.String     `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	DeliverySessionID null.String     `boil:"delivery_session_id" json:"delivery_session_id,omitempty" toml:"delivery_session_id" yaml:"delivery_session_id,omitempty"`
	Amount            decimal.Decimal `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	CreatedAt         null.Time       `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *orderSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderSessionColumns = struct {
	ID                string
	OrderItems        string
	UserID            string
	DeliverySessionID string
	Amount            string
	CreatedAt         string
}{
	ID:                "_id",
	OrderItems:        "order_items",
	UserID:            "user_id",
	DeliverySessionID: "delivery_session_id",
	Amount:            "amount",
	CreatedAt:         "created_at",
}

var OrderSessionTableColumns = struct {
	ID                string
	OrderItems        string
	UserID            string
	DeliverySessionID string
	Amount            string
	CreatedAt         string
}{
	ID:                "order_sessions._id",
	OrderItems:        "order_sessions.order_items",
	UserID:            "order_sessions.user_id",
	DeliverySessionID: "order_sessions.delivery_session_id",
	Amount:            "order_sessions.amount",
	CreatedAt:         "order_sessions.created_at",
}

// Generated where

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var OrderSessionWhere = struct {
	ID                whereHelpernull_String
	OrderItems        whereHelpernull_JSON
	UserID            whereHelpernull_String
	DeliverySessionID whereHelpernull_String
	Amount            whereHelperdecimal_Decimal
	CreatedAt         whereHelpernull_Time
}{
	ID:                whereHelpernull_String{field: "\"order_sessions\".\"_id\""},
	OrderItems:        whereHelpernull_JSON{field: "\"order_sessions\".\"order_items\""},
	UserID:            whereHelpernull_String{field: "\"order_sessions\".\"user_id\""},
	DeliverySessionID: whereHelpernull_String{field: "\"order_sessions\".\"delivery_session_id\""},
	Amount:            whereHelperdecimal_Decimal{field: "\"order_sessions\".\"amount\""},
	CreatedAt:         whereHelpernull_Time{field: "\"order_sessions\".\"created_at\""},
}

// OrderSessionRels is where relationship names are stored.
var OrderSessionRels = struct {
}{}

// orderSessionR is where relationships are stored.
type orderSessionR struct {
}

// NewStruct creates a new relationship struct
func (*orderSessionR) NewStruct() *orderSessionR {
	return &orderSessionR{}
}

// orderSessionL is where Load methods for each relationship are stored.
type orderSessionL struct{}

var (
	orderSessionAllColumns            = []string{"_id", "order_items", "user_id", "delivery_session_id", "amount", "created_at"}
	orderSessionColumnsWithoutDefault = []string{"_id", "amount", "created_at"}
	orderSessionColumnsWithDefault    = []string{"order_items", "user_id", "delivery_session_id"}
	orderSessionPrimaryKeyColumns     = []string{"_id"}
	orderSessionGeneratedColumns      = []string{}
)

type (
	// OrderSessionSlice is an alias for a slice of pointers to OrderSession.
	// This should almost always be used instead of []OrderSession.
	OrderSessionSlice []*OrderSession
	// OrderSessionHook is the signature for custom OrderSession hook methods
	OrderSessionHook func(context.Context, boil.ContextExecutor, *OrderSession) error

	orderSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderSessionType                 = reflect.TypeOf(&OrderSession{})
	orderSessionMapping              = queries.MakeStructMapping(orderSessionType)
	orderSessionPrimaryKeyMapping, _ = queries.BindMapping(orderSessionType, orderSessionMapping, orderSessionPrimaryKeyColumns)
	orderSessionInsertCacheMut       sync.RWMutex
	orderSessionInsertCache          = make(map[string]insertCache)
	orderSessionUpdateCacheMut       sync.RWMutex
	orderSessionUpdateCache          = make(map[string]updateCache)
	orderSessionUpsertCacheMut       sync.RWMutex
	orderSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderSessionAfterSelectMu sync.Mutex
var orderSessionAfterSelectHooks []OrderSessionHook

var orderSessionBeforeInsertMu sync.Mutex
var orderSessionBeforeInsertHooks []OrderSessionHook
var orderSessionAfterInsertMu sync.Mutex
var orderSessionAfterInsertHooks []OrderSessionHook

var orderSessionBeforeUpdateMu sync.Mutex
var orderSessionBeforeUpdateHooks []OrderSessionHook
var orderSessionAfterUpdateMu sync.Mutex
var orderSessionAfterUpdateHooks []OrderSessionHook

var orderSessionBeforeDeleteMu sync.Mutex
var orderSessionBeforeDeleteHooks []OrderSessionHook
var orderSessionAfterDeleteMu sync.Mutex
var orderSessionAfterDeleteHooks []OrderSessionHook

var orderSessionBeforeUpsertMu sync.Mutex
var orderSessionBeforeUpsertHooks []OrderSessionHook
var orderSessionAfterUpsertMu sync.Mutex
var orderSessionAfterUpsertHooks []OrderSessionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrderSession) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrderSession) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrderSession) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrderSession) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrderSession) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrderSession) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrderSession) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrderSession) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrderSession) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderSessionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderSessionHook registers your hook function for all future operations.
func AddOrderSessionHook(hookPoint boil.HookPoint, orderSessionHook OrderSessionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderSessionAfterSelectMu.Lock()
		orderSessionAfterSelectHooks = append(orderSessionAfterSelectHooks, orderSessionHook)
		orderSessionAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		orderSessionBeforeInsertMu.Lock()
		orderSessionBeforeInsertHooks = append(orderSessionBeforeInsertHooks, orderSessionHook)
		orderSessionBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		orderSessionAfterInsertMu.Lock()
		orderSessionAfterInsertHooks = append(orderSessionAfterInsertHooks, orderSessionHook)
		orderSessionAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		orderSessionBeforeUpdateMu.Lock()
		orderSessionBeforeUpdateHooks = append(orderSessionBeforeUpdateHooks, orderSessionHook)
		orderSessionBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		orderSessionAfterUpdateMu.Lock()
		orderSessionAfterUpdateHooks = append(orderSessionAfterUpdateHooks, orderSessionHook)
		orderSessionAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		orderSessionBeforeDeleteMu.Lock()
		orderSessionBeforeDeleteHooks = append(orderSessionBeforeDeleteHooks, orderSessionHook)
		orderSessionBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		orderSessionAfterDeleteMu.Lock()
		orderSessionAfterDeleteHooks = append(orderSessionAfterDeleteHooks, orderSessionHook)
		orderSessionAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		orderSessionBeforeUpsertMu.Lock()
		orderSessionBeforeUpsertHooks = append(orderSessionBeforeUpsertHooks, orderSessionHook)
		orderSessionBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		orderSessionAfterUpsertMu.Lock()
		orderSessionAfterUpsertHooks = append(orderSessionAfterUpsertHooks, orderSessionHook)
		orderSessionAfterUpsertMu.Unlock()
	}
}

// One returns a single orderSession record from the query.
func (q orderSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderSession, error) {
	o := &OrderSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "modelpg: failed to execute a one query for order_sessions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrderSession records from the query.
func (q orderSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderSessionSlice, error) {
	var o []*OrderSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "modelpg: failed to assign all query results to OrderSession slice")
	}

	if len(orderSessionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrderSession records in the query.
func (q orderSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to count order_sessions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orderSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "modelpg: failed to check if order_sessions exists")
	}

	return count > 0, nil
}

// OrderSessions retrieves all the records using an executor.
func OrderSessions(mods ...qm.QueryMod) orderSessionQuery {
	mods = append(mods, qm.From("\"order_sessions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"order_sessions\".*"})
	}

	return orderSessionQuery{q}
}

// FindOrderSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrderSession(ctx context.Context, exec boil.ContextExecutor, iD null.String, selectCols ...string) (*OrderSession, error) {
	orderSessionObj := &OrderSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"order_sessions\" where \"_id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderSessionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "modelpg: unable to select from order_sessions")
	}

	if err = orderSessionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderSessionObj, err
	}

	return orderSessionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("modelpg: no order_sessions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderSessionInsertCacheMut.RLock()
	cache, cached := orderSessionInsertCache[key]
	orderSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderSessionAllColumns,
			orderSessionColumnsWithDefault,
			orderSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderSessionType, orderSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderSessionType, orderSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"order_sessions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"order_sessions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "modelpg: unable to insert into order_sessions")
	}

	if !cached {
		orderSessionInsertCacheMut.Lock()
		orderSessionInsertCache[key] = cache
		orderSessionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OrderSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrderSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderSessionUpdateCacheMut.RLock()
	cache, cached := orderSessionUpdateCache[key]
	orderSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderSessionAllColumns,
			orderSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("modelpg: unable to update order_sessions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"order_sessions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orderSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderSessionType, orderSessionMapping, append(wl, orderSessionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "modelpg: unable to update order_sessions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by update for order_sessions")
	}

	if !cached {
		orderSessionUpdateCacheMut.Lock()
		orderSessionUpdateCache[key] = cache
		orderSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q orderSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to update all for order_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to retrieve rows affected for order_sessions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"order_sessions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderSessionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to update all in orderSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to retrieve rows affected all in update all orderSession")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrderSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("modelpg: no order_sessions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderSessionColumnsWithDefault, o)

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

	orderSessionUpsertCacheMut.RLock()
	cache, cached := orderSessionUpsertCache[key]
	orderSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			orderSessionAllColumns,
			orderSessionColumnsWithDefault,
			orderSessionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			orderSessionAllColumns,
			orderSessionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("modelpg: unable to upsert order_sessions, could not build update column list")
		}

		ret := strmangle.SetComplement(orderSessionAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(orderSessionPrimaryKeyColumns) == 0 {
				return errors.New("modelpg: unable to upsert order_sessions, could not build conflict column list")
			}

			conflict = make([]string, len(orderSessionPrimaryKeyColumns))
			copy(conflict, orderSessionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"order_sessions\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(orderSessionType, orderSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderSessionType, orderSessionMapping, ret)
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
		return errors.Wrap(err, "modelpg: unable to upsert order_sessions")
	}

	if !cached {
		orderSessionUpsertCacheMut.Lock()
		orderSessionUpsertCache[key] = cache
		orderSessionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OrderSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrderSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("modelpg: no OrderSession provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderSessionPrimaryKeyMapping)
	sql := "DELETE FROM \"order_sessions\" WHERE \"_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete from order_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by delete for order_sessions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orderSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("modelpg: no orderSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete all from order_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by deleteall for order_sessions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderSessionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"order_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderSessionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: unable to delete all from orderSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "modelpg: failed to get rows affected by deleteall for order_sessions")
	}

	if len(orderSessionAfterDeleteHooks) != 0 {
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
func (o *OrderSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrderSession(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"order_sessions\".* FROM \"order_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "modelpg: unable to reload all in OrderSessionSlice")
	}

	*o = slice

	return nil
}

// OrderSessionExists checks if the OrderSession row exists.
func OrderSessionExists(ctx context.Context, exec boil.ContextExecutor, iD null.String) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"order_sessions\" where \"_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "modelpg: unable to check if order_sessions exists")
	}

	return exists, nil
}

// Exists checks if the OrderSession row exists.
func (o *OrderSession) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrderSessionExists(ctx, exec, o.ID)
}

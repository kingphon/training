// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelpg

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWarehouseBankCards(t *testing.T) {
	t.Parallel()

	query := WarehouseBankCards()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWarehouseBankCardsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWarehouseBankCardsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := WarehouseBankCards().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWarehouseBankCardsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WarehouseBankCardSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWarehouseBankCardsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WarehouseBankCardExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if WarehouseBankCard exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WarehouseBankCardExists to return true, but got false.")
	}
}

func testWarehouseBankCardsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	warehouseBankCardFound, err := FindWarehouseBankCard(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if warehouseBankCardFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWarehouseBankCardsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = WarehouseBankCards().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWarehouseBankCardsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := WarehouseBankCards().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWarehouseBankCardsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	warehouseBankCardOne := &WarehouseBankCard{}
	warehouseBankCardTwo := &WarehouseBankCard{}
	if err = randomize.Struct(seed, warehouseBankCardOne, warehouseBankCardDBTypes, false, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}
	if err = randomize.Struct(seed, warehouseBankCardTwo, warehouseBankCardDBTypes, false, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = warehouseBankCardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = warehouseBankCardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WarehouseBankCards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWarehouseBankCardsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	warehouseBankCardOne := &WarehouseBankCard{}
	warehouseBankCardTwo := &WarehouseBankCard{}
	if err = randomize.Struct(seed, warehouseBankCardOne, warehouseBankCardDBTypes, false, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}
	if err = randomize.Struct(seed, warehouseBankCardTwo, warehouseBankCardDBTypes, false, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = warehouseBankCardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = warehouseBankCardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func warehouseBankCardBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func warehouseBankCardAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WarehouseBankCard) error {
	*o = WarehouseBankCard{}
	return nil
}

func testWarehouseBankCardsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &WarehouseBankCard{}
	o := &WarehouseBankCard{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard object: %s", err)
	}

	AddWarehouseBankCardHook(boil.BeforeInsertHook, warehouseBankCardBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardBeforeInsertHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.AfterInsertHook, warehouseBankCardAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardAfterInsertHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.AfterSelectHook, warehouseBankCardAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardAfterSelectHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.BeforeUpdateHook, warehouseBankCardBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardBeforeUpdateHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.AfterUpdateHook, warehouseBankCardAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardAfterUpdateHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.BeforeDeleteHook, warehouseBankCardBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardBeforeDeleteHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.AfterDeleteHook, warehouseBankCardAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardAfterDeleteHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.BeforeUpsertHook, warehouseBankCardBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardBeforeUpsertHooks = []WarehouseBankCardHook{}

	AddWarehouseBankCardHook(boil.AfterUpsertHook, warehouseBankCardAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	warehouseBankCardAfterUpsertHooks = []WarehouseBankCardHook{}
}

func testWarehouseBankCardsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWarehouseBankCardsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(warehouseBankCardColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWarehouseBankCardToOneWarehouseUsingWarehouse(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local WarehouseBankCard
	var foreign Warehouse

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, warehouseBankCardDBTypes, false, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, warehouseDBTypes, false, warehouseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Warehouse struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.WarehouseID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Warehouse().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddWarehouseHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Warehouse) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := WarehouseBankCardSlice{&local}
	if err = local.L.LoadWarehouse(ctx, tx, false, (*[]*WarehouseBankCard)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Warehouse == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Warehouse = nil
	if err = local.L.LoadWarehouse(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Warehouse == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testWarehouseBankCardToOneSetOpWarehouseUsingWarehouse(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WarehouseBankCard
	var b, c Warehouse

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, warehouseBankCardDBTypes, false, strmangle.SetComplement(warehouseBankCardPrimaryKeyColumns, warehouseBankCardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, warehouseDBTypes, false, strmangle.SetComplement(warehousePrimaryKeyColumns, warehouseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, warehouseDBTypes, false, strmangle.SetComplement(warehousePrimaryKeyColumns, warehouseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Warehouse{&b, &c} {
		err = a.SetWarehouse(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Warehouse != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.WarehouseBankCards[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.WarehouseID, x.ID) {
			t.Error("foreign key was wrong value", a.WarehouseID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.WarehouseID))
		reflect.Indirect(reflect.ValueOf(&a.WarehouseID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.WarehouseID, x.ID) {
			t.Error("foreign key was wrong value", a.WarehouseID, x.ID)
		}
	}
}

func testWarehouseBankCardsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWarehouseBankCardsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WarehouseBankCardSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWarehouseBankCardsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WarehouseBankCards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	warehouseBankCardDBTypes = map[string]string{`ID`: `text`, `Number`: `text`, `WarehouseID`: `text`, `IsDefault`: `boolean`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                        = bytes.MinRead
)

func testWarehouseBankCardsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(warehouseBankCardPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(warehouseBankCardAllColumns) == len(warehouseBankCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWarehouseBankCardsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(warehouseBankCardAllColumns) == len(warehouseBankCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WarehouseBankCard{}
	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, warehouseBankCardDBTypes, true, warehouseBankCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(warehouseBankCardAllColumns, warehouseBankCardPrimaryKeyColumns) {
		fields = warehouseBankCardAllColumns
	} else {
		fields = strmangle.SetComplement(
			warehouseBankCardAllColumns,
			warehouseBankCardPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WarehouseBankCardSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testWarehouseBankCardsUpsert(t *testing.T) {
	t.Parallel()

	if len(warehouseBankCardAllColumns) == len(warehouseBankCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := WarehouseBankCard{}
	if err = randomize.Struct(seed, &o, warehouseBankCardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert WarehouseBankCard: %s", err)
	}

	count, err := WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, warehouseBankCardDBTypes, false, warehouseBankCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WarehouseBankCard struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert WarehouseBankCard: %s", err)
	}

	count, err = WarehouseBankCards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

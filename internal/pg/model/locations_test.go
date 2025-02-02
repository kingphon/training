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

func testLocations(t *testing.T) {
	t.Parallel()

	query := Locations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLocationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
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

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLocationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Locations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLocationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LocationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLocationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LocationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Location exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LocationExists to return true, but got false.")
	}
}

func testLocationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	locationFound, err := FindLocation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if locationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLocationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Locations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLocationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Locations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLocationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	locationOne := &Location{}
	locationTwo := &Location{}
	if err = randomize.Struct(seed, locationOne, locationDBTypes, false, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}
	if err = randomize.Struct(seed, locationTwo, locationDBTypes, false, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = locationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = locationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Locations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLocationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	locationOne := &Location{}
	locationTwo := &Location{}
	if err = randomize.Struct(seed, locationOne, locationDBTypes, false, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}
	if err = randomize.Struct(seed, locationTwo, locationDBTypes, false, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = locationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = locationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func locationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func locationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Location) error {
	*o = Location{}
	return nil
}

func testLocationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Location{}
	o := &Location{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, locationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Location object: %s", err)
	}

	AddLocationHook(boil.BeforeInsertHook, locationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	locationBeforeInsertHooks = []LocationHook{}

	AddLocationHook(boil.AfterInsertHook, locationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	locationAfterInsertHooks = []LocationHook{}

	AddLocationHook(boil.AfterSelectHook, locationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	locationAfterSelectHooks = []LocationHook{}

	AddLocationHook(boil.BeforeUpdateHook, locationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	locationBeforeUpdateHooks = []LocationHook{}

	AddLocationHook(boil.AfterUpdateHook, locationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	locationAfterUpdateHooks = []LocationHook{}

	AddLocationHook(boil.BeforeDeleteHook, locationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	locationBeforeDeleteHooks = []LocationHook{}

	AddLocationHook(boil.AfterDeleteHook, locationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	locationAfterDeleteHooks = []LocationHook{}

	AddLocationHook(boil.BeforeUpsertHook, locationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	locationBeforeUpsertHooks = []LocationHook{}

	AddLocationHook(boil.AfterUpsertHook, locationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	locationAfterUpsertHooks = []LocationHook{}
}

func testLocationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLocationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(locationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLocationToManyDeliveries(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Location
	var b, c Delivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, deliveryDBTypes, false, deliveryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, deliveryDBTypes, false, deliveryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.LocationID, a.ID)
	queries.Assign(&c.LocationID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Deliveries().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.LocationID, b.LocationID) {
			bFound = true
		}
		if queries.Equal(v.LocationID, c.LocationID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LocationSlice{&a}
	if err = a.L.LoadDeliveries(ctx, tx, false, (*[]*Location)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Deliveries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Deliveries = nil
	if err = a.L.LoadDeliveries(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Deliveries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testLocationToManyAddOpDeliveries(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Location
	var b, c, d, e Delivery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, locationDBTypes, false, strmangle.SetComplement(locationPrimaryKeyColumns, locationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Delivery{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, deliveryDBTypes, false, strmangle.SetComplement(deliveryPrimaryKeyColumns, deliveryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Delivery{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDeliveries(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.LocationID) {
			t.Error("foreign key was wrong value", a.ID, first.LocationID)
		}
		if !queries.Equal(a.ID, second.LocationID) {
			t.Error("foreign key was wrong value", a.ID, second.LocationID)
		}

		if first.R.Location != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Location != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Deliveries[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Deliveries[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Deliveries().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testLocationToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Location
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, locationDBTypes, false, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.UserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := LocationSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Location)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testLocationToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Location
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, locationDBTypes, false, strmangle.SetComplement(locationPrimaryKeyColumns, locationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Locations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testLocationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
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

func testLocationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LocationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLocationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Locations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	locationDBTypes = map[string]string{`ID`: `text`, `Name`: `text`, `Address`: `text`, `Status`: `text`, `IsDefault`: `boolean`, `UserID`: `text`, `Latitude`: `numeric`, `Longitude`: `numeric`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testLocationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(locationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(locationAllColumns) == len(locationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, locationDBTypes, true, locationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLocationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(locationAllColumns) == len(locationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Location{}
	if err = randomize.Struct(seed, o, locationDBTypes, true, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, locationDBTypes, true, locationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(locationAllColumns, locationPrimaryKeyColumns) {
		fields = locationAllColumns
	} else {
		fields = strmangle.SetComplement(
			locationAllColumns,
			locationPrimaryKeyColumns,
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

	slice := LocationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLocationsUpsert(t *testing.T) {
	t.Parallel()

	if len(locationAllColumns) == len(locationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Location{}
	if err = randomize.Struct(seed, &o, locationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Location: %s", err)
	}

	count, err := Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, locationDBTypes, false, locationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Location: %s", err)
	}

	count, err = Locations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

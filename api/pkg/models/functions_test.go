// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFunctions(t *testing.T) {
	t.Parallel()

	query := Functions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFunctionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
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

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFunctionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Functions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFunctionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FunctionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFunctionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FunctionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Function exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FunctionExists to return true, but got false.")
	}
}

func testFunctionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	functionFound, err := FindFunction(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if functionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFunctionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Functions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFunctionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Functions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFunctionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	functionOne := &Function{}
	functionTwo := &Function{}
	if err = randomize.Struct(seed, functionOne, functionDBTypes, false, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}
	if err = randomize.Struct(seed, functionTwo, functionDBTypes, false, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = functionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = functionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Functions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFunctionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	functionOne := &Function{}
	functionTwo := &Function{}
	if err = randomize.Struct(seed, functionOne, functionDBTypes, false, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}
	if err = randomize.Struct(seed, functionTwo, functionDBTypes, false, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = functionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = functionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func functionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func functionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Function) error {
	*o = Function{}
	return nil
}

func testFunctionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Function{}
	o := &Function{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, functionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Function object: %s", err)
	}

	AddFunctionHook(boil.BeforeInsertHook, functionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	functionBeforeInsertHooks = []FunctionHook{}

	AddFunctionHook(boil.AfterInsertHook, functionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	functionAfterInsertHooks = []FunctionHook{}

	AddFunctionHook(boil.AfterSelectHook, functionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	functionAfterSelectHooks = []FunctionHook{}

	AddFunctionHook(boil.BeforeUpdateHook, functionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	functionBeforeUpdateHooks = []FunctionHook{}

	AddFunctionHook(boil.AfterUpdateHook, functionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	functionAfterUpdateHooks = []FunctionHook{}

	AddFunctionHook(boil.BeforeDeleteHook, functionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	functionBeforeDeleteHooks = []FunctionHook{}

	AddFunctionHook(boil.AfterDeleteHook, functionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	functionAfterDeleteHooks = []FunctionHook{}

	AddFunctionHook(boil.BeforeUpsertHook, functionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	functionBeforeUpsertHooks = []FunctionHook{}

	AddFunctionHook(boil.AfterUpsertHook, functionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	functionAfterUpsertHooks = []FunctionHook{}
}

func testFunctionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFunctionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(functionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFunctionToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Function
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, functionDBTypes, false, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FunctionSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Function)(&slice), nil); err != nil {
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
}

func testFunctionToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Function
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, functionDBTypes, false, strmangle.SetComplement(functionPrimaryKeyColumns, functionColumnsWithoutDefault)...); err != nil {
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

		if x.R.Functions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testFunctionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
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

func testFunctionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FunctionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFunctionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Functions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	functionDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `Tag`: `character varying`, `Hash`: `character varying`, `UserID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testFunctionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(functionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(functionAllColumns) == len(functionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, functionDBTypes, true, functionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFunctionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(functionAllColumns) == len(functionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Function{}
	if err = randomize.Struct(seed, o, functionDBTypes, true, functionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, functionDBTypes, true, functionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(functionAllColumns, functionPrimaryKeyColumns) {
		fields = functionAllColumns
	} else {
		fields = strmangle.SetComplement(
			functionAllColumns,
			functionPrimaryKeyColumns,
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

	slice := FunctionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFunctionsUpsert(t *testing.T) {
	t.Parallel()

	if len(functionAllColumns) == len(functionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Function{}
	if err = randomize.Struct(seed, &o, functionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Function: %s", err)
	}

	count, err := Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, functionDBTypes, false, functionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Function struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Function: %s", err)
	}

	count, err = Functions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

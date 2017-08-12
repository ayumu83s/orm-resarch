// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package sqlboiler

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testStores(t *testing.T) {
	t.Parallel()

	query := Stores(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testStoresDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = store.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStoresQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stores(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStoresSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StoreSlice{store}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testStoresExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := StoreExists(tx, store.StoreID)
	if err != nil {
		t.Errorf("Unable to check if Store exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StoreExistsG to return true, but got false.")
	}
}
func testStoresFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	storeFound, err := FindStore(tx, store.StoreID)
	if err != nil {
		t.Error(err)
	}

	if storeFound == nil {
		t.Error("want a record, got nil")
	}
}
func testStoresBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Stores(tx).Bind(store); err != nil {
		t.Error(err)
	}
}

func testStoresOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Stores(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStoresAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	storeOne := &Store{}
	storeTwo := &Store{}
	if err = randomize.Struct(seed, storeOne, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}
	if err = randomize.Struct(seed, storeTwo, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = storeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = storeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stores(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStoresCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	storeOne := &Store{}
	storeTwo := &Store{}
	if err = randomize.Struct(seed, storeOne, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}
	if err = randomize.Struct(seed, storeTwo, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = storeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = storeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func storeBeforeInsertHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeAfterInsertHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeAfterSelectHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeBeforeUpdateHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeAfterUpdateHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeBeforeDeleteHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeAfterDeleteHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeBeforeUpsertHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func storeAfterUpsertHook(e boil.Executor, o *Store) error {
	*o = Store{}
	return nil
}

func testStoresHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Store{}
	o := &Store{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, storeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Store object: %s", err)
	}

	AddStoreHook(boil.BeforeInsertHook, storeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	storeBeforeInsertHooks = []StoreHook{}

	AddStoreHook(boil.AfterInsertHook, storeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	storeAfterInsertHooks = []StoreHook{}

	AddStoreHook(boil.AfterSelectHook, storeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	storeAfterSelectHooks = []StoreHook{}

	AddStoreHook(boil.BeforeUpdateHook, storeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	storeBeforeUpdateHooks = []StoreHook{}

	AddStoreHook(boil.AfterUpdateHook, storeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	storeAfterUpdateHooks = []StoreHook{}

	AddStoreHook(boil.BeforeDeleteHook, storeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	storeBeforeDeleteHooks = []StoreHook{}

	AddStoreHook(boil.AfterDeleteHook, storeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	storeAfterDeleteHooks = []StoreHook{}

	AddStoreHook(boil.BeforeUpsertHook, storeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	storeBeforeUpsertHooks = []StoreHook{}

	AddStoreHook(boil.AfterUpsertHook, storeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	storeAfterUpsertHooks = []StoreHook{}
}
func testStoresInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStoresInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx, storeColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStoreToManyCustomers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c Customer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, customerDBTypes, false, customerColumnsWithDefault...)
	randomize.Struct(seed, &c, customerDBTypes, false, customerColumnsWithDefault...)

	b.StoreID = a.StoreID
	c.StoreID = a.StoreID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	customer, err := a.Customers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range customer {
		if v.StoreID == b.StoreID {
			bFound = true
		}
		if v.StoreID == c.StoreID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := StoreSlice{&a}
	if err = a.L.LoadCustomers(tx, false, (*[]*Store)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Customers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Customers = nil
	if err = a.L.LoadCustomers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Customers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", customer)
	}
}

func testStoreToManyInventories(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c Inventory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, inventoryDBTypes, false, inventoryColumnsWithDefault...)
	randomize.Struct(seed, &c, inventoryDBTypes, false, inventoryColumnsWithDefault...)

	b.StoreID = a.StoreID
	c.StoreID = a.StoreID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	inventory, err := a.Inventories(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range inventory {
		if v.StoreID == b.StoreID {
			bFound = true
		}
		if v.StoreID == c.StoreID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := StoreSlice{&a}
	if err = a.L.LoadInventories(tx, false, (*[]*Store)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Inventories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Inventories = nil
	if err = a.L.LoadInventories(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Inventories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", inventory)
	}
}

func testStoreToManyStaffs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c Staff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, staffDBTypes, false, staffColumnsWithDefault...)
	randomize.Struct(seed, &c, staffDBTypes, false, staffColumnsWithDefault...)

	b.StoreID = a.StoreID
	c.StoreID = a.StoreID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	staff, err := a.Staffs(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range staff {
		if v.StoreID == b.StoreID {
			bFound = true
		}
		if v.StoreID == c.StoreID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := StoreSlice{&a}
	if err = a.L.LoadStaffs(tx, false, (*[]*Store)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Staffs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Staffs = nil
	if err = a.L.LoadStaffs(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Staffs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", staff)
	}
}

func testStoreToManyAddOpCustomers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c, d, e Customer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Customer{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Customer{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCustomers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.StoreID != first.StoreID {
			t.Error("foreign key was wrong value", a.StoreID, first.StoreID)
		}
		if a.StoreID != second.StoreID {
			t.Error("foreign key was wrong value", a.StoreID, second.StoreID)
		}

		if first.R.Store != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Store != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Customers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Customers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Customers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testStoreToManyAddOpInventories(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c, d, e Inventory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Inventory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, inventoryDBTypes, false, strmangle.SetComplement(inventoryPrimaryKeyColumns, inventoryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Inventory{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddInventories(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.StoreID != first.StoreID {
			t.Error("foreign key was wrong value", a.StoreID, first.StoreID)
		}
		if a.StoreID != second.StoreID {
			t.Error("foreign key was wrong value", a.StoreID, second.StoreID)
		}

		if first.R.Store != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Store != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Inventories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Inventories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Inventories(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testStoreToManyAddOpStaffs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c, d, e Staff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Staff{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, staffDBTypes, false, strmangle.SetComplement(staffPrimaryKeyColumns, staffColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Staff{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStaffs(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.StoreID != first.StoreID {
			t.Error("foreign key was wrong value", a.StoreID, first.StoreID)
		}
		if a.StoreID != second.StoreID {
			t.Error("foreign key was wrong value", a.StoreID, second.StoreID)
		}

		if first.R.Store != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Store != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Staffs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Staffs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Staffs(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testStoreToOneAddressUsingAddress(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Store
	var foreign Address

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AddressID = foreign.AddressID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Address(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AddressID != foreign.AddressID {
		t.Errorf("want: %v, got %v", foreign.AddressID, check.AddressID)
	}

	slice := StoreSlice{&local}
	if err = local.L.LoadAddress(tx, false, (*[]*Store)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Address = nil
	if err = local.L.LoadAddress(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStoreToOneStaffUsingManagerStaff(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Store
	var foreign Staff

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, staffDBTypes, false, staffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Staff struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ManagerStaffID = foreign.StaffID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ManagerStaff(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StaffID != foreign.StaffID {
		t.Errorf("want: %v, got %v", foreign.StaffID, check.StaffID)
	}

	slice := StoreSlice{&local}
	if err = local.L.LoadManagerStaff(tx, false, (*[]*Store)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.ManagerStaff == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ManagerStaff = nil
	if err = local.L.LoadManagerStaff(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ManagerStaff == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStoreToOneSetOpAddressUsingAddress(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Address{&b, &c} {
		err = a.SetAddress(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Address != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stores[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AddressID != x.AddressID {
			t.Error("foreign key was wrong value", a.AddressID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AddressID))
		reflect.Indirect(reflect.ValueOf(&a.AddressID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AddressID != x.AddressID {
			t.Error("foreign key was wrong value", a.AddressID, x.AddressID)
		}
	}
}
func testStoreToOneSetOpStaffUsingManagerStaff(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Store
	var b, c Staff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, staffDBTypes, false, strmangle.SetComplement(staffPrimaryKeyColumns, staffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, staffDBTypes, false, strmangle.SetComplement(staffPrimaryKeyColumns, staffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Staff{&b, &c} {
		err = a.SetManagerStaff(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ManagerStaff != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ManagerStaffStore != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ManagerStaffID != x.StaffID {
			t.Error("foreign key was wrong value", a.ManagerStaffID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ManagerStaffID))
		reflect.Indirect(reflect.ValueOf(&a.ManagerStaffID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ManagerStaffID != x.StaffID {
			t.Error("foreign key was wrong value", a.ManagerStaffID, x.StaffID)
		}
	}
}
func testStoresReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = store.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testStoresReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := StoreSlice{store}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testStoresSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Stores(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	storeDBTypes = map[string]string{`AddressID`: `smallint`, `LastUpdate`: `timestamp`, `ManagerStaffID`: `tinyint`, `StoreID`: `tinyint`}
	_            = bytes.MinRead
)

func testStoresUpdate(t *testing.T) {
	t.Parallel()

	if len(storeColumns) == len(storePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err = store.Update(tx); err != nil {
		t.Error(err)
	}
}

func testStoresSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(storeColumns) == len(storePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	store := &Store{}
	if err = randomize.Struct(seed, store, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, store, storeDBTypes, true, storePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(storeColumns, storePrimaryKeyColumns) {
		fields = storeColumns
	} else {
		fields = strmangle.SetComplement(
			storeColumns,
			storePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(store))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := StoreSlice{store}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testStoresUpsert(t *testing.T) {
	t.Parallel()

	if len(storeColumns) == len(storePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	store := Store{}
	if err = randomize.Struct(seed, &store, storeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = store.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Store: %s", err)
	}

	count, err := Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &store, storeDBTypes, false, storePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err = store.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Store: %s", err)
	}

	count, err = Stores(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

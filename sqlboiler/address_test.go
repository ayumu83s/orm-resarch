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

func testAddresses(t *testing.T) {
	t.Parallel()

	query := Addresses(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAddressesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = address.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAddressesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Addresses(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAddressesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AddressSlice{address}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAddressesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AddressExists(tx, address.AddressID)
	if err != nil {
		t.Errorf("Unable to check if Address exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AddressExistsG to return true, but got false.")
	}
}
func testAddressesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	addressFound, err := FindAddress(tx, address.AddressID)
	if err != nil {
		t.Error(err)
	}

	if addressFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAddressesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Addresses(tx).Bind(address); err != nil {
		t.Error(err)
	}
}

func testAddressesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Addresses(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAddressesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	addressOne := &Address{}
	addressTwo := &Address{}
	if err = randomize.Struct(seed, addressOne, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}
	if err = randomize.Struct(seed, addressTwo, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = addressOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = addressTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Addresses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAddressesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	addressOne := &Address{}
	addressTwo := &Address{}
	if err = randomize.Struct(seed, addressOne, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}
	if err = randomize.Struct(seed, addressTwo, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = addressOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = addressTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func addressBeforeInsertHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressAfterInsertHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressAfterSelectHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressBeforeUpdateHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressAfterUpdateHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressBeforeDeleteHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressAfterDeleteHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressBeforeUpsertHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func addressAfterUpsertHook(e boil.Executor, o *Address) error {
	*o = Address{}
	return nil
}

func testAddressesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Address{}
	o := &Address{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, addressDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Address object: %s", err)
	}

	AddAddressHook(boil.BeforeInsertHook, addressBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	addressBeforeInsertHooks = []AddressHook{}

	AddAddressHook(boil.AfterInsertHook, addressAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	addressAfterInsertHooks = []AddressHook{}

	AddAddressHook(boil.AfterSelectHook, addressAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	addressAfterSelectHooks = []AddressHook{}

	AddAddressHook(boil.BeforeUpdateHook, addressBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	addressBeforeUpdateHooks = []AddressHook{}

	AddAddressHook(boil.AfterUpdateHook, addressAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	addressAfterUpdateHooks = []AddressHook{}

	AddAddressHook(boil.BeforeDeleteHook, addressBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	addressBeforeDeleteHooks = []AddressHook{}

	AddAddressHook(boil.AfterDeleteHook, addressAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	addressAfterDeleteHooks = []AddressHook{}

	AddAddressHook(boil.BeforeUpsertHook, addressBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	addressBeforeUpsertHooks = []AddressHook{}

	AddAddressHook(boil.AfterUpsertHook, addressAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	addressAfterUpsertHooks = []AddressHook{}
}
func testAddressesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAddressesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx, addressColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAddressToManyCustomers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Address
	var b, c Customer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, customerDBTypes, false, customerColumnsWithDefault...)
	randomize.Struct(seed, &c, customerDBTypes, false, customerColumnsWithDefault...)

	b.AddressID = a.AddressID
	c.AddressID = a.AddressID
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
		if v.AddressID == b.AddressID {
			bFound = true
		}
		if v.AddressID == c.AddressID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AddressSlice{&a}
	if err = a.L.LoadCustomers(tx, false, (*[]*Address)(&slice)); err != nil {
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

func testAddressToManyStaffs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Address
	var b, c Staff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, staffDBTypes, false, staffColumnsWithDefault...)
	randomize.Struct(seed, &c, staffDBTypes, false, staffColumnsWithDefault...)

	b.AddressID = a.AddressID
	c.AddressID = a.AddressID
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
		if v.AddressID == b.AddressID {
			bFound = true
		}
		if v.AddressID == c.AddressID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AddressSlice{&a}
	if err = a.L.LoadStaffs(tx, false, (*[]*Address)(&slice)); err != nil {
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

func testAddressToManyStores(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Address
	var b, c Store

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, storeDBTypes, false, storeColumnsWithDefault...)
	randomize.Struct(seed, &c, storeDBTypes, false, storeColumnsWithDefault...)

	b.AddressID = a.AddressID
	c.AddressID = a.AddressID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	store, err := a.Stores(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range store {
		if v.AddressID == b.AddressID {
			bFound = true
		}
		if v.AddressID == c.AddressID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AddressSlice{&a}
	if err = a.L.LoadStores(tx, false, (*[]*Address)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Stores); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Stores = nil
	if err = a.L.LoadStores(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Stores); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", store)
	}
}

func testAddressToManyAddOpCustomers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Address
	var b, c, d, e Customer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
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

		if a.AddressID != first.AddressID {
			t.Error("foreign key was wrong value", a.AddressID, first.AddressID)
		}
		if a.AddressID != second.AddressID {
			t.Error("foreign key was wrong value", a.AddressID, second.AddressID)
		}

		if first.R.Address != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Address != &a {
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
func testAddressToManyAddOpStaffs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Address
	var b, c, d, e Staff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
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

		if a.AddressID != first.AddressID {
			t.Error("foreign key was wrong value", a.AddressID, first.AddressID)
		}
		if a.AddressID != second.AddressID {
			t.Error("foreign key was wrong value", a.AddressID, second.AddressID)
		}

		if first.R.Address != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Address != &a {
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
func testAddressToManyAddOpStores(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Address
	var b, c, d, e Store

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Store{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Store{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddStores(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.AddressID != first.AddressID {
			t.Error("foreign key was wrong value", a.AddressID, first.AddressID)
		}
		if a.AddressID != second.AddressID {
			t.Error("foreign key was wrong value", a.AddressID, second.AddressID)
		}

		if first.R.Address != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Address != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Stores[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Stores[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Stores(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAddressToOneCityUsingCity(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Address
	var foreign City

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cityDBTypes, false, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.CityID = foreign.CityID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.City(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.CityID != foreign.CityID {
		t.Errorf("want: %v, got %v", foreign.CityID, check.CityID)
	}

	slice := AddressSlice{&local}
	if err = local.L.LoadCity(tx, false, (*[]*Address)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.City == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.City = nil
	if err = local.L.LoadCity(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.City == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAddressToOneSetOpCityUsingCity(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Address
	var b, c City

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*City{&b, &c} {
		err = a.SetCity(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.City != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Addresses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CityID != x.CityID {
			t.Error("foreign key was wrong value", a.CityID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CityID))
		reflect.Indirect(reflect.ValueOf(&a.CityID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CityID != x.CityID {
			t.Error("foreign key was wrong value", a.CityID, x.CityID)
		}
	}
}
func testAddressesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = address.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAddressesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AddressSlice{address}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAddressesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Addresses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	addressDBTypes = map[string]string{`Address`: `varchar`, `Address2`: `varchar`, `AddressID`: `smallint`, `CityID`: `smallint`, `District`: `varchar`, `LastUpdate`: `timestamp`, `Phone`: `varchar`, `PostalCode`: `varchar`}
	_              = bytes.MinRead
)

func testAddressesUpdate(t *testing.T) {
	t.Parallel()

	if len(addressColumns) == len(addressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err = address.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAddressesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(addressColumns) == len(addressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	address := &Address{}
	if err = randomize.Struct(seed, address, addressDBTypes, true, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, address, addressDBTypes, true, addressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(addressColumns, addressPrimaryKeyColumns) {
		fields = addressColumns
	} else {
		fields = strmangle.SetComplement(
			addressColumns,
			addressPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(address))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AddressSlice{address}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAddressesUpsert(t *testing.T) {
	t.Parallel()

	if len(addressColumns) == len(addressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	address := Address{}
	if err = randomize.Struct(seed, &address, addressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = address.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Address: %s", err)
	}

	count, err := Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &address, addressDBTypes, false, addressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err = address.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Address: %s", err)
	}

	count, err = Addresses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

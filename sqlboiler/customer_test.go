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

func testCustomers(t *testing.T) {
	t.Parallel()

	query := Customers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCustomersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = customer.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Customers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCustomersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CustomerSlice{customer}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCustomersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CustomerExists(tx, customer.CustomerID)
	if err != nil {
		t.Errorf("Unable to check if Customer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CustomerExistsG to return true, but got false.")
	}
}
func testCustomersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	customerFound, err := FindCustomer(tx, customer.CustomerID)
	if err != nil {
		t.Error(err)
	}

	if customerFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCustomersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Customers(tx).Bind(customer); err != nil {
		t.Error(err)
	}
}

func testCustomersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Customers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCustomersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customerOne := &Customer{}
	customerTwo := &Customer{}
	if err = randomize.Struct(seed, customerOne, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err = randomize.Struct(seed, customerTwo, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customerOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = customerTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Customers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCustomersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	customerOne := &Customer{}
	customerTwo := &Customer{}
	if err = randomize.Struct(seed, customerOne, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err = randomize.Struct(seed, customerTwo, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customerOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = customerTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func customerBeforeInsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterInsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterSelectHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeUpdateHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterUpdateHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeDeleteHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterDeleteHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerBeforeUpsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func customerAfterUpsertHook(e boil.Executor, o *Customer) error {
	*o = Customer{}
	return nil
}

func testCustomersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Customer{}
	o := &Customer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, customerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Customer object: %s", err)
	}

	AddCustomerHook(boil.BeforeInsertHook, customerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	customerBeforeInsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterInsertHook, customerAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	customerAfterInsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterSelectHook, customerAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	customerAfterSelectHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeUpdateHook, customerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	customerBeforeUpdateHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterUpdateHook, customerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	customerAfterUpdateHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeDeleteHook, customerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	customerBeforeDeleteHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterDeleteHook, customerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	customerAfterDeleteHooks = []CustomerHook{}

	AddCustomerHook(boil.BeforeUpsertHook, customerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	customerBeforeUpsertHooks = []CustomerHook{}

	AddCustomerHook(boil.AfterUpsertHook, customerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	customerAfterUpsertHooks = []CustomerHook{}
}
func testCustomersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx, customerColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCustomerToManyPayments(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Customer
	var b, c Payment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, paymentDBTypes, false, paymentColumnsWithDefault...)
	randomize.Struct(seed, &c, paymentDBTypes, false, paymentColumnsWithDefault...)

	b.CustomerID = a.CustomerID
	c.CustomerID = a.CustomerID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	payment, err := a.Payments(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range payment {
		if v.CustomerID == b.CustomerID {
			bFound = true
		}
		if v.CustomerID == c.CustomerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CustomerSlice{&a}
	if err = a.L.LoadPayments(tx, false, (*[]*Customer)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Payments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Payments = nil
	if err = a.L.LoadPayments(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Payments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", payment)
	}
}

func testCustomerToManyRentals(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Customer
	var b, c Rental

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, rentalDBTypes, false, rentalColumnsWithDefault...)
	randomize.Struct(seed, &c, rentalDBTypes, false, rentalColumnsWithDefault...)

	b.CustomerID = a.CustomerID
	c.CustomerID = a.CustomerID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	rental, err := a.Rentals(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range rental {
		if v.CustomerID == b.CustomerID {
			bFound = true
		}
		if v.CustomerID == c.CustomerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CustomerSlice{&a}
	if err = a.L.LoadRentals(tx, false, (*[]*Customer)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Rentals); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Rentals = nil
	if err = a.L.LoadRentals(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Rentals); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", rental)
	}
}

func testCustomerToManyAddOpPayments(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Customer
	var b, c, d, e Payment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Payment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, paymentDBTypes, false, strmangle.SetComplement(paymentPrimaryKeyColumns, paymentColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Payment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPayments(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CustomerID != first.CustomerID {
			t.Error("foreign key was wrong value", a.CustomerID, first.CustomerID)
		}
		if a.CustomerID != second.CustomerID {
			t.Error("foreign key was wrong value", a.CustomerID, second.CustomerID)
		}

		if first.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Payments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Payments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Payments(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCustomerToManyAddOpRentals(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Customer
	var b, c, d, e Rental

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Rental{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rentalDBTypes, false, strmangle.SetComplement(rentalPrimaryKeyColumns, rentalColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Rental{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRentals(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CustomerID != first.CustomerID {
			t.Error("foreign key was wrong value", a.CustomerID, first.CustomerID)
		}
		if a.CustomerID != second.CustomerID {
			t.Error("foreign key was wrong value", a.CustomerID, second.CustomerID)
		}

		if first.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Customer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Rentals[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Rentals[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Rentals(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCustomerToOneAddressUsingAddress(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Customer
	var foreign Address

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
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

	slice := CustomerSlice{&local}
	if err = local.L.LoadAddress(tx, false, (*[]*Customer)(&slice)); err != nil {
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

func testCustomerToOneStoreUsingStore(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Customer
	var foreign Store

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StoreID = foreign.StoreID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Store(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.StoreID != foreign.StoreID {
		t.Errorf("want: %v, got %v", foreign.StoreID, check.StoreID)
	}

	slice := CustomerSlice{&local}
	if err = local.L.LoadStore(tx, false, (*[]*Customer)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Store == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Store = nil
	if err = local.L.LoadStore(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Store == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCustomerToOneSetOpAddressUsingAddress(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Customer
	var b, c Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
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

		if x.R.Customers[0] != &a {
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
func testCustomerToOneSetOpStoreUsingStore(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Customer
	var b, c Store

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Store{&b, &c} {
		err = a.SetStore(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Store != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Customers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StoreID != x.StoreID {
			t.Error("foreign key was wrong value", a.StoreID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StoreID))
		reflect.Indirect(reflect.ValueOf(&a.StoreID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StoreID != x.StoreID {
			t.Error("foreign key was wrong value", a.StoreID, x.StoreID)
		}
	}
}
func testCustomersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = customer.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCustomersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CustomerSlice{customer}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCustomersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Customers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	customerDBTypes = map[string]string{`Active`: `tinyint`, `AddressID`: `smallint`, `CreateDate`: `datetime`, `CustomerID`: `smallint`, `Email`: `varchar`, `FirstName`: `varchar`, `LastName`: `varchar`, `LastUpdate`: `timestamp`, `StoreID`: `tinyint`}
	_               = bytes.MinRead
)

func testCustomersUpdate(t *testing.T) {
	t.Parallel()

	if len(customerColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err = customer.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCustomersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(customerColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	customer := &Customer{}
	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, customer, customerDBTypes, true, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(customerColumns, customerPrimaryKeyColumns) {
		fields = customerColumns
	} else {
		fields = strmangle.SetComplement(
			customerColumns,
			customerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(customer))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CustomerSlice{customer}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCustomersUpsert(t *testing.T) {
	t.Parallel()

	if len(customerColumns) == len(customerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	customer := Customer{}
	if err = randomize.Struct(seed, &customer, customerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = customer.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Customer: %s", err)
	}

	count, err := Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &customer, customerDBTypes, false, customerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err = customer.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Customer: %s", err)
	}

	count, err = Customers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

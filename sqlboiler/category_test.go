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

func testCategories(t *testing.T) {
	t.Parallel()

	query := Categories(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = category.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Categories(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CategorySlice{category}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CategoryExists(tx, category.CategoryID)
	if err != nil {
		t.Errorf("Unable to check if Category exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CategoryExistsG to return true, but got false.")
	}
}
func testCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	categoryFound, err := FindCategory(tx, category.CategoryID)
	if err != nil {
		t.Error(err)
	}

	if categoryFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Categories(tx).Bind(category); err != nil {
		t.Error(err)
	}
}

func testCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Categories(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	categoryOne := &Category{}
	categoryTwo := &Category{}
	if err = randomize.Struct(seed, categoryOne, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTwo, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = categoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = categoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Categories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	categoryOne := &Category{}
	categoryTwo := &Category{}
	if err = randomize.Struct(seed, categoryOne, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}
	if err = randomize.Struct(seed, categoryTwo, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = categoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = categoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func categoryBeforeInsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterInsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterSelectHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeUpdateHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterUpdateHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeDeleteHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterDeleteHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryBeforeUpsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func categoryAfterUpsertHook(e boil.Executor, o *Category) error {
	*o = Category{}
	return nil
}

func testCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Category{}
	o := &Category{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, categoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Category object: %s", err)
	}

	AddCategoryHook(boil.BeforeInsertHook, categoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	categoryBeforeInsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterInsertHook, categoryAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	categoryAfterInsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterSelectHook, categoryAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	categoryAfterSelectHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeUpdateHook, categoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	categoryBeforeUpdateHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterUpdateHook, categoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	categoryAfterUpdateHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeDeleteHook, categoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	categoryBeforeDeleteHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterDeleteHook, categoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	categoryAfterDeleteHooks = []CategoryHook{}

	AddCategoryHook(boil.BeforeUpsertHook, categoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	categoryBeforeUpsertHooks = []CategoryHook{}

	AddCategoryHook(boil.AfterUpsertHook, categoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	categoryAfterUpsertHooks = []CategoryHook{}
}
func testCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx, categoryColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCategoryToManyFilmCategories(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Category
	var b, c FilmCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, filmCategoryDBTypes, false, filmCategoryColumnsWithDefault...)
	randomize.Struct(seed, &c, filmCategoryDBTypes, false, filmCategoryColumnsWithDefault...)

	b.CategoryID = a.CategoryID
	c.CategoryID = a.CategoryID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	filmCategory, err := a.FilmCategories(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range filmCategory {
		if v.CategoryID == b.CategoryID {
			bFound = true
		}
		if v.CategoryID == c.CategoryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CategorySlice{&a}
	if err = a.L.LoadFilmCategories(tx, false, (*[]*Category)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FilmCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FilmCategories = nil
	if err = a.L.LoadFilmCategories(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FilmCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", filmCategory)
	}
}

func testCategoryToManyAddOpFilmCategories(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Category
	var b, c, d, e FilmCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*FilmCategory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmCategoryDBTypes, false, strmangle.SetComplement(filmCategoryPrimaryKeyColumns, filmCategoryColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*FilmCategory{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFilmCategories(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.CategoryID != first.CategoryID {
			t.Error("foreign key was wrong value", a.CategoryID, first.CategoryID)
		}
		if a.CategoryID != second.CategoryID {
			t.Error("foreign key was wrong value", a.CategoryID, second.CategoryID)
		}

		if first.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FilmCategories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FilmCategories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FilmCategories(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = category.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CategorySlice{category}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Categories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	categoryDBTypes = map[string]string{`CategoryID`: `tinyint`, `LastUpdate`: `timestamp`, `Name`: `varchar`}
	_               = bytes.MinRead
)

func testCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if len(categoryColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err = category.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(categoryColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	category := &Category{}
	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, category, categoryDBTypes, true, categoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(categoryColumns, categoryPrimaryKeyColumns) {
		fields = categoryColumns
	} else {
		fields = strmangle.SetComplement(
			categoryColumns,
			categoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(category))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CategorySlice{category}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(categoryColumns) == len(categoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	category := Category{}
	if err = randomize.Struct(seed, &category, categoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = category.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Category: %s", err)
	}

	count, err := Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &category, categoryDBTypes, false, categoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err = category.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Category: %s", err)
	}

	count, err = Categories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

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

func testFilms(t *testing.T) {
	t.Parallel()

	query := Films(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFilmsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = film.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilmsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Films(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilmsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FilmSlice{film}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFilmsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FilmExists(tx, film.FilmID)
	if err != nil {
		t.Errorf("Unable to check if Film exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FilmExistsG to return true, but got false.")
	}
}
func testFilmsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	filmFound, err := FindFilm(tx, film.FilmID)
	if err != nil {
		t.Error(err)
	}

	if filmFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFilmsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Films(tx).Bind(film); err != nil {
		t.Error(err)
	}
}

func testFilmsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Films(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFilmsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	filmOne := &Film{}
	filmTwo := &Film{}
	if err = randomize.Struct(seed, filmOne, filmDBTypes, false, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}
	if err = randomize.Struct(seed, filmTwo, filmDBTypes, false, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = filmOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = filmTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Films(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFilmsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	filmOne := &Film{}
	filmTwo := &Film{}
	if err = randomize.Struct(seed, filmOne, filmDBTypes, false, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}
	if err = randomize.Struct(seed, filmTwo, filmDBTypes, false, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = filmOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = filmTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func filmBeforeInsertHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmAfterInsertHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmAfterSelectHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmBeforeUpdateHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmAfterUpdateHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmBeforeDeleteHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmAfterDeleteHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmBeforeUpsertHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func filmAfterUpsertHook(e boil.Executor, o *Film) error {
	*o = Film{}
	return nil
}

func testFilmsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Film{}
	o := &Film{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, filmDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Film object: %s", err)
	}

	AddFilmHook(boil.BeforeInsertHook, filmBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	filmBeforeInsertHooks = []FilmHook{}

	AddFilmHook(boil.AfterInsertHook, filmAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	filmAfterInsertHooks = []FilmHook{}

	AddFilmHook(boil.AfterSelectHook, filmAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	filmAfterSelectHooks = []FilmHook{}

	AddFilmHook(boil.BeforeUpdateHook, filmBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	filmBeforeUpdateHooks = []FilmHook{}

	AddFilmHook(boil.AfterUpdateHook, filmAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	filmAfterUpdateHooks = []FilmHook{}

	AddFilmHook(boil.BeforeDeleteHook, filmBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	filmBeforeDeleteHooks = []FilmHook{}

	AddFilmHook(boil.AfterDeleteHook, filmAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	filmAfterDeleteHooks = []FilmHook{}

	AddFilmHook(boil.BeforeUpsertHook, filmBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	filmBeforeUpsertHooks = []FilmHook{}

	AddFilmHook(boil.AfterUpsertHook, filmAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	filmAfterUpsertHooks = []FilmHook{}
}
func testFilmsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFilmsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx, filmColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFilmToManyFilmActors(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c FilmActor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, filmActorDBTypes, false, filmActorColumnsWithDefault...)
	randomize.Struct(seed, &c, filmActorDBTypes, false, filmActorColumnsWithDefault...)

	b.FilmID = a.FilmID
	c.FilmID = a.FilmID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	filmActor, err := a.FilmActors(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range filmActor {
		if v.FilmID == b.FilmID {
			bFound = true
		}
		if v.FilmID == c.FilmID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FilmSlice{&a}
	if err = a.L.LoadFilmActors(tx, false, (*[]*Film)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FilmActors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FilmActors = nil
	if err = a.L.LoadFilmActors(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FilmActors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", filmActor)
	}
}

func testFilmToManyFilmCategories(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c FilmCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, filmCategoryDBTypes, false, filmCategoryColumnsWithDefault...)
	randomize.Struct(seed, &c, filmCategoryDBTypes, false, filmCategoryColumnsWithDefault...)

	b.FilmID = a.FilmID
	c.FilmID = a.FilmID
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
		if v.FilmID == b.FilmID {
			bFound = true
		}
		if v.FilmID == c.FilmID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FilmSlice{&a}
	if err = a.L.LoadFilmCategories(tx, false, (*[]*Film)(&slice)); err != nil {
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

func testFilmToManyInventories(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c Inventory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, inventoryDBTypes, false, inventoryColumnsWithDefault...)
	randomize.Struct(seed, &c, inventoryDBTypes, false, inventoryColumnsWithDefault...)

	b.FilmID = a.FilmID
	c.FilmID = a.FilmID
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
		if v.FilmID == b.FilmID {
			bFound = true
		}
		if v.FilmID == c.FilmID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FilmSlice{&a}
	if err = a.L.LoadInventories(tx, false, (*[]*Film)(&slice)); err != nil {
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

func testFilmToManyAddOpFilmActors(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c, d, e FilmActor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*FilmActor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmActorDBTypes, false, strmangle.SetComplement(filmActorPrimaryKeyColumns, filmActorColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*FilmActor{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFilmActors(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.FilmID != first.FilmID {
			t.Error("foreign key was wrong value", a.FilmID, first.FilmID)
		}
		if a.FilmID != second.FilmID {
			t.Error("foreign key was wrong value", a.FilmID, second.FilmID)
		}

		if first.R.Film != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Film != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FilmActors[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FilmActors[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FilmActors(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testFilmToManyAddOpFilmCategories(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c, d, e FilmCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
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

		if a.FilmID != first.FilmID {
			t.Error("foreign key was wrong value", a.FilmID, first.FilmID)
		}
		if a.FilmID != second.FilmID {
			t.Error("foreign key was wrong value", a.FilmID, second.FilmID)
		}

		if first.R.Film != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Film != &a {
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
func testFilmToManyAddOpInventories(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c, d, e Inventory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
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

		if a.FilmID != first.FilmID {
			t.Error("foreign key was wrong value", a.FilmID, first.FilmID)
		}
		if a.FilmID != second.FilmID {
			t.Error("foreign key was wrong value", a.FilmID, second.FilmID)
		}

		if first.R.Film != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Film != &a {
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
func testFilmToOneLanguageUsingLanguage(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Film
	var foreign Language

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, filmDBTypes, false, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, languageDBTypes, false, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.LanguageID = foreign.LanguageID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Language(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.LanguageID != foreign.LanguageID {
		t.Errorf("want: %v, got %v", foreign.LanguageID, check.LanguageID)
	}

	slice := FilmSlice{&local}
	if err = local.L.LoadLanguage(tx, false, (*[]*Film)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Language == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Language = nil
	if err = local.L.LoadLanguage(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Language == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFilmToOneLanguageUsingOriginalLanguage(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Film
	var foreign Language

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, languageDBTypes, false, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	local.OriginalLanguageID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.OriginalLanguageID.Uint8 = foreign.LanguageID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.OriginalLanguage(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.LanguageID != foreign.LanguageID {
		t.Errorf("want: %v, got %v", foreign.LanguageID, check.LanguageID)
	}

	slice := FilmSlice{&local}
	if err = local.L.LoadOriginalLanguage(tx, false, (*[]*Film)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.OriginalLanguage == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OriginalLanguage = nil
	if err = local.L.LoadOriginalLanguage(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.OriginalLanguage == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFilmToOneSetOpLanguageUsingLanguage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c Language

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Language{&b, &c} {
		err = a.SetLanguage(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Language != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Films[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.LanguageID != x.LanguageID {
			t.Error("foreign key was wrong value", a.LanguageID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LanguageID))
		reflect.Indirect(reflect.ValueOf(&a.LanguageID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.LanguageID != x.LanguageID {
			t.Error("foreign key was wrong value", a.LanguageID, x.LanguageID)
		}
	}
}
func testFilmToOneSetOpLanguageUsingOriginalLanguage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b, c Language

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Language{&b, &c} {
		err = a.SetOriginalLanguage(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OriginalLanguage != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OriginalLanguageFilms[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OriginalLanguageID.Uint8 != x.LanguageID {
			t.Error("foreign key was wrong value", a.OriginalLanguageID.Uint8)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OriginalLanguageID.Uint8))
		reflect.Indirect(reflect.ValueOf(&a.OriginalLanguageID.Uint8)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OriginalLanguageID.Uint8 != x.LanguageID {
			t.Error("foreign key was wrong value", a.OriginalLanguageID.Uint8, x.LanguageID)
		}
	}
}

func testFilmToOneRemoveOpLanguageUsingOriginalLanguage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Film
	var b Language

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOriginalLanguage(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOriginalLanguage(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.OriginalLanguage(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.OriginalLanguage != nil {
		t.Error("R struct entry should be nil")
	}

	if a.OriginalLanguageID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.OriginalLanguageFilms) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testFilmsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = film.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFilmsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FilmSlice{film}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFilmsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Films(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	filmDBTypes = map[string]string{`Description`: `text`, `FilmID`: `smallint`, `LanguageID`: `tinyint`, `LastUpdate`: `timestamp`, `Length`: `smallint`, `OriginalLanguageID`: `tinyint`, `Rating`: `enum('G','PG','PG-13','R','NC-17')`, `ReleaseYear`: `varchar`, `RentalDuration`: `tinyint`, `RentalRate`: `decimal`, `ReplacementCost`: `decimal`, `SpecialFeatures`: `set`, `Title`: `varchar`}
	_           = bytes.MinRead
)

func testFilmsUpdate(t *testing.T) {
	t.Parallel()

	if len(filmColumns) == len(filmPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	if err = film.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFilmsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(filmColumns) == len(filmPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	film := &Film{}
	if err = randomize.Struct(seed, film, filmDBTypes, true, filmColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, film, filmDBTypes, true, filmPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(filmColumns, filmPrimaryKeyColumns) {
		fields = filmColumns
	} else {
		fields = strmangle.SetComplement(
			filmColumns,
			filmPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(film))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FilmSlice{film}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFilmsUpsert(t *testing.T) {
	t.Parallel()

	if len(filmColumns) == len(filmPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	film := Film{}
	if err = randomize.Struct(seed, &film, filmDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = film.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Film: %s", err)
	}

	count, err := Films(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &film, filmDBTypes, false, filmPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Film struct: %s", err)
	}

	if err = film.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Film: %s", err)
	}

	count, err = Films(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

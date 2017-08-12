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

func testLanguages(t *testing.T) {
	t.Parallel()

	query := Languages(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testLanguagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = language.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLanguagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Languages(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLanguagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LanguageSlice{language}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testLanguagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := LanguageExists(tx, language.LanguageID)
	if err != nil {
		t.Errorf("Unable to check if Language exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LanguageExistsG to return true, but got false.")
	}
}
func testLanguagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	languageFound, err := FindLanguage(tx, language.LanguageID)
	if err != nil {
		t.Error(err)
	}

	if languageFound == nil {
		t.Error("want a record, got nil")
	}
}
func testLanguagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Languages(tx).Bind(language); err != nil {
		t.Error(err)
	}
}

func testLanguagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Languages(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLanguagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	languageOne := &Language{}
	languageTwo := &Language{}
	if err = randomize.Struct(seed, languageOne, languageDBTypes, false, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}
	if err = randomize.Struct(seed, languageTwo, languageDBTypes, false, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = languageOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = languageTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Languages(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLanguagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	languageOne := &Language{}
	languageTwo := &Language{}
	if err = randomize.Struct(seed, languageOne, languageDBTypes, false, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}
	if err = randomize.Struct(seed, languageTwo, languageDBTypes, false, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = languageOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = languageTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func languageBeforeInsertHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageAfterInsertHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageAfterSelectHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageBeforeUpdateHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageAfterUpdateHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageBeforeDeleteHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageAfterDeleteHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageBeforeUpsertHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func languageAfterUpsertHook(e boil.Executor, o *Language) error {
	*o = Language{}
	return nil
}

func testLanguagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Language{}
	o := &Language{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, languageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Language object: %s", err)
	}

	AddLanguageHook(boil.BeforeInsertHook, languageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	languageBeforeInsertHooks = []LanguageHook{}

	AddLanguageHook(boil.AfterInsertHook, languageAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	languageAfterInsertHooks = []LanguageHook{}

	AddLanguageHook(boil.AfterSelectHook, languageAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	languageAfterSelectHooks = []LanguageHook{}

	AddLanguageHook(boil.BeforeUpdateHook, languageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	languageBeforeUpdateHooks = []LanguageHook{}

	AddLanguageHook(boil.AfterUpdateHook, languageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	languageAfterUpdateHooks = []LanguageHook{}

	AddLanguageHook(boil.BeforeDeleteHook, languageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	languageBeforeDeleteHooks = []LanguageHook{}

	AddLanguageHook(boil.AfterDeleteHook, languageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	languageAfterDeleteHooks = []LanguageHook{}

	AddLanguageHook(boil.BeforeUpsertHook, languageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	languageBeforeUpsertHooks = []LanguageHook{}

	AddLanguageHook(boil.AfterUpsertHook, languageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	languageAfterUpsertHooks = []LanguageHook{}
}
func testLanguagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLanguagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx, languageColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLanguageToManyFilms(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Language
	var b, c Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, filmDBTypes, false, filmColumnsWithDefault...)
	randomize.Struct(seed, &c, filmDBTypes, false, filmColumnsWithDefault...)

	b.LanguageID = a.LanguageID
	c.LanguageID = a.LanguageID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	film, err := a.Films(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range film {
		if v.LanguageID == b.LanguageID {
			bFound = true
		}
		if v.LanguageID == c.LanguageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LanguageSlice{&a}
	if err = a.L.LoadFilms(tx, false, (*[]*Language)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Films); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Films = nil
	if err = a.L.LoadFilms(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Films); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", film)
	}
}

func testLanguageToManyOriginalLanguageFilms(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Language
	var b, c Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, filmDBTypes, false, filmColumnsWithDefault...)
	randomize.Struct(seed, &c, filmDBTypes, false, filmColumnsWithDefault...)

	b.OriginalLanguageID.Valid = true
	c.OriginalLanguageID.Valid = true
	b.OriginalLanguageID.Uint8 = a.LanguageID
	c.OriginalLanguageID.Uint8 = a.LanguageID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	film, err := a.OriginalLanguageFilms(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range film {
		if v.OriginalLanguageID.Uint8 == b.OriginalLanguageID.Uint8 {
			bFound = true
		}
		if v.OriginalLanguageID.Uint8 == c.OriginalLanguageID.Uint8 {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LanguageSlice{&a}
	if err = a.L.LoadOriginalLanguageFilms(tx, false, (*[]*Language)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OriginalLanguageFilms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OriginalLanguageFilms = nil
	if err = a.L.LoadOriginalLanguageFilms(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OriginalLanguageFilms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", film)
	}
}

func testLanguageToManyAddOpFilms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Language
	var b, c, d, e Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Film{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Film{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFilms(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.LanguageID != first.LanguageID {
			t.Error("foreign key was wrong value", a.LanguageID, first.LanguageID)
		}
		if a.LanguageID != second.LanguageID {
			t.Error("foreign key was wrong value", a.LanguageID, second.LanguageID)
		}

		if first.R.Language != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Language != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Films[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Films[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Films(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testLanguageToManyAddOpOriginalLanguageFilms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Language
	var b, c, d, e Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Film{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Film{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOriginalLanguageFilms(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.LanguageID != first.OriginalLanguageID.Uint8 {
			t.Error("foreign key was wrong value", a.LanguageID, first.OriginalLanguageID.Uint8)
		}
		if a.LanguageID != second.OriginalLanguageID.Uint8 {
			t.Error("foreign key was wrong value", a.LanguageID, second.OriginalLanguageID.Uint8)
		}

		if first.R.OriginalLanguage != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.OriginalLanguage != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OriginalLanguageFilms[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OriginalLanguageFilms[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OriginalLanguageFilms(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testLanguageToManySetOpOriginalLanguageFilms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Language
	var b, c, d, e Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Film{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetOriginalLanguageFilms(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.OriginalLanguageFilms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetOriginalLanguageFilms(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.OriginalLanguageFilms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.OriginalLanguageID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.OriginalLanguageID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.LanguageID != d.OriginalLanguageID.Uint8 {
		t.Error("foreign key was wrong value", a.LanguageID, d.OriginalLanguageID.Uint8)
	}
	if a.LanguageID != e.OriginalLanguageID.Uint8 {
		t.Error("foreign key was wrong value", a.LanguageID, e.OriginalLanguageID.Uint8)
	}

	if b.R.OriginalLanguage != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.OriginalLanguage != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.OriginalLanguage != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.OriginalLanguage != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.OriginalLanguageFilms[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.OriginalLanguageFilms[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testLanguageToManyRemoveOpOriginalLanguageFilms(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Language
	var b, c, d, e Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, languageDBTypes, false, strmangle.SetComplement(languagePrimaryKeyColumns, languageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Film{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddOriginalLanguageFilms(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.OriginalLanguageFilms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveOriginalLanguageFilms(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.OriginalLanguageFilms(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.OriginalLanguageID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.OriginalLanguageID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.OriginalLanguage != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.OriginalLanguage != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.OriginalLanguage != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.OriginalLanguage != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.OriginalLanguageFilms) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.OriginalLanguageFilms[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.OriginalLanguageFilms[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testLanguagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = language.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testLanguagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LanguageSlice{language}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testLanguagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Languages(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	languageDBTypes = map[string]string{`LanguageID`: `tinyint`, `LastUpdate`: `timestamp`, `Name`: `char`}
	_               = bytes.MinRead
)

func testLanguagesUpdate(t *testing.T) {
	t.Parallel()

	if len(languageColumns) == len(languagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	if err = language.Update(tx); err != nil {
		t.Error(err)
	}
}

func testLanguagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(languageColumns) == len(languagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	language := &Language{}
	if err = randomize.Struct(seed, language, languageDBTypes, true, languageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, language, languageDBTypes, true, languagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(languageColumns, languagePrimaryKeyColumns) {
		fields = languageColumns
	} else {
		fields = strmangle.SetComplement(
			languageColumns,
			languagePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(language))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := LanguageSlice{language}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testLanguagesUpsert(t *testing.T) {
	t.Parallel()

	if len(languageColumns) == len(languagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	language := Language{}
	if err = randomize.Struct(seed, &language, languageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = language.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Language: %s", err)
	}

	count, err := Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &language, languageDBTypes, false, languagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Language struct: %s", err)
	}

	if err = language.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Language: %s", err)
	}

	count, err = Languages(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

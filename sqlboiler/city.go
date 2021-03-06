// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package sqlboiler

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// City is an object representing the database table.
type City struct {
	CityID     uint16    `boil:"city_id" json:"city_id" toml:"city_id" yaml:"city_id"`
	City       string    `boil:"city" json:"city" toml:"city" yaml:"city"`
	CountryID  uint16    `boil:"country_id" json:"country_id" toml:"country_id" yaml:"country_id"`
	LastUpdate time.Time `boil:"last_update" json:"last_update" toml:"last_update" yaml:"last_update"`

	R *cityR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cityL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CityColumns = struct {
	CityID     string
	City       string
	CountryID  string
	LastUpdate string
}{
	CityID:     "city_id",
	City:       "city",
	CountryID:  "country_id",
	LastUpdate: "last_update",
}

// cityR is where relationships are stored.
type cityR struct {
	Country   *Country
	Addresses AddressSlice
}

// cityL is where Load methods for each relationship are stored.
type cityL struct{}

var (
	cityColumns               = []string{"city_id", "city", "country_id", "last_update"}
	cityColumnsWithoutDefault = []string{"city", "country_id"}
	cityColumnsWithDefault    = []string{"city_id", "last_update"}
	cityPrimaryKeyColumns     = []string{"city_id"}
)

type (
	// CitySlice is an alias for a slice of pointers to City.
	// This should generally be used opposed to []City.
	CitySlice []*City
	// CityHook is the signature for custom City hook methods
	CityHook func(boil.Executor, *City) error

	cityQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cityType                 = reflect.TypeOf(&City{})
	cityMapping              = queries.MakeStructMapping(cityType)
	cityPrimaryKeyMapping, _ = queries.BindMapping(cityType, cityMapping, cityPrimaryKeyColumns)
	cityInsertCacheMut       sync.RWMutex
	cityInsertCache          = make(map[string]insertCache)
	cityUpdateCacheMut       sync.RWMutex
	cityUpdateCache          = make(map[string]updateCache)
	cityUpsertCacheMut       sync.RWMutex
	cityUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cityBeforeInsertHooks []CityHook
var cityBeforeUpdateHooks []CityHook
var cityBeforeDeleteHooks []CityHook
var cityBeforeUpsertHooks []CityHook

var cityAfterInsertHooks []CityHook
var cityAfterSelectHooks []CityHook
var cityAfterUpdateHooks []CityHook
var cityAfterDeleteHooks []CityHook
var cityAfterUpsertHooks []CityHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *City) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cityBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *City) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cityBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *City) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cityBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *City) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cityBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *City) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cityAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *City) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cityAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *City) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cityAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *City) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cityAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *City) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cityAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCityHook registers your hook function for all future operations.
func AddCityHook(hookPoint boil.HookPoint, cityHook CityHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cityBeforeInsertHooks = append(cityBeforeInsertHooks, cityHook)
	case boil.BeforeUpdateHook:
		cityBeforeUpdateHooks = append(cityBeforeUpdateHooks, cityHook)
	case boil.BeforeDeleteHook:
		cityBeforeDeleteHooks = append(cityBeforeDeleteHooks, cityHook)
	case boil.BeforeUpsertHook:
		cityBeforeUpsertHooks = append(cityBeforeUpsertHooks, cityHook)
	case boil.AfterInsertHook:
		cityAfterInsertHooks = append(cityAfterInsertHooks, cityHook)
	case boil.AfterSelectHook:
		cityAfterSelectHooks = append(cityAfterSelectHooks, cityHook)
	case boil.AfterUpdateHook:
		cityAfterUpdateHooks = append(cityAfterUpdateHooks, cityHook)
	case boil.AfterDeleteHook:
		cityAfterDeleteHooks = append(cityAfterDeleteHooks, cityHook)
	case boil.AfterUpsertHook:
		cityAfterUpsertHooks = append(cityAfterUpsertHooks, cityHook)
	}
}

// OneP returns a single city record from the query, and panics on error.
func (q cityQuery) OneP() *City {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single city record from the query.
func (q cityQuery) One() (*City, error) {
	o := &City{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: failed to execute a one query for city")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all City records from the query, and panics on error.
func (q cityQuery) AllP() CitySlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all City records from the query.
func (q cityQuery) All() (CitySlice, error) {
	var o []*City

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler: failed to assign all query results to City slice")
	}

	if len(cityAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all City records in the query, and panics on error.
func (q cityQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all City records in the query.
func (q cityQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to count city rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cityQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cityQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: failed to check if city exists")
	}

	return count > 0, nil
}

// CountryG pointed to by the foreign key.
func (o *City) CountryG(mods ...qm.QueryMod) countryQuery {
	return o.Country(boil.GetDB(), mods...)
}

// Country pointed to by the foreign key.
func (o *City) Country(exec boil.Executor, mods ...qm.QueryMod) countryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("country_id=?", o.CountryID),
	}

	queryMods = append(queryMods, mods...)

	query := Countries(exec, queryMods...)
	queries.SetFrom(query.Query, "`country`")

	return query
}

// AddressesG retrieves all the address's address.
func (o *City) AddressesG(mods ...qm.QueryMod) addressQuery {
	return o.Addresses(boil.GetDB(), mods...)
}

// Addresses retrieves all the address's address with an executor.
func (o *City) Addresses(exec boil.Executor, mods ...qm.QueryMod) addressQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`address`.`city_id`=?", o.CityID),
	)

	query := Addresses(exec, queryMods...)
	queries.SetFrom(query.Query, "`address`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`address`.*"})
	}

	return query
}

// LoadCountry allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cityL) LoadCountry(e boil.Executor, singular bool, maybeCity interface{}) error {
	var slice []*City
	var object *City

	count := 1
	if singular {
		object = maybeCity.(*City)
	} else {
		slice = *maybeCity.(*[]*City)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &cityR{}
		}
		args[0] = object.CountryID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &cityR{}
			}
			args[i] = obj.CountryID
		}
	}

	query := fmt.Sprintf(
		"select * from `country` where `country_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Country")
	}
	defer results.Close()

	var resultSlice []*Country
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Country")
	}

	if len(cityAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Country = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CountryID == foreign.CountryID {
				local.R.Country = foreign
				break
			}
		}
	}

	return nil
}

// LoadAddresses allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cityL) LoadAddresses(e boil.Executor, singular bool, maybeCity interface{}) error {
	var slice []*City
	var object *City

	count := 1
	if singular {
		object = maybeCity.(*City)
	} else {
		slice = *maybeCity.(*[]*City)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &cityR{}
		}
		args[0] = object.CityID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &cityR{}
			}
			args[i] = obj.CityID
		}
	}

	query := fmt.Sprintf(
		"select * from `address` where `city_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load address")
	}
	defer results.Close()

	var resultSlice []*Address
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice address")
	}

	if len(addressAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Addresses = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CityID == foreign.CityID {
				local.R.Addresses = append(local.R.Addresses, foreign)
				break
			}
		}
	}

	return nil
}

// SetCountryG of the city to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.Cities.
// Uses the global database handle.
func (o *City) SetCountryG(insert bool, related *Country) error {
	return o.SetCountry(boil.GetDB(), insert, related)
}

// SetCountryP of the city to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.Cities.
// Panics on error.
func (o *City) SetCountryP(exec boil.Executor, insert bool, related *Country) {
	if err := o.SetCountry(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetCountryGP of the city to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.Cities.
// Uses the global database handle and panics on error.
func (o *City) SetCountryGP(insert bool, related *Country) {
	if err := o.SetCountry(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetCountry of the city to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.Cities.
func (o *City) SetCountry(exec boil.Executor, insert bool, related *Country) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `city` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"country_id"}),
		strmangle.WhereClause("`", "`", 0, cityPrimaryKeyColumns),
	)
	values := []interface{}{related.CountryID, o.CityID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CountryID = related.CountryID

	if o.R == nil {
		o.R = &cityR{
			Country: related,
		}
	} else {
		o.R.Country = related
	}

	if related.R == nil {
		related.R = &countryR{
			Cities: CitySlice{o},
		}
	} else {
		related.R.Cities = append(related.R.Cities, o)
	}

	return nil
}

// AddAddressesG adds the given related objects to the existing relationships
// of the city, optionally inserting them as new records.
// Appends related to o.R.Addresses.
// Sets related.R.City appropriately.
// Uses the global database handle.
func (o *City) AddAddressesG(insert bool, related ...*Address) error {
	return o.AddAddresses(boil.GetDB(), insert, related...)
}

// AddAddressesP adds the given related objects to the existing relationships
// of the city, optionally inserting them as new records.
// Appends related to o.R.Addresses.
// Sets related.R.City appropriately.
// Panics on error.
func (o *City) AddAddressesP(exec boil.Executor, insert bool, related ...*Address) {
	if err := o.AddAddresses(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddAddressesGP adds the given related objects to the existing relationships
// of the city, optionally inserting them as new records.
// Appends related to o.R.Addresses.
// Sets related.R.City appropriately.
// Uses the global database handle and panics on error.
func (o *City) AddAddressesGP(insert bool, related ...*Address) {
	if err := o.AddAddresses(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddAddresses adds the given related objects to the existing relationships
// of the city, optionally inserting them as new records.
// Appends related to o.R.Addresses.
// Sets related.R.City appropriately.
func (o *City) AddAddresses(exec boil.Executor, insert bool, related ...*Address) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CityID = o.CityID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `address` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"city_id"}),
				strmangle.WhereClause("`", "`", 0, addressPrimaryKeyColumns),
			)
			values := []interface{}{o.CityID, rel.AddressID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CityID = o.CityID
		}
	}

	if o.R == nil {
		o.R = &cityR{
			Addresses: related,
		}
	} else {
		o.R.Addresses = append(o.R.Addresses, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &addressR{
				City: o,
			}
		} else {
			rel.R.City = o
		}
	}
	return nil
}

// CitiesG retrieves all records.
func CitiesG(mods ...qm.QueryMod) cityQuery {
	return Cities(boil.GetDB(), mods...)
}

// Cities retrieves all the records using an executor.
func Cities(exec boil.Executor, mods ...qm.QueryMod) cityQuery {
	mods = append(mods, qm.From("`city`"))
	return cityQuery{NewQuery(exec, mods...)}
}

// FindCityG retrieves a single record by ID.
func FindCityG(cityID uint16, selectCols ...string) (*City, error) {
	return FindCity(boil.GetDB(), cityID, selectCols...)
}

// FindCityGP retrieves a single record by ID, and panics on error.
func FindCityGP(cityID uint16, selectCols ...string) *City {
	retobj, err := FindCity(boil.GetDB(), cityID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCity retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCity(exec boil.Executor, cityID uint16, selectCols ...string) (*City, error) {
	cityObj := &City{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `city` where `city_id`=?", sel,
	)

	q := queries.Raw(exec, query, cityID)

	err := q.Bind(cityObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: unable to select from city")
	}

	return cityObj, nil
}

// FindCityP retrieves a single record by ID with an executor, and panics on error.
func FindCityP(exec boil.Executor, cityID uint16, selectCols ...string) *City {
	retobj, err := FindCity(exec, cityID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *City) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *City) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *City) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *City) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("sqlboiler: no city provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cityColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cityInsertCacheMut.RLock()
	cache, cached := cityInsertCache[key]
	cityInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cityColumns,
			cityColumnsWithDefault,
			cityColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cityType, cityMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cityType, cityMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `city` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `city` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `city` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cityPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to insert into city")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.CityID = uint16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cityMapping["CityID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CityID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to populate default values for city")
	}

CacheNoHooks:
	if !cached {
		cityInsertCacheMut.Lock()
		cityInsertCache[key] = cache
		cityInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single City record. See Update for
// whitelist behavior description.
func (o *City) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single City record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *City) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the City, and panics on error.
// See Update for whitelist behavior description.
func (o *City) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the City.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *City) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cityUpdateCacheMut.RLock()
	cache, cached := cityUpdateCache[key]
	cityUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			cityColumns,
			cityPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("sqlboiler: unable to update city, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `city` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cityPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cityType, cityMapping, append(wl, cityPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to update city row")
	}

	if !cached {
		cityUpdateCacheMut.Lock()
		cityUpdateCache[key] = cache
		cityUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cityQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cityQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to update all for city")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CitySlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CitySlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CitySlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CitySlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("sqlboiler: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `city` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cityPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to update all in city slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *City) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *City) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *City) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *City) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("sqlboiler: no city provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cityColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	cityUpsertCacheMut.RLock()
	cache, cached := cityUpsertCache[key]
	cityUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			cityColumns,
			cityColumnsWithDefault,
			cityColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			cityColumns,
			cityPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("sqlboiler: unable to upsert city, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "city", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `city` WHERE `city_id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(cityType, cityMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cityType, cityMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to upsert for city")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.CityID = uint16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cityMapping["CityID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CityID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to populate default values for city")
	}

CacheNoHooks:
	if !cached {
		cityUpsertCacheMut.Lock()
		cityUpsertCache[key] = cache
		cityUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single City record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *City) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single City record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *City) DeleteG() error {
	if o == nil {
		return errors.New("sqlboiler: no City provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single City record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *City) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single City record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *City) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("sqlboiler: no City provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cityPrimaryKeyMapping)
	sql := "DELETE FROM `city` WHERE `city_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to delete from city")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cityQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cityQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("sqlboiler: no cityQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to delete all from city")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CitySlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CitySlice) DeleteAllG() error {
	if o == nil {
		return errors.New("sqlboiler: no City slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CitySlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CitySlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("sqlboiler: no City slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cityBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `city` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cityPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to delete all from city slice")
	}

	if len(cityAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *City) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *City) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *City) ReloadG() error {
	if o == nil {
		return errors.New("sqlboiler: no City provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *City) Reload(exec boil.Executor) error {
	ret, err := FindCity(exec, o.CityID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CitySlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CitySlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CitySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("sqlboiler: empty CitySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CitySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cities := CitySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `city`.* FROM `city` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cityPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cities)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to reload all in CitySlice")
	}

	*o = cities

	return nil
}

// CityExists checks if the City row exists.
func CityExists(exec boil.Executor, cityID uint16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `city` where `city_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cityID)
	}

	row := exec.QueryRow(sql, cityID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: unable to check if city exists")
	}

	return exists, nil
}

// CityExistsG checks if the City row exists.
func CityExistsG(cityID uint16) (bool, error) {
	return CityExists(boil.GetDB(), cityID)
}

// CityExistsGP checks if the City row exists. Panics on error.
func CityExistsGP(cityID uint16) bool {
	e, err := CityExists(boil.GetDB(), cityID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CityExistsP checks if the City row exists. Panics on error.
func CityExistsP(exec boil.Executor, cityID uint16) bool {
	e, err := CityExists(exec, cityID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

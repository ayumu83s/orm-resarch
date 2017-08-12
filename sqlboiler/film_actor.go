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

// FilmActor is an object representing the database table.
type FilmActor struct {
	ActorID    uint16    `boil:"actor_id" json:"actor_id" toml:"actor_id" yaml:"actor_id"`
	FilmID     uint16    `boil:"film_id" json:"film_id" toml:"film_id" yaml:"film_id"`
	LastUpdate time.Time `boil:"last_update" json:"last_update" toml:"last_update" yaml:"last_update"`

	R *filmActorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L filmActorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FilmActorColumns = struct {
	ActorID    string
	FilmID     string
	LastUpdate string
}{
	ActorID:    "actor_id",
	FilmID:     "film_id",
	LastUpdate: "last_update",
}

// filmActorR is where relationships are stored.
type filmActorR struct {
	Actor *Actor
	Film  *Film
}

// filmActorL is where Load methods for each relationship are stored.
type filmActorL struct{}

var (
	filmActorColumns               = []string{"actor_id", "film_id", "last_update"}
	filmActorColumnsWithoutDefault = []string{"actor_id", "film_id"}
	filmActorColumnsWithDefault    = []string{"last_update"}
	filmActorPrimaryKeyColumns     = []string{"actor_id", "film_id"}
)

type (
	// FilmActorSlice is an alias for a slice of pointers to FilmActor.
	// This should generally be used opposed to []FilmActor.
	FilmActorSlice []*FilmActor
	// FilmActorHook is the signature for custom FilmActor hook methods
	FilmActorHook func(boil.Executor, *FilmActor) error

	filmActorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	filmActorType                 = reflect.TypeOf(&FilmActor{})
	filmActorMapping              = queries.MakeStructMapping(filmActorType)
	filmActorPrimaryKeyMapping, _ = queries.BindMapping(filmActorType, filmActorMapping, filmActorPrimaryKeyColumns)
	filmActorInsertCacheMut       sync.RWMutex
	filmActorInsertCache          = make(map[string]insertCache)
	filmActorUpdateCacheMut       sync.RWMutex
	filmActorUpdateCache          = make(map[string]updateCache)
	filmActorUpsertCacheMut       sync.RWMutex
	filmActorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var filmActorBeforeInsertHooks []FilmActorHook
var filmActorBeforeUpdateHooks []FilmActorHook
var filmActorBeforeDeleteHooks []FilmActorHook
var filmActorBeforeUpsertHooks []FilmActorHook

var filmActorAfterInsertHooks []FilmActorHook
var filmActorAfterSelectHooks []FilmActorHook
var filmActorAfterUpdateHooks []FilmActorHook
var filmActorAfterDeleteHooks []FilmActorHook
var filmActorAfterUpsertHooks []FilmActorHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FilmActor) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FilmActor) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FilmActor) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FilmActor) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FilmActor) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FilmActor) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FilmActor) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FilmActor) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FilmActor) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range filmActorAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFilmActorHook registers your hook function for all future operations.
func AddFilmActorHook(hookPoint boil.HookPoint, filmActorHook FilmActorHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		filmActorBeforeInsertHooks = append(filmActorBeforeInsertHooks, filmActorHook)
	case boil.BeforeUpdateHook:
		filmActorBeforeUpdateHooks = append(filmActorBeforeUpdateHooks, filmActorHook)
	case boil.BeforeDeleteHook:
		filmActorBeforeDeleteHooks = append(filmActorBeforeDeleteHooks, filmActorHook)
	case boil.BeforeUpsertHook:
		filmActorBeforeUpsertHooks = append(filmActorBeforeUpsertHooks, filmActorHook)
	case boil.AfterInsertHook:
		filmActorAfterInsertHooks = append(filmActorAfterInsertHooks, filmActorHook)
	case boil.AfterSelectHook:
		filmActorAfterSelectHooks = append(filmActorAfterSelectHooks, filmActorHook)
	case boil.AfterUpdateHook:
		filmActorAfterUpdateHooks = append(filmActorAfterUpdateHooks, filmActorHook)
	case boil.AfterDeleteHook:
		filmActorAfterDeleteHooks = append(filmActorAfterDeleteHooks, filmActorHook)
	case boil.AfterUpsertHook:
		filmActorAfterUpsertHooks = append(filmActorAfterUpsertHooks, filmActorHook)
	}
}

// OneP returns a single filmActor record from the query, and panics on error.
func (q filmActorQuery) OneP() *FilmActor {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single filmActor record from the query.
func (q filmActorQuery) One() (*FilmActor, error) {
	o := &FilmActor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: failed to execute a one query for film_actor")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FilmActor records from the query, and panics on error.
func (q filmActorQuery) AllP() FilmActorSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FilmActor records from the query.
func (q filmActorQuery) All() (FilmActorSlice, error) {
	var o []*FilmActor

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler: failed to assign all query results to FilmActor slice")
	}

	if len(filmActorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FilmActor records in the query, and panics on error.
func (q filmActorQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FilmActor records in the query.
func (q filmActorQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to count film_actor rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q filmActorQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q filmActorQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: failed to check if film_actor exists")
	}

	return count > 0, nil
}

// ActorG pointed to by the foreign key.
func (o *FilmActor) ActorG(mods ...qm.QueryMod) actorQuery {
	return o.Actor(boil.GetDB(), mods...)
}

// Actor pointed to by the foreign key.
func (o *FilmActor) Actor(exec boil.Executor, mods ...qm.QueryMod) actorQuery {
	queryMods := []qm.QueryMod{
		qm.Where("actor_id=?", o.ActorID),
	}

	queryMods = append(queryMods, mods...)

	query := Actors(exec, queryMods...)
	queries.SetFrom(query.Query, "`actor`")

	return query
}

// FilmG pointed to by the foreign key.
func (o *FilmActor) FilmG(mods ...qm.QueryMod) filmQuery {
	return o.Film(boil.GetDB(), mods...)
}

// Film pointed to by the foreign key.
func (o *FilmActor) Film(exec boil.Executor, mods ...qm.QueryMod) filmQuery {
	queryMods := []qm.QueryMod{
		qm.Where("film_id=?", o.FilmID),
	}

	queryMods = append(queryMods, mods...)

	query := Films(exec, queryMods...)
	queries.SetFrom(query.Query, "`film`")

	return query
} // LoadActor allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (filmActorL) LoadActor(e boil.Executor, singular bool, maybeFilmActor interface{}) error {
	var slice []*FilmActor
	var object *FilmActor

	count := 1
	if singular {
		object = maybeFilmActor.(*FilmActor)
	} else {
		slice = *maybeFilmActor.(*[]*FilmActor)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &filmActorR{}
		}
		args[0] = object.ActorID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &filmActorR{}
			}
			args[i] = obj.ActorID
		}
	}

	query := fmt.Sprintf(
		"select * from `actor` where `actor_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Actor")
	}
	defer results.Close()

	var resultSlice []*Actor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Actor")
	}

	if len(filmActorAfterSelectHooks) != 0 {
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
		object.R.Actor = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ActorID == foreign.ActorID {
				local.R.Actor = foreign
				break
			}
		}
	}

	return nil
}

// LoadFilm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (filmActorL) LoadFilm(e boil.Executor, singular bool, maybeFilmActor interface{}) error {
	var slice []*FilmActor
	var object *FilmActor

	count := 1
	if singular {
		object = maybeFilmActor.(*FilmActor)
	} else {
		slice = *maybeFilmActor.(*[]*FilmActor)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &filmActorR{}
		}
		args[0] = object.FilmID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &filmActorR{}
			}
			args[i] = obj.FilmID
		}
	}

	query := fmt.Sprintf(
		"select * from `film` where `film_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Film")
	}
	defer results.Close()

	var resultSlice []*Film
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Film")
	}

	if len(filmActorAfterSelectHooks) != 0 {
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
		object.R.Film = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FilmID == foreign.FilmID {
				local.R.Film = foreign
				break
			}
		}
	}

	return nil
}

// SetActorG of the film_actor to the related item.
// Sets o.R.Actor to related.
// Adds o to related.R.FilmActors.
// Uses the global database handle.
func (o *FilmActor) SetActorG(insert bool, related *Actor) error {
	return o.SetActor(boil.GetDB(), insert, related)
}

// SetActorP of the film_actor to the related item.
// Sets o.R.Actor to related.
// Adds o to related.R.FilmActors.
// Panics on error.
func (o *FilmActor) SetActorP(exec boil.Executor, insert bool, related *Actor) {
	if err := o.SetActor(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetActorGP of the film_actor to the related item.
// Sets o.R.Actor to related.
// Adds o to related.R.FilmActors.
// Uses the global database handle and panics on error.
func (o *FilmActor) SetActorGP(insert bool, related *Actor) {
	if err := o.SetActor(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetActor of the film_actor to the related item.
// Sets o.R.Actor to related.
// Adds o to related.R.FilmActors.
func (o *FilmActor) SetActor(exec boil.Executor, insert bool, related *Actor) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `film_actor` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"actor_id"}),
		strmangle.WhereClause("`", "`", 0, filmActorPrimaryKeyColumns),
	)
	values := []interface{}{related.ActorID, o.ActorID, o.FilmID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ActorID = related.ActorID

	if o.R == nil {
		o.R = &filmActorR{
			Actor: related,
		}
	} else {
		o.R.Actor = related
	}

	if related.R == nil {
		related.R = &actorR{
			FilmActors: FilmActorSlice{o},
		}
	} else {
		related.R.FilmActors = append(related.R.FilmActors, o)
	}

	return nil
}

// SetFilmG of the film_actor to the related item.
// Sets o.R.Film to related.
// Adds o to related.R.FilmActors.
// Uses the global database handle.
func (o *FilmActor) SetFilmG(insert bool, related *Film) error {
	return o.SetFilm(boil.GetDB(), insert, related)
}

// SetFilmP of the film_actor to the related item.
// Sets o.R.Film to related.
// Adds o to related.R.FilmActors.
// Panics on error.
func (o *FilmActor) SetFilmP(exec boil.Executor, insert bool, related *Film) {
	if err := o.SetFilm(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetFilmGP of the film_actor to the related item.
// Sets o.R.Film to related.
// Adds o to related.R.FilmActors.
// Uses the global database handle and panics on error.
func (o *FilmActor) SetFilmGP(insert bool, related *Film) {
	if err := o.SetFilm(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetFilm of the film_actor to the related item.
// Sets o.R.Film to related.
// Adds o to related.R.FilmActors.
func (o *FilmActor) SetFilm(exec boil.Executor, insert bool, related *Film) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `film_actor` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"film_id"}),
		strmangle.WhereClause("`", "`", 0, filmActorPrimaryKeyColumns),
	)
	values := []interface{}{related.FilmID, o.ActorID, o.FilmID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FilmID = related.FilmID

	if o.R == nil {
		o.R = &filmActorR{
			Film: related,
		}
	} else {
		o.R.Film = related
	}

	if related.R == nil {
		related.R = &filmR{
			FilmActors: FilmActorSlice{o},
		}
	} else {
		related.R.FilmActors = append(related.R.FilmActors, o)
	}

	return nil
}

// FilmActorsG retrieves all records.
func FilmActorsG(mods ...qm.QueryMod) filmActorQuery {
	return FilmActors(boil.GetDB(), mods...)
}

// FilmActors retrieves all the records using an executor.
func FilmActors(exec boil.Executor, mods ...qm.QueryMod) filmActorQuery {
	mods = append(mods, qm.From("`film_actor`"))
	return filmActorQuery{NewQuery(exec, mods...)}
}

// FindFilmActorG retrieves a single record by ID.
func FindFilmActorG(actorID uint16, filmID uint16, selectCols ...string) (*FilmActor, error) {
	return FindFilmActor(boil.GetDB(), actorID, filmID, selectCols...)
}

// FindFilmActorGP retrieves a single record by ID, and panics on error.
func FindFilmActorGP(actorID uint16, filmID uint16, selectCols ...string) *FilmActor {
	retobj, err := FindFilmActor(boil.GetDB(), actorID, filmID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFilmActor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFilmActor(exec boil.Executor, actorID uint16, filmID uint16, selectCols ...string) (*FilmActor, error) {
	filmActorObj := &FilmActor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `film_actor` where `actor_id`=? AND `film_id`=?", sel,
	)

	q := queries.Raw(exec, query, actorID, filmID)

	err := q.Bind(filmActorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: unable to select from film_actor")
	}

	return filmActorObj, nil
}

// FindFilmActorP retrieves a single record by ID with an executor, and panics on error.
func FindFilmActorP(exec boil.Executor, actorID uint16, filmID uint16, selectCols ...string) *FilmActor {
	retobj, err := FindFilmActor(exec, actorID, filmID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FilmActor) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FilmActor) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FilmActor) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FilmActor) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("sqlboiler: no film_actor provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(filmActorColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	filmActorInsertCacheMut.RLock()
	cache, cached := filmActorInsertCache[key]
	filmActorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			filmActorColumns,
			filmActorColumnsWithDefault,
			filmActorColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(filmActorType, filmActorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(filmActorType, filmActorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `film_actor` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `film_actor` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `film_actor` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, filmActorPrimaryKeyColumns))
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

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to insert into film_actor")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ActorID,
		o.FilmID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to populate default values for film_actor")
	}

CacheNoHooks:
	if !cached {
		filmActorInsertCacheMut.Lock()
		filmActorInsertCache[key] = cache
		filmActorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FilmActor record. See Update for
// whitelist behavior description.
func (o *FilmActor) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FilmActor record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FilmActor) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FilmActor, and panics on error.
// See Update for whitelist behavior description.
func (o *FilmActor) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FilmActor.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FilmActor) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	filmActorUpdateCacheMut.RLock()
	cache, cached := filmActorUpdateCache[key]
	filmActorUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			filmActorColumns,
			filmActorPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("sqlboiler: unable to update film_actor, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `film_actor` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, filmActorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(filmActorType, filmActorMapping, append(wl, filmActorPrimaryKeyColumns...))
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
		return errors.Wrap(err, "sqlboiler: unable to update film_actor row")
	}

	if !cached {
		filmActorUpdateCacheMut.Lock()
		filmActorUpdateCache[key] = cache
		filmActorUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q filmActorQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q filmActorQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to update all for film_actor")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FilmActorSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FilmActorSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FilmActorSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FilmActorSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filmActorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `film_actor` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filmActorPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to update all in filmActor slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FilmActor) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FilmActor) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FilmActor) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FilmActor) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("sqlboiler: no film_actor provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(filmActorColumnsWithDefault, o)

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

	filmActorUpsertCacheMut.RLock()
	cache, cached := filmActorUpsertCache[key]
	filmActorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			filmActorColumns,
			filmActorColumnsWithDefault,
			filmActorColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			filmActorColumns,
			filmActorPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("sqlboiler: unable to upsert film_actor, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "film_actor", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `film_actor` WHERE `actor_id`=? AND `film_id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(filmActorType, filmActorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(filmActorType, filmActorMapping, ret)
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

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to upsert for film_actor")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ActorID,
		o.FilmID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to populate default values for film_actor")
	}

CacheNoHooks:
	if !cached {
		filmActorUpsertCacheMut.Lock()
		filmActorUpsertCache[key] = cache
		filmActorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FilmActor record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FilmActor) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FilmActor record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FilmActor) DeleteG() error {
	if o == nil {
		return errors.New("sqlboiler: no FilmActor provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FilmActor record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FilmActor) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FilmActor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FilmActor) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("sqlboiler: no FilmActor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), filmActorPrimaryKeyMapping)
	sql := "DELETE FROM `film_actor` WHERE `actor_id`=? AND `film_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to delete from film_actor")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q filmActorQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q filmActorQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("sqlboiler: no filmActorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to delete all from film_actor")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FilmActorSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FilmActorSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("sqlboiler: no FilmActor slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FilmActorSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FilmActorSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("sqlboiler: no FilmActor slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(filmActorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filmActorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `film_actor` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filmActorPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to delete all from filmActor slice")
	}

	if len(filmActorAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FilmActor) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FilmActor) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FilmActor) ReloadG() error {
	if o == nil {
		return errors.New("sqlboiler: no FilmActor provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FilmActor) Reload(exec boil.Executor) error {
	ret, err := FindFilmActor(exec, o.ActorID, o.FilmID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FilmActorSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FilmActorSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FilmActorSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("sqlboiler: empty FilmActorSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FilmActorSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	filmActors := FilmActorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filmActorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `film_actor`.* FROM `film_actor` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filmActorPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&filmActors)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to reload all in FilmActorSlice")
	}

	*o = filmActors

	return nil
}

// FilmActorExists checks if the FilmActor row exists.
func FilmActorExists(exec boil.Executor, actorID uint16, filmID uint16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `film_actor` where `actor_id`=? AND `film_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, actorID, filmID)
	}

	row := exec.QueryRow(sql, actorID, filmID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: unable to check if film_actor exists")
	}

	return exists, nil
}

// FilmActorExistsG checks if the FilmActor row exists.
func FilmActorExistsG(actorID uint16, filmID uint16) (bool, error) {
	return FilmActorExists(boil.GetDB(), actorID, filmID)
}

// FilmActorExistsGP checks if the FilmActor row exists. Panics on error.
func FilmActorExistsGP(actorID uint16, filmID uint16) bool {
	e, err := FilmActorExists(boil.GetDB(), actorID, filmID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FilmActorExistsP checks if the FilmActor row exists. Panics on error.
func FilmActorExistsP(exec boil.Executor, actorID uint16, filmID uint16) bool {
	e, err := FilmActorExists(exec, actorID, filmID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
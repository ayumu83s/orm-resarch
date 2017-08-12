package sqlboiler

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ayumu83s/orm-resarch/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vattle/sqlboiler/boil"
	. "github.com/vattle/sqlboiler/queries/qm"
)

func New(ds *structs.DataSource) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		ds.User, ds.Password, ds.Host, ds.Port, ds.DbName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	boil.SetDB(db)
	boil.DebugMode = true
	return db, err
}

func Sample(db *sql.DB) {
	selectOne(db, 1)
	selectOne2(db, 2)
	selectOne3(db, 3)

	count()
	countByFirstName()
	searchFilmActorByActorID(1)
	searchFilmActorByActorID2(2)

	//searchFilmActorByActorIDWithActor(1)
	//searchFilmActorByActorIDWithActor2(1)
	searchFilmActorByActorIDWithActor3(2)
}

// ID指定で1件引くヤツ
func selectOne(db *sql.DB, id int) {
	// select * from `actor` where `actor_id`=?
	actor, err := FindActorG(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("selectOne: %+v\n", actor)
}

// ID指定で1件引くヤツ
func selectOne2(db *sql.DB, id int) {
	// SELECT * FROM `actor` WHERE (actor_id=?) LIMIT 1;
	// OneをAllにするとLIMIT 1の指定がない
	actor, err := Actors(db, Where("actor_id=?", id)).One()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("selectOne2: %+v\n", actor)
}

// ID指定で1件引くヤツ
func selectOne3(db *sql.DB, id int) {
	// SELECT * FROM `actor` WHERE actor_id=?
	// boil.SetDB(db)のDBを利用する
	actor, err := ActorsG(SQL("SELECT * FROM `actor` WHERE actor_id=?", id)).One()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("selectOne3: %+v\n", actor)
}

// 条件指定しないcount
func count() {
	// SELECT COUNT(*) FROM `actor`;
	count, err := ActorsG().Count()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count: %d\n", count)
}

// 条件指定のcount
func countByFirstName() {
	// SELECT COUNT(*) FROM `actor` WHERE (first_name LIKE ?);
	count, err := ActorsG(Where("first_name LIKE ?", "JO%")).Count()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count: %d\n", count)
}

func searchFilmActorByActorID(id int) {
	// SELECT * FROM `film_actor` WHERE (actor_id=?);
	filmActors, err := FilmActorsG(Where("actor_id=?", id)).All()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range filmActors {
		fmt.Printf("searchFilmActorByActorID: %+v\n", v)
	}
}

func searchFilmActorByActorID2(id int) {
	// SELECT film_id, last_update FROM `film_actor` WHERE actor_id=?
	filmActors, err := FilmActorsG(SQL("SELECT film_id, last_update FROM `film_actor` WHERE actor_id=?", id)).All()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range filmActors {
		fmt.Printf("searchFilmActorByActorID: %+v\n", v)
	}
}

// joinするやつ
func searchFilmActorByActorIDWithActor(id int) {
	type FilmActorsWithActor struct {
		FilmActor `boil:",bind"`
		Actor     `boil:",bind"`
	}

	var filmActorsWithActors []FilmActorsWithActor
	// SELECT film_actor.*, actor.* FROM `film_actor` INNER JOIN actor ON (film_actor.actor_id = actor.actor_id) WHERE (film_actor.actor_id = ?);
	err := NewQueryG(
		Select("film_actor.*, actor.*"),
		From("film_actor"),
		InnerJoin("actor ON (film_actor.actor_id = actor.actor_id)"),
		Where("film_actor.actor_id = ?", id),
	).Bind(&filmActorsWithActors)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range filmActorsWithActors {
		fmt.Printf("searchFilmActorByActorIDWithActor: %+v\n", v)
	}
}

// joinするやつ
// うごくけど、構造体の再定義がひたすらダルい
func searchFilmActorByActorIDWithActor2(id int) {
	type FilmActorsWithActor2 struct {
		ActorID    uint16
		FilmID     uint16
		LastUpdate time.Time
		FirstName  string
		LastName   string
		FilmTitle  string
	}

	var filmActorsWithActors []FilmActorsWithActor2
	err := NewQueryG(
		SQL(`
			SELECT
				film_actor.actor_id AS actor_id,
				film_actor.film_id AS film_id,
				film_actor.last_update AS last_update,
				actor.first_name AS first_name,
				actor.last_name AS last_name,
				film.title AS film_title
			FROM film_actor 
			INNER JOIN actor ON (
				film_actor.actor_id = actor.actor_id
			)
			INNER JOIN film ON (
				film_actor.film_id = film.film_id
			)
			WHERE film_actor.actor_id=?
		`, id),
	).Bind(&filmActorsWithActors)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range filmActorsWithActors {
		fmt.Printf("searchFilmActorByActorIDWithActor: %+v\n", v)
	}
}

// joinするやつ
// このあたりが落とし所な感じ
func searchFilmActorByActorIDWithActor3(id int) {
	type FilmActorsWithActor3 struct {
		FilmActor `boil:",bind"`
		Actor     `boil:",bind"`
		Film      `boil:",bind"`
	}

	var filmActorsWithActors []FilmActorsWithActor3
	err := NewQueryG(
		SQL(`
			SELECT
				film_actor.*,
				actor.*,
				film.*
			FROM film_actor 
			INNER JOIN actor ON (
				film_actor.actor_id = actor.actor_id
			)
			INNER JOIN film ON (
				film_actor.film_id = film.film_id
			)
			WHERE film_actor.actor_id=?
		`, id),
	).Bind(&filmActorsWithActors)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range filmActorsWithActors {
		fmt.Printf(
			"actor_id:%d, actor_name:%s, film_id:%d, film_title:%s\n",
			v.FilmActor.ActorID,
			v.Actor.FirstName,
			v.FilmActor.FilmID,
			v.Film.Title,
		)
	}
}

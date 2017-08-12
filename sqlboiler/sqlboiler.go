package sqlboiler

import (
	"database/sql"
	"fmt"

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

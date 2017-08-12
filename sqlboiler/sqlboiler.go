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

package xorm

import (
	"fmt"

	"github.com/ayumu83s/orm-resarch/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func New(ds *structs.DataSource) (*xorm.Engine, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		ds.User, ds.Password, ds.Host, ds.Port, ds.DbName,
	)

	engine, err := xorm.NewEngine("mysql", dataSourceName)
	engine.ShowSQL(true)
	//err = engine.Sync2(new(structs.Actor))
	return engine, err
}

func Sample(db *xorm.Engine) {
	selectOne(db, 1)
	selectOne2(db, 0)
	selectOne3(db, 3)

	count(db, 1)
	count(db, 0)
	count2(db, 1)
	//count3(db, 2)
}

// ID指定で1件引くヤツ
func selectOne(db *xorm.Engine, id int) {
	var actor = structs.Actor{
		ActorId: id,
	}
	// SELECT * FROM actor WHERE actor_id = ? LIMIT 1
	// ただし、idが0だと構造体の初期値と一緒だからか
	// SELECT * FROM actor LIMIT 1 とかいう地獄のSQLを発行するのでマジ危険
	has, err := db.Get(&actor)
	if err != nil {
		fmt.Println(err)
	}
	if has {
		fmt.Printf("selectOne: %+v\n", actor)
	}
}

// ID指定で1件引くヤツ2
func selectOne2(db *xorm.Engine, id int) {
	// SELECT * FROM actor WHERE actor_id = ? LIMIT 1
	// これなら0が来てもちゃんとwhere句つかってくれる
	var actor = structs.Actor{}
	has, err := db.Table("actor").Where("actor_id = ?", id).Get(&actor)
	if err != nil {
		fmt.Println(err)
	}
	if has {
		fmt.Printf("selectOne2: %+v\n", actor)
	}
}

// ID指定で1件引くヤツ3
func selectOne3(db *xorm.Engine, id int) {
	// SQL指定で結果を構造体にマッピングする
	// これこれ。こういうのでいいんだよ。
	var actor = structs.Actor{}
	sql := fmt.Sprintf("SELECT * FROM actor WHERE actor_id = %d", id)
	has, err := db.Sql(sql).Get(&actor)
	if err != nil {
		fmt.Println(err)
	}
	if has {
		fmt.Printf("selectOne3: %+v\n", actor)
	}
}

// 件数をカウントする
func count(db *xorm.Engine, actorID int) {
	// SELECT count(*) FROM film_actor WHERE actor_id = ?
	// ただし、actorIDが0だと構造体の初期値と一緒だからか
	// SELECT count(*) FROM film_actor とかいう地獄のSQLを発行するのでマジ危険
	var filmActor = structs.FilmActor{
		ActorId: actorID,
	}
	counts, err := db.Count(&filmActor)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count: %d\n", counts)
}

// 件数をカウントする2
func count2(db *xorm.Engine, actorID int) {
	// SQL指定で結果をマッピングする
	sql := fmt.Sprintf("SELECT count(*) FROM film_actor WHERE actor_id = %d", actorID)
	var counts int64
	_, err := db.Sql(sql).Get(&counts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count2: %d\n", counts)
}

// 件数をカウントする3
func count3(db *xorm.Engine, actorID int) {
	// これはエラーになる
	// Countには対応する構造体にしか渡せない
	var counts int64
	_, err := db.Table("film_actor").Where("actor_id = ?", actorID).Count(&counts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count3: %d\n", counts)
}

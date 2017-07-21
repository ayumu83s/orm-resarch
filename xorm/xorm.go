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

	search(db, 1)
	search(db, 0)
	search2(db, 1)
	search2(db, 0)
	search3(db, 1)
	search3(db, 0)
	search4(db, 1)
	search4(db, 0)
	search5(db, 1)
	search5(db, 0)
	search6(db, 1)
	search6(db, 0)

	//joinSearch(db, 1)
	joinSearch2(db, 1)

	//update(db, 1, "ACADEMY DINOSAUR2")
	//update2(db, 1, "ACADEMY DINOSAUR2")
	//update3(db, 1, "ACADEMY DINOSAUR2")
	update4(db, 1, "A Corua (La Corua)2")
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

// 複数のレコード取得
func search(db *xorm.Engine, actorID int) {
	// SELECT * FROM film_actor WHERE actor_id = ?
	// Where指定するとactorsが初期値でもwhere句が付く
	var actors = []structs.FilmActor{}
	err := db.Table("film_actor").Where("actor_id = ?", actorID).Find(&actors)
	if err != nil {
		fmt.Println(err)
	}
	var count int
	for _, v := range actors {
		count++
		fmt.Printf("search: %+v\n", v)
	}
	fmt.Printf("search: %d\n", count)
}

// 複数のレコード取得2
func search2(db *xorm.Engine, actorID int) {
	// これも条件指定が漏れると全件検索するので危険
	var actor = structs.FilmActor{
		ActorId: actorID,
	}
	var count int
	// SELECT * FROM film_actor WHERE actor_id = ?
	err := db.Iterate(&actor, func(idx int, bean interface{}) error {
		//actor := bean.(*structs.FilmActor)
		count++
		//fmt.Printf("search2: %+v\n", v)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("search2: %d\n", count)
}

// 複数のレコード取得3
func search3(db *xorm.Engine, actorID int) {
	// Where設定すれば漏れはない
	// actorで条件指定するとand条件に変わる。whereとマッピング用の構造体はどっちかにした方が良さそう
	//SELECT * FROM film_actor WHERE actor_id = ? AND actor_id = ?
	var actor = structs.FilmActor{}
	var count int
	err := db.Where("actor_id = ?", actorID).Iterate(&actor, func(idx int, bean interface{}) error {
		//actor := bean.(*structs.FilmActor)
		count++
		//fmt.Printf("search2: %+v\n", v)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("search3: %d\n", count)
}

// 複数のレコード取得4
func search4(db *xorm.Engine, actorID int) {
	var filmActors = []structs.FilmActor{}
	sql := fmt.Sprintf("SELECT * FROM film_actor WHERE actor_id = %d", actorID)

	// SELECT * FROM film_actor WHERE actor_id = ?
	// Getはhas, errを返すけど、Findはerrだけ
	err := db.Sql(sql).Find(&filmActors)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range filmActors {
		fmt.Printf("search4: %+v\n", v)
	}
}

// rowsを使ってみる版
func search5(db *xorm.Engine, actorID int) {
	var actor = structs.FilmActor{}
	var count int
	rows, err := db.Where("actor_id = ?", actorID).Rows(&actor)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		bean := new(structs.FilmActor)
		err = rows.Scan(bean)
		fmt.Printf("search5: %+v\n", bean)
		count++
	}
	fmt.Printf("search5: %d\n", count)
}

// SQL指定でRowsを使う版
// 一度に大量のレコードを引くときは、メモリ消費をおさえられるのかも。
func search6(db *xorm.Engine, actorID int) {
	var filmActor = structs.FilmActor{}
	sql := fmt.Sprintf("SELECT * FROM film_actor WHERE actor_id = %d", actorID)
	var count int

	rows, err := db.Sql(sql).Rows(&filmActor)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		bean := new(structs.FilmActor)
		err = rows.Scan(bean)
		fmt.Printf("search6: %+v\n", bean)
		count++
	}
	fmt.Printf("search6: %d\n", count)
}

// 他のテーブルjoinして検索
func joinSearch(db *xorm.Engine, actorID int) {
	type FilmActorDetail struct {
		structs.FilmActor `xorm:"extends"` // tagがないとマッピングできない
		//structs.FilmActor
		FirstName string
		LastName  string
	}
	var filmActorDetail = []FilmActorDetail{}
	err := db.Table("film_actor").
		Select("film_actor.*, actor.first_name, actor.last_name").
		Join("INNER", "actor", "film_actor.actor_id = actor.actor_id").
		Where("film_actor.actor_id = ?", actorID).
		Find(&filmActorDetail)

	if err != nil {
		fmt.Println(err)
	}
	var count int
	for _, v := range filmActorDetail {
		count++
		fmt.Printf("joinSearch: %+v\n", v)
		// fmt.Printf("FilmId: %d\n", v.FilmId)
		// fmt.Printf("ActorId: %d\n", v.ActorId)
		// fmt.Printf("FirstName: %s\n", v.FirstName)
		fmt.Println("----")
	}
	fmt.Printf("joinSearch: %d\n", count)
}

// 他のテーブルjoinして検索2
func joinSearch2(db *xorm.Engine, actorID int) {
	type FilmActorDetail struct {
		structs.FilmActor `xorm:"extends"` // tagがないとマッピングできない
		//structs.FilmActor
		ActorFirstName string // 別名でも規則さえあえばマッピングできる
		ActorLastName  string
	}
	var filmActorDetail = []FilmActorDetail{}
	sql := fmt.Sprintf(`
		SELECT
			film_actor.*, actor.first_name AS actor_first_name, actor.last_name AS actor_last_name
		FROM film_actor
		INNER JOIN actor ON (
			film_actor.actor_id = actor.actor_id
		)
		WHERE
			film_actor.actor_id = %d
	`, actorID)

	err := db.Sql(sql).Find(&filmActorDetail)

	if err != nil {
		fmt.Println(err)
	}
	var count int
	for _, v := range filmActorDetail {
		count++
		fmt.Printf("joinSearch: %+v\n", v)
		// fmt.Printf("FilmId: %d\n", v.FilmId)
		// fmt.Printf("ActorId: %d\n", v.ActorId)
		// fmt.Printf("FirstName: %s\n", v.FirstName)
		fmt.Println("----")
	}
	fmt.Printf("joinSearch: %d\n", count)
}

// timeoutする…
// Error 1205: Lock wait timeout exceeded; try restarting transaction
// film_listも更新しようとしている？？
func update(db *xorm.Engine, filmID int, title string) {
	var fileText = structs.FilmText{
		Title: title,
	}
	// UPDATE film_text SET title = "title" WHERE film_id = 1
	affected, err := db.Where("film_id = ?", filmID).Update(&fileText)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("update: %d\n", affected)
}

// timeoutする…
func update2(db *xorm.Engine, filmID int, title string) {
	// UPDATE film_text SET title = "title" WHERE film_id = 1
	affected, err := db.Exec(
		"UPDATE film_text SET title = ? WHERE film_id = ?",
		title, filmID,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("update: %d\n", affected)
}

// timeoutする…
func update3(db *xorm.Engine, filmID int, title string) {
	// UPDATE film_text SET title = "title" WHERE film_id = 1
	sql := fmt.Sprintf(`
		UPDATE film_text SET title = "%s" WHERE film_id = %d
		`, title, filmID,
	)
	results, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("update: %d\n", results)
}

// リレーションがなければサクッと完了
func update4(db *xorm.Engine, cityID int, cityName string) {
	var city = structs.City{
		City: cityName,
	}
	// UPDATE city SET city = "cityName" WHERE city_id = 1
	affected, err := db.Where("city_id = ?", cityID).Update(&city)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("update4: %d\n", affected)
}

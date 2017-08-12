package sqlboiler

import (
	"database/sql"
	"fmt"

	"github.com/ayumu83s/orm-resarch/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vattle/sqlboiler/boil"
)

func New(ds *structs.DataSource) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		ds.User, ds.Password, ds.Host, ds.Port, ds.DbName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	boil.SetDB(db)
	return db, err
}

func Sample(db *sql.DB) {
	selectOne(db, 1)
}

// ID指定で1件引くヤツ
func selectOne(db *sql.DB, id int) {
	actor, err := FindActorG(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("selectOne: %+v\n", actor)
}

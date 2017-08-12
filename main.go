package main

import (
	"fmt"

	"github.com/ayumu83s/orm-resarch/sqlboiler"
	"github.com/ayumu83s/orm-resarch/structs"
)

func main() {
	ds := &structs.DataSource{
		DbName:   "sakila",
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "",
	}

	// db, err := xorm.New(ds)
	// if err != nil {
	// 	fmt.Printf("err: %#v\n", err)
	// }
	// xorm.Sample(db)
	db, err := sqlboiler.New(ds)
	if err != nil {
		fmt.Printf("err: %#v\n", err)
	}
	sqlboiler.Sample(db)
	db.Close()
}

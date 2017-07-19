package main

import (
	"fmt"

	"github.com/ayumu83s/orm-resarch/structs"
	"github.com/ayumu83s/orm-resarch/xorm"
)

func main() {
	ds := &structs.DataSource{
		DbName:   "sakila",
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "",
	}

	db, err := xorm.New(ds)
	if err != nil {
		fmt.Printf("err: %#v\n", err)
	}
	xorm.Sample(db)
}

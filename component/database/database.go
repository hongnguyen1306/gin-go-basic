package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

func ConnectDatabase() *pg.DB {
	opt, err := pg.ParseURL("postgres://postgres:123@host.docker.internal:5435/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	var n int
	_, err2 := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(n)

	return db
}

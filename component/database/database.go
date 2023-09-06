package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

func ConnectDatabase() *pg.DB {
	opt, err := pg.ParseURL("postgres://postgres:123@postgres:5435/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	fmt.Println("pass", opt)
	db := pg.Connect(opt)
	return db
}

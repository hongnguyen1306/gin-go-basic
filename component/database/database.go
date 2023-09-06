package database

import (
	"github.com/go-pg/pg/v10"
)

func ConnectDatabase() *pg.DB {
	opt, err := pg.ParseURL("postgres://postgres:postgres@postgres:5432/postgres")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	return db
}

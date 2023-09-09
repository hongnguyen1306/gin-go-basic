package database

import (
	"app/common"

	"github.com/go-pg/pg/v10"
)

func New(config *common.Config) *pg.DB {
	opt, _ := pg.ParseURL(config.Dsn)
	db := pg.Connect(opt)
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	return db
}

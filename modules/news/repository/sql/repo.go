package sql

import "github.com/go-pg/pg/v10"

type sqlRepo struct {
	db *pg.DB
}

func NewSQLRepo(db *pg.DB) *sqlRepo {
	return &sqlRepo{db}
}

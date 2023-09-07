package sql

import (
	"app/modules/user/entity"
	"context"

	"github.com/go-pg/pg/v10"
)

type sqlRepo struct {
	db *pg.DB
}

// FindUser implements business.LoginStore.
func (*sqlRepo) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*entity.User, error) {
	panic("unimplemented")
}

func NewSQLRepo(db *pg.DB) *sqlRepo {
	return &sqlRepo{db: db}
}

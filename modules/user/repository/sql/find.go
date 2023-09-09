package sql

import (
	"app/modules/user/entity"
	"context"

	"github.com/go-pg/pg/v10"
	"golang.org/x/exp/maps"
)

func (sql *sqlRepo) FindData(context context.Context, condition map[string]interface{}) (*entity.User, error) {

	var data entity.User

	if err := sql.db.Model(&data).Where("? = ?", pg.Ident(maps.Keys(condition)[0]), maps.Values(condition)[0]).Select(); err != nil {
		return nil, err
	}

	return &data, nil
}

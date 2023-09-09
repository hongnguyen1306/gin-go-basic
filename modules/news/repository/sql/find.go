package sql

import (
	"app/modules/news/entity"
	"context"

	"github.com/go-pg/pg/v10"
	"golang.org/x/exp/maps"
)

func (sql *sqlRepo) Find(context context.Context, condition map[string]interface{}) (*entity.News, error) {
	var data entity.News

	if err := sql.db.Model(&data).Where("? = ?", pg.Ident(maps.Keys(condition)[0]), maps.Values(condition)[0]).Select(); err != nil {
		return nil, err
	}

	return &data, nil
}

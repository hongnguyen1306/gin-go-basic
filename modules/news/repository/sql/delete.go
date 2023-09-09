package sql

import (
	"app/modules/news/entity"
	"context"

	"github.com/go-pg/pg/v10/orm"
)

func (sql *sqlRepo) Delete(context context.Context, id string) (orm.Result, error) {
	data := &entity.News{Id: id}
	res, err := sql.db.Model(data).Where("id = ?", id).Delete()
	if err != nil {
		return nil, err
	}

	return res, nil
}

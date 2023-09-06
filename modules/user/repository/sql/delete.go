package sql

import (
	"app/modules/user/entity"
	"context"

	"github.com/go-pg/pg/v10/orm"
)

func (sql *sqlRepo) DeleteData(context context.Context, id string) (orm.Result, error) {

	data := &entity.User{Id: id}
	res, err := sql.db.Model(data).Where("id = ?", id).Delete()
	if err != nil {
		return nil, err
	}

	return res, nil
}

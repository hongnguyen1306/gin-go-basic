package sql

import (
	"app/modules/user/entity"
	"context"
)

func (sql *sqlRepo) FindData(context context.Context, id string) (*entity.User, error) {

	data := &entity.User{Id: id}
	if err := sql.db.Model(data).WherePK().Select(); err != nil {
		return nil, err
	}

	return data, nil
}

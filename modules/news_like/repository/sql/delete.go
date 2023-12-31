package sql

import (
	"app/modules/news_like/entity"
	"context"

	"github.com/go-pg/pg/v10/orm"
)

func (sql *sqlRepo) Delete(context context.Context, userId string, newsId string) (orm.Result, error) {
	data := &entity.Like{}
	res, err := sql.db.Model(data).Where("user_id = ? and news_id = ?", userId, newsId).Delete()

	if err != nil {
		return nil, err
	}

	return res, nil
}

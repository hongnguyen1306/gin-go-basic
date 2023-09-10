package sql

import (
	"app/modules/news_like/entity"
	"context"
)

func (s *sqlRepo) List(context context.Context) ([]entity.Like, error) {
	var result []entity.Like

	if err := s.db.Model(&result).Select(); err != nil {
		return nil, err
	}

	return result, nil

}

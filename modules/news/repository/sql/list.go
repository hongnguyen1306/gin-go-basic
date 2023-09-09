package sql

import (
	"app/modules/news/entity"
	"context"
)

func (s *sqlRepo) List(context context.Context) ([]entity.News, error) {
	var result []entity.News

	if err := s.db.Model(&result).Select(); err != nil {
		return nil, err
	}

	return result, nil

}

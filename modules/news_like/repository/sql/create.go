package sql

import (
	"app/modules/news_like/entity"
	"context"
)

func (s *sqlRepo) Create(c context.Context, data *entity.Like) error {
	_, err := s.db.WithContext(c).Model(data).Insert()
	if err != nil {
		return err
	}

	return nil
}

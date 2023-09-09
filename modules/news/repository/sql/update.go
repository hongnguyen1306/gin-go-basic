package sql

import (
	"app/modules/news/entity"
	"context"
	"time"
)

func (s *sqlRepo) Update(c context.Context, id string, data *entity.NewsUpdate) error {

	data.UpdatedAt = time.Now()

	_, err := s.db.Model(data).Where("id = ?", id).Update()
	if err != nil {
		return err
	}
	return nil
}

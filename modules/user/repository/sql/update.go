package sql

import (
	"app/modules/user/entity"
	"context"
	"time"
)

func (s *sqlRepo) UpdateData(context context.Context, id string, data *entity.UserUpdate) error {
	data.UpdatedAt = time.Now()

	_, err := s.db.Model(data).Where("id = ?", id).Update()
	if err != nil {
		return err
	}
	return nil
}

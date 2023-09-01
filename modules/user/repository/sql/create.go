package sql

import (
	"app/modules/user/entity"
	"context"
)

func (s *sqlRepo) Create(c context.Context, data *entity.User) error {
	_, err := s.db.Model(data).Insert()
	if err != nil {
		return err
	}

	return nil
}

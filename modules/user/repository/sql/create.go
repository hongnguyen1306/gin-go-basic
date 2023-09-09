package sql

import (
	"app/modules/user/entity"
	"context"
)

func (s *sqlRepo) CreateData(c context.Context, data *entity.User) error {
	_, err := s.db.WithContext(c).Model(data).Insert()
	if err != nil {
		return err
	}

	return nil
}

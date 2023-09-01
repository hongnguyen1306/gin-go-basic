package sql

import (
	"app/modules/user/entity"
	"context"
)

func (s *sqlRepo) ListData(context context.Context) ([]entity.User, error) {
	var result []entity.User

	if err := s.db.Model(&result).Select(); err != nil {
		return nil, err
	}

	return result, nil

}

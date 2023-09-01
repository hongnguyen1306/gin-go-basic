package business

import (
	"app/modules/user/entity"
	"context"
)

func (biz *business) ListUser(ctx context.Context) ([]entity.User, error) {
	result, err := biz.userRepo.ListData(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}

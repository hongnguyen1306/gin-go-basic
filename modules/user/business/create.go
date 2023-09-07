package business

import (
	"app/modules/user/entity"
	"context"
)

func (biz *business) CreateUser(ctx context.Context, data *entity.User) error {
	if err := biz.userRepo.CreateData(ctx, data); err != nil {
		return err
	}
	return nil
}

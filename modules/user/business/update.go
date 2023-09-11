package business

import (
	"app/modules/user/entity"
	"context"
)

func (biz *business) UpdateUser(ctx context.Context, id string, data *entity.UserUpdate) (*entity.User, error) {
	err := biz.userRepo.UpdateData(ctx, id, data)
	if err != nil {
		return nil, err
	}

	updatedUser, err := biz.userRepo.FindData(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

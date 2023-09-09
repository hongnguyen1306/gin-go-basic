package business

import (
	"app/modules/user/entity"
	"context"
)

func (biz *business) FindUser(ctx context.Context, id string) (*entity.User, error) {
	result, err := biz.userRepo.FindData(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}

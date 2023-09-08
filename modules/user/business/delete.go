package business

import (
	"context"

	"github.com/go-pg/pg/v10/orm"
)

func (biz *business) DeleteUser(ctx context.Context, id string) (orm.Result, error) {
	_, err := biz.userRepo.FindData(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	result, err := biz.userRepo.DeleteData(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

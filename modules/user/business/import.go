package business

import (
	"app/modules/user/entity"
	"context"
)

func (biz *business) ImportUserCSV(ctx context.Context, data []*entity.User) error {
	if err := biz.userRepo.ImportDataCSV(ctx, data); err != nil {
		return err
	}
	return nil
}

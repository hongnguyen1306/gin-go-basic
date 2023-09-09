package business

import (
	"context"

	"github.com/go-pg/pg/v10/orm"
)

func (biz *business) DeleteNews(ctx context.Context, id string) (orm.Result, error) {
	_, err := biz.newsRepo.Find(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	result, err := biz.newsRepo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

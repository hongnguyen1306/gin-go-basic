package business

import (
	"app/modules/news/entity"
	"context"
)

func (biz *business) CreateNews(ctx context.Context, data *entity.News) error {
	if err := biz.newsRepo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

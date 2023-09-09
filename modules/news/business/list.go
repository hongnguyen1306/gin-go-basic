package business

import (
	"app/modules/news/entity"
	"context"
)

func (biz *business) ListNews(ctx context.Context) ([]entity.News, error) {
	result, err := biz.newsRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

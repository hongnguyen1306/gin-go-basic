package business

import (
	"app/modules/news/entity"
	"context"
)

func (biz *business) FindNews(ctx context.Context, id string) (*entity.News, error) {
	result, err := biz.newsRepo.Find(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}

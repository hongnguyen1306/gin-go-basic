package business

import (
	"app/modules/news/entity"
	"context"
)

func (biz *business) UpdateNews(ctx context.Context, id string, data *entity.NewsUpdate) (*entity.News, error) {
	err := biz.newsRepo.Update(ctx, id, data)
	if err != nil {
		return nil, err
	}

	updatedNews, err := biz.newsRepo.Find(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return updatedNews, nil
}

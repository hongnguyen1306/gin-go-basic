package business

import (
	"app/modules/news_like/entity"
	"context"
)

type UserLikeNewsRepo interface {
	Create(ctx context.Context, data *entity.Like) error
}

type userLikeNewsBusiness struct {
	repo UserLikeNewsRepo
}

func NewUserLikeNewsBusiness(repo UserLikeNewsRepo) *userLikeNewsBusiness {
	return &userLikeNewsBusiness{repo}
}

func (biz *userLikeNewsBusiness) LikeNews(ctx context.Context, data *entity.Like) error {
	err := biz.repo.Create(ctx, data)

	if err != nil {
		return err
	}
	return nil
}

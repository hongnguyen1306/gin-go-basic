package business

import (
	"context"

	"github.com/go-pg/pg/v10/orm"
)

type UserUnlikeNewsRepo interface {
	Delete(ctx context.Context, userId, newsId string) (orm.Result, error)
}

type userUnlikeNewsBusiness struct {
	repo UserUnlikeNewsRepo
}

func NewUserUnlikeNewsBusiness(repo UserUnlikeNewsRepo) *userUnlikeNewsBusiness {
	return &userUnlikeNewsBusiness{repo}
}

func (biz *userUnlikeNewsBusiness) UnlikeNews(ctx context.Context, userId, newsId string) error {
	_, err := biz.repo.Delete(ctx, userId, newsId)

	if err != nil {
		return err
	}

	return nil
}

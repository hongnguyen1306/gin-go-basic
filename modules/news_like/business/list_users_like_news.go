package business

import (
	"app/common"
	"app/modules/news_like/entity"
	"context"
)

type ListUserLikeNewsRepo interface {
	GetUsersLikeNews(ctx context.Context, conditions map[string]interface{}, filter *entity.Filter, paging *common.Paging) ([]common.SimpleUser, error)
}

type listUserLikeNews struct {
	repo ListUserLikeNewsRepo
}

func NewListUserLikeNews(repo ListUserLikeNewsRepo) *listUserLikeNews {
	return &listUserLikeNews{repo}
}

func (biz *listUserLikeNews) ListUser(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]common.SimpleUser, error) {
	users, err := biz.repo.GetUsersLikeNews(ctx, nil, filter, paging)

	if err != nil {
		return nil, err
	}

	return users, nil
}

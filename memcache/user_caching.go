package memcache

import (
	"app/modules/user/entity"
	"context"
	"fmt"
)

type RealStore interface {
	FindData(ctx context.Context, condition map[string]interface{}) (*entity.User, error)
}

type userCaching struct {
	store     Caching
	realStore RealStore
}

func NewUserCaching(store Caching, realStore RealStore) *userCaching {
	return &userCaching{
		store:     store,
		realStore: realStore,
	}
}

func (uc *userCaching) FindData(ctx context.Context, condition map[string]interface{}) (*entity.User, error) {
	userId := condition["id"]
	key := fmt.Sprintf("user-%d", userId)
	userInCache := uc.store.Read(key)
	if userInCache != nil {
		return userInCache.(*entity.User), nil
	}

	user, err := uc.realStore.FindData(ctx, condition)
	if err != nil {
		return nil, err
	}

	uc.store.Write(key, user)
	return user, nil
}

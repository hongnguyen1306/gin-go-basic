package business

import (
	"app/modules/news/entity"
	"context"

	"github.com/go-pg/pg/v10/orm"
)

type NewsRepository interface {
	Create(context context.Context, data *entity.News) error
	Find(context context.Context, condition map[string]interface{}) (*entity.News, error)
	List(context context.Context) ([]entity.News, error)
	Update(context context.Context, id string, data *entity.NewsUpdate) error
	Delete(context context.Context, id string) (orm.Result, error)
}

type business struct {
	newsRepo NewsRepository
}

func NewBusiness(newsRepo NewsRepository) *business {
	return &business{newsRepo}
}

package business

import (
	"app/modules/user/entity"
	"context"

	"github.com/go-pg/pg/v10/orm"
)

type UserRepository interface {
	Create(context context.Context, data *entity.User) error
	ListData(context context.Context) ([]entity.User, error)
	FindData(context context.Context, id string) (*entity.User, error)
	DeleteData(context context.Context, id string) (orm.Result, error)
	ImportDataCSV(context context.Context, data []*entity.User) error
}

type business struct {
	userRepo UserRepository
}

func NewBusiness(userRepo UserRepository) *business {
	return &business{userRepo: userRepo}
}

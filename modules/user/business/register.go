package business

import (
	"app/common"
	"app/modules/user/entity"
	"context"
	"fmt"

	guuid "github.com/google/uuid"
)

type RegisterStorage interface {
	FindData(ctx context.Context, condition map[string]interface{}) (*entity.User, error)
	CreateData(ctx context.Context, data *entity.User) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterStorage(registerStorage RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *entity.User) error {
	user, err := biz.registerStorage.FindData(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		fmt.Println(err)
	}

	if user != nil {
		return entity.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	userCreate := &entity.User{
		Id:           guuid.New().String(),
		FullName:     data.FullName,
		EmployeeCode: data.EmployeeCode,
		Email:        data.Email,
		Role:         data.Role,
		Password:     data.Password,
	}

	userCreate.Password = biz.hasher.Hash(userCreate.Password + salt)
	userCreate.Salt = salt
	if err := biz.registerStorage.CreateData(ctx, userCreate); err != nil {
		panic(err)
	}
	return nil
}

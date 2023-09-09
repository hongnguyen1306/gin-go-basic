package business

import (
	"app/component/hasher"
	"app/component/tokenprovider"
	"app/modules/user/entity"
	"context"
)

type LoginStore interface {
	FindData(ctx context.Context, condition map[string]interface{}) (*entity.User, error)
}

type loginBusiness struct {
	storeUser     LoginStore
	tokenProvider tokenprovider.Provider
	hasher        hasher.Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStore, tokenProvider tokenprovider.Provider, hasher hasher.Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *entity.UserLogin) (*entity.Account, error) {
	user, err := biz.storeUser.FindData(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, err
	}
	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, err
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, err
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(accessToken, refreshToken)

	return account, nil
}

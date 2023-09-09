package middleware

import (
	"app/common"
	"app/component/app_context"
	"app/component/tokenprovider/jwt"
	"app/modules/user/entity"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenStore interface {
	FindData(ctx context.Context, condition map[string]interface{}) (*entity.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(err,
		"wrong authen header",
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	fmt.Println("header ", s)
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequireAuth(appCtx app_context.AppContext, authenStore AuthenStore) func(c *gin.Context) {
	tokenProvide := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvide.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authenStore.FindData(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}
		c.Set("user", user)
		c.Next()
	}
}

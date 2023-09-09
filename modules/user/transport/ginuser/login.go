package ginuser

import (
	"app/common"
	"app/component/app_context"
	"app/component/hasher/md5"
	"app/component/tokenprovider/jwt"
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var SecretKey = os.Getenv("SYSTEM_SECRET")

func HandleLogin(appCtx app_context.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("SecretKey ", SecretKey)
		tokenProvider := jwt.NewTokenJWTProvider(SecretKey)
		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		md5 := md5.NewMd5Hash()

		biz := business.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(account, nil, nil))
	}
}

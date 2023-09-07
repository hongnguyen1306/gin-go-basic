package ginuser

import (
	"app/component/hasher/md5"
	"app/component/tokenprovider/jwt"
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

var SecretKey = os.Getenv("SYSTEM_SECRET")

func HandleLogin(db *pg.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		tokenProvider := jwt.NewTokenJWTProvider(SecretKey)
		store := sql.NewSQLRepo(db)
		md5 := md5.NewMd5Hash()

		biz := business.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(200, account)
	}
}

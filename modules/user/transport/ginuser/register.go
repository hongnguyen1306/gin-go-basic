package ginuser

import (
	"app/component/hasher/md5"
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func HandleRegister(db *pg.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.User
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		store := sql.NewSQLRepo(db)
		md5 := md5.NewMd5Hash()
		biz := business.NewRegisterStorage(store, md5)

		if err := biz.Register(c.Request.Context(), data); err != nil {
			panic(err)
		}
		c.JSON(200, "OK!!!")
	}

}

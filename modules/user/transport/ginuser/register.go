package ginuser

import (
	"app/component/app_context"
	"app/component/hasher/md5"
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRegister(appCtx app_context.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.User
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		md5 := md5.NewMd5Hash()
		biz := business.NewRegisterStorage(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			log.Printf("Error while find a user, Reason: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "OK!!!",
		})
	}

}

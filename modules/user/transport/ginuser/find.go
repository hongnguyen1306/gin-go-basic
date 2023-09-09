package ginuser

import (
	"app/component/app_context"
	"app/modules/user/business"
	"app/modules/user/repository/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFindUser(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("userId")

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		result, err := biz.FindUser(c.Request.Context(), id)
		if err != nil {
			log.Printf("Error while find a user, Reason: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   result,
		})

	}
}

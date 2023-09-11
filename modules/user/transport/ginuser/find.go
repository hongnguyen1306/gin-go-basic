package ginuser

import (
	"app/common"
	"app/component/app_context"
	"app/modules/user/business"
	"app/modules/user/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFindUser(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		result, err := biz.FindUser(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))

	}
}

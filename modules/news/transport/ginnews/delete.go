package ginnews

import (
	"app/common"
	"app/component/app_context"
	"app/modules/news/business"
	"app/modules/news/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandelDeleteNews(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		_, err := biz.DeleteNews(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse("Deleted", nil, nil))
	}
}

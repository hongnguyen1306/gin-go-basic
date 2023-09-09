package ginnews

import (
	"app/common"
	"app/component/app_context"
	"app/modules/news/business"
	"app/modules/news/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HanldListNews(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		result, err := biz.ListNews(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}

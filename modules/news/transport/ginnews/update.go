package ginnews

import (
	"app/common"
	"app/component/app_context"
	"app/modules/news/business"
	"app/modules/news/entity"
	"app/modules/news/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUpdateNews(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("newsId")

		var data entity.NewsUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)
		result, err := biz.UpdateNews(c.Request.Context(), id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))

	}
}

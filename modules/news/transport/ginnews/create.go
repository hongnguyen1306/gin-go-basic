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

func HandelCreateNews(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entity.News

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.CreatorId = requester.GetUserId()

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		if err := biz.CreateNews(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))

	}
}

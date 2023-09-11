package ginnewslike

import (
	"app/common"
	"app/component/app_context"
	"app/modules/news_like/business"
	"app/modules/news_like/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserUnlikeNews(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsId := c.Param("newsId")

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewUserUnlikeNewsBusiness(store)

		err := biz.UnlikeNews(c.Request.Context(), requester.GetUserId(), newsId)
		if err != nil {
			panic(err)
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse("Unlike", nil, nil))
	}
}

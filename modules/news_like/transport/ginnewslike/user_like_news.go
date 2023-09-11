package ginnewslike

import (
	"app/common"
	"app/component/app_context"
	"app/modules/news_like/business"
	"app/modules/news_like/entity"
	"app/modules/news_like/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserLikeNews(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsId := c.Param("id")

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := entity.Like{
			NewsId: newsId,
			UserId: requester.GetUserId(),
		}

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewUserLikeNewsBusiness(store)

		err := biz.LikeNews(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
	}
}

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

func HandleListUserLikeNews(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsId := c.Param("id")

		filter := entity.Filter{
			NewsId: newsId,
		}

		var paging common.Paging

		if err := c.ShouldBindJSON(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		paging.Fulfill()
		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewListUserLikeNews(store)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

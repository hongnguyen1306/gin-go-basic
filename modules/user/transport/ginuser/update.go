package ginuser

import (
	"app/common"
	"app/component/app_context"
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUpdateaUser(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()

		var data entity.UserUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)
		result, err := biz.UpdateUser(c.Request.Context(), userId, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))

	}
}

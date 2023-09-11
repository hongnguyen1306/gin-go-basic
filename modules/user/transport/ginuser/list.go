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

func HandleListUser(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		if requester.GetRole() != entity.SuperAdmin {
			c.JSON(http.StatusForbidden, common.NewFailResponse("Unauthorized"))
			return
		}

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		result, err := biz.ListUser(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))

	}
}

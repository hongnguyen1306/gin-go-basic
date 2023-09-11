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

func HandleDeleteUser(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("userId")

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		if requester.GetRole() != entity.SuperAdmin {
			c.JSON(http.StatusForbidden, common.NewFailResponse("Bạn không có quyền thực hiện hành động này"))
			return
		}

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		_, err := biz.DeleteUser(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse("Deleted!!", nil, nil))
	}
}

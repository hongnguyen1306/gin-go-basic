package ginuser

import (
	"app/modules/user/business"
	"app/modules/user/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func HandleListUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		store := sql.NewSQLRepo(db)
		biz := business.NewBusiness(store)

		result, err := biz.ListUser(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})

	}
}

package ginuser

import (
	"app/modules/user/business"
	"app/modules/user/repository/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func HandleDeleteUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("userId")
		fmt.Println("id ", id)

		store := sql.NewSQLRepo(db)
		biz := business.NewBusiness(store)

		result, err := biz.DeleteUser(c.Request.Context(), id)
		if err != nil {
			log.Printf("Error while find a user, Reason: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}

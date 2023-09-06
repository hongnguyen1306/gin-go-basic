package ginuser

import (
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	guuid "github.com/google/uuid"
)

func HandleCreateUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("req", user)
		data := &entity.User{
			Id:           guuid.New().String(),
			FullName:     user.FullName,
			EmployeeCode: user.EmployeeCode,
			Email:        user.Email,
			Role:         user.Role,
		}

		store := sql.NewSQLRepo(db)
		biz := business.NewBusiness(store)

		insertError := biz.CreateUser(c.Request.Context(), data)
		if insertError != nil {
			log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User created Successfully",
		})
	}
}

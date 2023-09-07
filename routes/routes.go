package routes

import (
	"app/modules/user/transport/ginuser"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func SetupRouter(dbConnect *pg.DB) *gin.Engine {
	r := gin.Default()
	userRoutes := r.Group("/api")

	{
		userRoutes.GET("/", welcome)
		userRoutes.POST("/register", ginuser.HandleRegister(dbConnect))
		userRoutes.POST("/user", ginuser.HandleCreateUser(dbConnect))
		userRoutes.POST("/login", ginuser.HandleLogin(dbConnect))
		userRoutes.POST("/import", ginuser.HandleImportUserCsv(dbConnect))
		userRoutes.GET("/users", ginuser.HandleListUser(dbConnect))
		userRoutes.GET("/user/:userId", ginuser.HandleFindUser(dbConnect))
		userRoutes.DELETE("/user/:userId", ginuser.HandleDeleteUser(dbConnect))

	}
	return r
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

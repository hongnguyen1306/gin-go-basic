package routes

import (
	"app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userRoutes := r.Group("/api")

	{
		userRoutes.GET("/", welcome)
		userRoutes.POST("/user", controllers.CreateUser)
		userRoutes.GET("/users", controllers.GetAllUsers)
		userRoutes.GET("/user/:userId", controllers.GetSingleUser)
	}
	return r
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

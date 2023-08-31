package routes

import (
	"app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(address string) error {
	r := gin.Default()
	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.GET("/", welcome)
		userRoutes.POST("/user", controllers.CreateUser)
		userRoutes.GET("/users", controllers.GetAllUsers)
		userRoutes.GET("/user/:userId", controllers.GetSingleUser)
	}
	return r.Run(address)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

package routes

import (
	"app/component/app_context"
	"app/memcache"
	"app/middleware"
	"app/modules/user/repository/sql"
	"app/modules/user/transport/ginuser"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx app_context.AppContext) *gin.Engine {
	r := gin.Default()

	userStore := sql.NewSQLRepo(appCtx.GetMainDBConnection())
	userCaching := memcache.NewUserCaching(memcache.NewCaching(), userStore)
	userRoutes := r.Group("/api")
	{
		userRoutes.GET("/", welcome)
		userRoutes.POST("/register", ginuser.HandleRegister(appCtx))
		userRoutes.POST("/login", ginuser.HandleLogin(appCtx))
		userRoutes.POST("/import", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleImportUserCsv(appCtx))
		userRoutes.GET("/users", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleListUser(appCtx))
		userRoutes.GET("/user/:userId", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleFindUser(appCtx))
		userRoutes.DELETE("/user/:userId", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleDeleteUser(appCtx))
	}
	return r
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

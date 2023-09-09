package routes

import (
	"app/component/app_context"
	"app/memcache"
	"app/middleware"
	"app/modules/news/transport/ginnews"
	"app/modules/user/repository/sql"
	"app/modules/user/transport/ginuser"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx app_context.AppContext) *gin.Engine {
	r := gin.Default()

	userStore := sql.NewSQLRepo(appCtx.GetMainDBConnection())
	userCaching := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/", welcome)
		apiRoutes.POST("/register", ginuser.HandleRegister(appCtx))
		apiRoutes.POST("/login", ginuser.HandleLogin(appCtx))
		apiRoutes.POST("/import", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleImportUserCsv(appCtx))
		apiRoutes.GET("/users", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleListUser(appCtx))
		apiRoutes.GET("/user/:userId", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleFindUser(appCtx))
		apiRoutes.PATCH("/user/update", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleUpdateaUser(appCtx))
		apiRoutes.DELETE("/user/:userId", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleDeleteUser(appCtx))
	}

	newsRoutes := apiRoutes.Group("/news", middleware.RequireAuth(appCtx, userCaching))
	{
		newsRoutes.POST("", ginnews.HandelCreateNews(appCtx))
		newsRoutes.GET("", ginnews.HanldListNews(appCtx))
		newsRoutes.GET("/:newsId", ginnews.HanldFindNews(appCtx))
		newsRoutes.PATCH("/:newsId", ginnews.HandleUpdateNews(appCtx))
		newsRoutes.DELETE("/:newsId", ginnews.HandelDeleteNews(appCtx))
	}
	return r
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

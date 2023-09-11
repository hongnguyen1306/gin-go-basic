package routes

import (
	"app/component/app_context"
	"app/memcache"
	"app/middleware"
	"app/modules/news/transport/ginnews"
	"app/modules/news_like/transport/ginnewslike"
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

		userRoutes := apiRoutes.Group("/users")
		{
			userRoutes.POST("/register", ginuser.HandleRegister(appCtx))
			userRoutes.POST("/login", ginuser.HandleLogin(appCtx))
			userRoutes.POST("/import", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleImportUserCsv(appCtx))
			userRoutes.GET("", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleListUser(appCtx))
			userRoutes.GET("/:id", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleFindUser(appCtx))
			userRoutes.DELETE("/:id", middleware.RequireAuth(appCtx, userCaching), ginuser.HandleDeleteUser(appCtx))
		}

		newsRoutes := apiRoutes.Group("/news", middleware.RequireAuth(appCtx, userCaching))
		{
			newsRoutes.POST("", ginnews.HandelCreateNews(appCtx))
			newsRoutes.GET("", ginnews.HanldListNews(appCtx))
			newsRoutes.GET("/:id", ginnews.HanldFindNews(appCtx))
			newsRoutes.PATCH("/:id", ginnews.HandleUpdateNews(appCtx))
			newsRoutes.DELETE("/:id", ginnews.HandelDeleteNews(appCtx))
		}

		newsLikeRoutes := apiRoutes.Group("/news/:id/like", middleware.RequireAuth(appCtx, userCaching))
		{
			newsLikeRoutes.GET("/users", ginnewslike.HandleListUserLikeNews(appCtx))
			newsLikeRoutes.POST("", ginnewslike.HandleUserLikeNews(appCtx))
			newsLikeRoutes.DELETE("", ginnewslike.HandleUserUnlikeNews(appCtx))
		}
	}

	return r
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

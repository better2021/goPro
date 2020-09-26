package route

import (
	"ginGo/controller"
	"ginGo/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"time":time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.Use(middleware.CorsMiddleware(),middleware.RecoveryMiddleware()) // 使用跨域中间件 和 cover()中间件
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.POST("api/auth/info",middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("",categoryController.Create)
	categoryRoutes.PUT("/:id",categoryController.Update)
	categoryRoutes.GET("",categoryController.Search)
	categoryRoutes.DELETE("/:id",categoryController.Delete)

	postRoutes := r.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("",postController.Create)
	postRoutes.PUT("/:id",postController.Update)
	postRoutes.GET("",postController.Search)
	postRoutes.DELETE("/:id",postController.Delete)

	postRoutes.GET("/page/list",postController.PageList)

	return r
}

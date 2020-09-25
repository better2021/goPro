package route

import (
	"ginGo/controller"
	"ginGo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CorsMiddleware()) // 使用跨域中间件
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.POST("api/auth/info",middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("",categoryController.Create)
	categoryRoutes.PUT("/:id",categoryController.Update)
	categoryRoutes.GET("",categoryController.Search)
	categoryRoutes.DELETE("/:id",categoryController.Delete)

	return r
}

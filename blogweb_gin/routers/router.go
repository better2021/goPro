package routers

import (
	"blogweb_gin/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//设置session midddleware
	store := cookie.NewStore([]byte("loginuser"))

	//router.Use(controllers.SessionMiddleware())
	router.Use(sessions.Sessions("mysession", store))
	{

		//注册：
		router.GET("/register", controllers.RegisterGet)
		router.POST("register",controllers.RegisterPost)
	}

	return router

}

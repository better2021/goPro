package ginStudy

import (
	"basicGo/basic/gin/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Lession01(){
	r := gin.Default()
	// 在路由函数中使用中间件
	r.GET("/",middleware.Mymid() ,func(c *gin.Context) {
		username := c.Query("username")
		fmt.Println(username)

		c.JSON(http.StatusOK, gin.H {
			"msg":"hello "+username,
		})
	})
	r.Run(":80")
}


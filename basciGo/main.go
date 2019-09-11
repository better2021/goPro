package main

import (
	"basicGo/control"
	"github.com/gin-gonic/gin"
)


func main()  {
	router:= gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200,"hello go")
	})

	api:=router.Group("/api")
	{
		api.GET("/user",control.UserList)
		api.POST("/user",control.UserCreat)
		api.PUT("/user/:id",control.UserUpdate)
		api.DELETE("/user/:id",control.UserDelete)
	}

	router.Run(":80")
}

package main

import (
	"basicGo/control"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello go")
	})

	api := router.Group("/api")
	{
		api.GET("/user", control.UserList)
		api.POST("/user", control.UserCreat)
		api.PUT("/user/:id", control.UserUpdate)
		api.DELETE("/user/:id", control.UserDelete)

		api.GET("/product", control.ProductList)
		api.POST("/product", control.ProductCreat)
		api.PUT("/product/:id", control.ProductUpdate)
		api.DELETE("/product/:id", control.ProductDelete)

		api.GET("/menu", control.ClassList)
		api.POST("/menu", control.ClasstCreat)
		api.PUT("/menu/:id", control.ClasstUpdate)
		api.DELETE("/menu/:id", control.ClasstDelete)
	}

	router.Run(":80")
}

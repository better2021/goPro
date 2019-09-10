package main

import "github.com/gin-gonic/gin"


func main()  {
	router:= gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200,"hello go")
	})
	router.Run(":80")
}

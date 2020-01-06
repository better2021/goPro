package ginStudy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Lession01(){
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		fmt.Println(username)

		c.JSON(http.StatusOK, gin.H {
			"msg":"hello "+username,
		})
	})
	r.Run(":80")
}
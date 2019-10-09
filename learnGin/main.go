package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Login struct {
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c.Request.URL,"--")
		agent := c.GetHeader("User-Agent")
		c.JSON(200,gin.H{
			"message":"pong",
			"User-Agent":agent,
		})
	})

	// 路劲传参
	r.GET("/video/:vid", func(c *gin.Context) {
		name := c.Param("vid")
		c.JSON(http.StatusOK,gin.H{
			"message":"hello " + name,
		})
	})

	// Query传参
	 r.GET("/welcome", func(c *gin.Context) {
		 firstname := c.DefaultQuery("firstname","Guest")
		 lastname := c.Query("lastname")
		 c.String(http.StatusOK,"Hello %s %s",firstname,lastname)
	 })

	// form表单传值
	r.POST("/form_post", func(c *gin.Context) {
		message  :=  strings.TrimSpace(c.PostForm("message"))
		nick := c.DefaultPostForm("nick","feiyu")

		if message==""{
			c.String(200,"请输入message")
			return
		}

		c.JSON(200,gin.H{
			"status":http.StatusOK,
			"message":message,
			"nick":nick,
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBind(&json);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		fmt.Println(c.GetHeader("Token"),"--")

		if json.User!="admin" || json.Password != "123"{
			c.JSON(http.StatusUnauthorized,gin.H{
				"message":"账号或密码错误,StatusUnauthorized",
			})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"message":"登录成功！",
			})
		}
	})

	r.Run(":8081")
}

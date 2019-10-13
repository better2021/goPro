package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
			"path":c.Request.URL.Path,
			"method":c.Request.Method,
		})
	})

	// 结构体返回
	r.GET("/getJson/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id == 1 {
			c.JSON(200,map[string]interface{}{
				"code":1,
				"message":"ok",
				"isHttps":false,
			})
		}else if id == 2{
			resp := Login{User:"coco",Password:"123456"}
			c.JSON(200,&resp)
		}else {
			c.String(401,"参数只能是1或者2")
		}
	})

	// 加载静态页面
	r.GET("/static/page", func(c *gin.Context) {
		// 设置html目录
		r.LoadHTMLGlob("./static/*")
		c.HTML(http.StatusOK,"hello.html",nil)
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
		var jsons Login
		if err := c.ShouldBind(&jsons);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		
		fmt.Println(c.GetHeader("Token"),"--")
		fmt.Println(jsons,"++")

		if jsons.User!="admin" || jsons.Password != "123"{
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

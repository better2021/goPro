package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"path"
	"time"
)


type UserInfo struct {
	Username string `form:"username"`
	Password string	`form:"password"`
}

// 中间件
func mq(c *gin.Context){
	time := time.Now()
	fmt.Println(time)
	// c.Next()
}

func main() {
	r := gin.Default()

	r.Use(mq) // 全局使用中间件
	r.LoadHTMLFiles("./login.html","./index.html")

	r.GET("/",mq,func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"hello":"hello",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/web",func(c *gin.Context){
		name := c.Query("name")
		sex:= c.DefaultQuery("sex","女生")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"sex":sex,
		})
	})

	r.GET("/login",func(c *gin.Context){
		c.HTML(http.StatusOK,"login.html",nil)
	})

	//r.POST("/login",func(c *gin.Context){
	//	username := c.PostForm("username")
	//	password := c.DefaultPostForm("password","***")
	//	c.HTML(http.StatusOK,"index.html",gin.H{
	//		"Name":username,
	//		"Password":password,
	//	})
	//})

	r.POST("/user",func(c *gin.Context){
		var u UserInfo // 声明一个UserInfo类型的u变量
		err := c.ShouldBind(&u)
		fmt.Printf("%#v\n",u)
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})


	r.POST("/upload",func(c *gin.Context){
		file,err := c.FormFile("f1")
		if err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			// dst := fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./img",file.Filename)
			c.SaveUploadedFile(file,dst)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})

	// 重定向
	r.GET("/index",func(c *gin.Context){
		c.Redirect(http.StatusMovedPermanently,"https://www.baidu.com")
	})


	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}







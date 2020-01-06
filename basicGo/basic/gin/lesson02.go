package ginStudy

import (
	"basicGo/basic/gin/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func login(c *gin.Context){
	var user User

	fmt.Println(c.PostForm("username"))
	fmt.Println(c.PostForm("password"))

	err := c.ShouldBind(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"username":user.Username,
		"paddword":user.Password,
	})
}

// 文件上传
func upload(c *gin.Context){
	file,err := c.FormFile("file")
	if err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
		})
		return
	}
	fmt.Println(file.Filename)
	dist := "./gin/static/"+file.Filename // 文件上传后保存的地址（相对地址）
	errs := c.SaveUploadedFile(file,dist)
	if errs!=nil{
		fmt.Println(errs)
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
	})
}

func Lession02(){
	router := gin.Default()
	// 全局中间件
	router.Use(middleware.Mymid())

	router.POST("/login",login)
	router.POST("/upload",upload)
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	//router.Static("/static","./static")
	router.Run(":3000")
}
package ginStudy

import (
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

func Lession02(){
	router := gin.Default()
	router.POST("/login",login)
	router.Run(":3000")
}
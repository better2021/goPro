package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc{
	// Do some initialization logic here
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginuser := session.Get("loginuser")

		fmt.Println("loginuser",loginuser)
		if loginuser != nil{
			c.Set("IsLogin",true)
			c.Set("Loginuser",loginuser)
		}else {
			c.Set("IsLogin",false)
		}

		isLogin,_ := c.Get("IsLogin")
		loginuser2,_ := c.Get("Loginuser")
		fmt.Println("middleware...",isLogin,loginuser2)

		c.Next()
	}
}

// 获取session
func GetSession(c *gin.Context) bool {

	session := sessions.Default(c)
	loginuser := session.Get("loginuser")

	fmt.Println("loginuser",loginuser)

	if loginuser!=nil{
		return true
	}else {
		return false
	}
}
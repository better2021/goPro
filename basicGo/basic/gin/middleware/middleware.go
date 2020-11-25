package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 中间件需要返回gin.HandlerFunc函数，多个中间件通过Next函数来依次执行
func Mymid() gin.HandlerFunc{
	return func(c *gin.Context) {
		host := c.Request.Host
		fmt.Printf("before:%s\n",host)
		c.Next()
		fmt.Println("next...")
	}
}


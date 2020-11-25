package middleware

import (
	"fmt"
	"ginGo/response"
	"github.com/gin-gonic/gin"
)

/* panic恢复中间件 */
func RecoveryMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover();err!=nil{
				response.Fail(ctx,fmt.Sprint(err),nil)
			}
		}()

		ctx.Next()
	}
}
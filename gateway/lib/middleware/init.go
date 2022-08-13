package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接受服务
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["userService"] = service[0]
		ctx.Keys["recordService"] = service[1]
		ctx.Next()
	}
}

// Error
func HandlerErr() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", err),
				})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

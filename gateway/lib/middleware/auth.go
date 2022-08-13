package middleware

import (
	"gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// token认证中间件（权限控制）
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization") //postman测试：在Headers中添加： key：Authorization；value：beishuo:xxx(token值)

		// 无效的token
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//index:=strings.Index(tokenString,auth+":") //找到token前缀对应的位置
		//tokenString = tokenString[index+len(auth)+1:] //截取真实的token（开始位置为：索引开始的位置+关键字符的长度+1(:的长度为1)）
		//fmt.Println("截取之后的数据：",tokenString)
		token, _, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid { //解析错误或者过期等
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 将认证转移，节省代码，不在这边写，否则需要引入model和数据库的连击

		//// 验证通过后获取claim 中的userId
		//userId := claims.UserId
		////判定
		//var user model.Api_user
		//common.DB.First(&user, userId)
		//if user.ID == 0 { //如果没有读取到内容，说明token值有误
		//	ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		//	ctx.Abort()
		//	return
		//}
		//ctx.Set("user", user) //将key-value值存储到context中

		ctx.Next()
	}
}

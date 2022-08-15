package lib

import (
	"gateway/lib/handlers"
	"gateway/lib/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 多参数
func NewRoute(service ...interface{}) *gin.Engine {
	route := gin.Default()
	// 注册中间件
	route.Use(middleware.CORSMiddleware(), middleware.HandlerErr(), middleware.InitMiddleware(service))

	// session cookie
	store := cookie.NewStore([]byte("cookie"))
	route.Use(sessions.Sessions("session", store))

	//路由
	urls := route.Group("/todoList")
	{
		urls.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "PONG"})
		})
		urls.POST("/user/register", handlers.UserRegister)
		urls.POST("/user/login", handlers.UserLogin)

		// 需要登录认证的接口
		urls.POST("/formRecord", middleware.AuthMiddleware(), handlers.FormRecord)     // 增
		urls.DELETE("/record/:id", middleware.AuthMiddleware(), handlers.DeleteRecord) // 删
		urls.PUT("/record/:id", middleware.AuthMiddleware(), handlers.UpdateRecord)    // 改

		urls.GET("/recordList", middleware.AuthMiddleware(), handlers.ListRecord) // 查
		urls.GET("/recordInfo", middleware.AuthMiddleware(), handlers.RecordInfo)
		urls.GET("/recordHistory", middleware.AuthMiddleware(), handlers.OpHistory)
	}

	return route
}

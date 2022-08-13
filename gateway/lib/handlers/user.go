package handlers

import (
	"context"
	"gateway/pkg/utils"
	"gateway/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
	var req service.UserRequest
	utils.UserError(ctx.Bind(&req))

	// 从gin.Key中取出服务实例
	userService := ctx.Keys["userService"].(service.UserService)
	resp, err := userService.UserRegister(context.Background(), &req)

	utils.UserError(err)
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	var req service.UserRequest
	utils.UserError(ctx.Bind(&req))
	// 从gin.Key中取出服务实例
	userService := ctx.Keys["userService"].(service.UserService)
	resp, err := userService.UserLogin(context.Background(), &req)

	utils.UserError(err)
	// 签发token
	token, err := utils.ReleaseToken(uint(resp.UserInfo.ID))
	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  "成功",
		"data": gin.H{
			"user":  resp.UserInfo,
			"token": token,
		},
	})
}

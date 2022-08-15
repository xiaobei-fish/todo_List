package handlers

import (
	"context"
	"errors"
	"gateway/pkg/logging"
	"gateway/pkg/utils"
	"gateway/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 生成一条记录
func FormRecord(ctx *gin.Context) {
	var req service.RecordRequest
	utils.RecordError(ctx.Bind(&req))

	// 从gin.Key中取出服务实例
	recordService := ctx.Keys["recordService"].(service.RecordService)
	// token认证
	token, claim, err := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil || !token.Valid { //解析错误或者过期等
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	// id绑定
	req.Uid = uint64(claim.Id)
	resp, err := recordService.FormRecord(context.Background(), &req)

	utils.RecordError(err)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "new a record successfully", "data": resp.RecordInfo})
}

// 删除一条记录
func DeleteRecord(ctx *gin.Context) {
	var req service.RecordRequest
	utils.RecordError(ctx.Bind(&req))

	// 从gin.Key中取出服务实例
	recordService := ctx.Keys["recordService"].(service.RecordService)
	// token认证
	token, claim, err := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil || !token.Valid { //解析错误或者过期等
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	// id绑定
	recordId, _ := strconv.Atoi(ctx.Param("id"))
	req.Id = uint64(recordId)
	req.Uid = uint64(claim.Id)
	resp, err := recordService.DeleteRecord(context.Background(), &req)

	utils.RecordError(err)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "delete a record successfully", "data": resp.RecordInfo})
}

// 更新一条记录
func UpdateRecord(ctx *gin.Context) {
	var req service.RecordRequest
	utils.RecordError(ctx.Bind(&req))

	// 从gin.Key中取出服务实例
	recordService := ctx.Keys["recordService"].(service.RecordService)
	// token认证
	token, claim, err := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil || !token.Valid { //解析错误或者过期等
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	// id绑定
	recordId, _ := strconv.Atoi(ctx.Param("id"))
	req.Id = uint64(recordId)
	req.Uid = uint64(claim.Id)
	resp, err := recordService.UpdateRecord(context.Background(), &req)

	utils.RecordError(err)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "update a record successfully", "data": resp.RecordInfo})
}

// 返回一条记录
func RecordInfo(ctx *gin.Context) {
	var req service.RecordRequest
	utils.RecordError(ctx.Bind(&req))

	// 从gin.Key中取出服务实例
	recordService := ctx.Keys["recordService"].(service.RecordService)
	// token认证
	token, claim, err := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil || !token.Valid { //解析错误或者过期等
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	// id绑定
	recordId, _ := strconv.Atoi(ctx.Param("id"))
	req.Id = uint64(recordId)
	req.Uid = uint64(claim.Id)
	resp, err := recordService.GetRecord(context.Background(), &req)

	utils.RecordError(err)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "get a record successfully", "data": resp.RecordInfo})
}

// 返回记录列表
func ListRecord(ctx *gin.Context) {
	var req service.RecordRequest
	utils.RecordError(ctx.Bind(&req))

	// 从gin.Key中取出服务实例
	recordService := ctx.Keys["recordService"].(service.RecordService)
	// token认证
	token, claim, err := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil || !token.Valid { //解析错误或者过期等
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	// id绑定
	req.Uid = uint64(claim.Id)
	resp, err := recordService.GetRecordsList(context.Background(), &req)

	utils.RecordError(err)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "get recordList successfully", "data": resp.RecordList, "cnt:": resp.Count})
}

// 返回历史记录
func OpHistory(ctx *gin.Context) {
	var req service.HistoryRequest
	err := ctx.Bind(&req)
	if err != nil {
		err = errors.New("historyService | err: " + err.Error())
		logging.Info(err)
		panic(err)
	}
	// 从gin.Key中取出服务实例
	recordService := ctx.Keys["recordService"].(service.RecordService)
	// token认证
	token, claim, err := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err != nil || !token.Valid { //解析错误或者过期等
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	// 绑定
	req.Uid = uint64(claim.Id)
	resp, err := recordService.OpHistory(context.Background(), &req)

	if err != nil {
		err = errors.New("historyService | err: " + err.Error())
		logging.Info(err)
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "get history successfully", "data": resp.History})
}

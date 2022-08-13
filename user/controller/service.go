package controller

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"
	"user/model"
	"user/service"
)

// 用户注册服务
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserInfoResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	count := 0 // 记录计数
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	// 插入数据
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserInfo = SetUser(user) // 数据绑定
	return nil
}

// 用户登录服务
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, resp *service.UserInfoResponse) error {
	var user model.User
	resp.Code = http.StatusOK
	// 查询用户
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		// 查询不到用户记录
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = http.StatusBadRequest
			return nil
		}
		// 其他报错
		resp.Code = http.StatusInternalServerError
		return nil
	}
	// 密码错误
	if user.CheckPassword(req.Password) == false {
		resp.Code = http.StatusBadRequest
		return nil
	}
	resp.UserInfo = SetUser(user) // 数据绑定
	return nil
}

// 用户绑定到服务的userModel
func SetUser(o_user model.User) *service.UserModel {
	userModel := service.UserModel{
		ID:        uint32(o_user.ID),
		UserName:  o_user.UserName,
		CreatedAt: o_user.CreatedAt.Unix(),
		UpdatedAt: o_user.UpdatedAt.Unix(),
	}
	return &userModel
}

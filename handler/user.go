package handler

import (
	"context"
	"github.com/golineshop/user/domain/model"
	"github.com/golineshop/user/domain/service"
	user "github.com/golineshop/user/proto/user"
)

type UserController struct {
	UserService service.UserService
}

// Register 注册
func (u *UserController) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest, userRegisterResponse *user.UserRegisterResponse) error {
	userModel := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserService.AddUser(userModel)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}

// Login 登陆
func (u *UserController) Login(ctx context.Context, userLoginRequest *user.UserLoginRequest, userLoginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}
	userLoginResponse.IsSuccess = isOk
	return nil
}

// GetUserInfo 获取用户信息
func (u *UserController) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest, userInfoResponse *user.UserInfoResponse) error {
	userModel, err := u.UserService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = userToResponse(userModel)
	return nil
}

// UserForResponse 类型转化
func userToResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}

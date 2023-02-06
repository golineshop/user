package handler

import (
	"context"
	"github.com/golineshop/user/domain/model"
	"github.com/golineshop/user/domain/service"
	proto "github.com/golineshop/user/proto"
)

type UserController struct {
	UserService service.IUserService
}

// Register 注册
func (u *UserController) Register(ctx context.Context, userRegisterRequest *proto.UserRegisterRequest, userRegisterResponse *proto.UserRegisterResponse) error {
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
func (u *UserController) Login(ctx context.Context, userLoginRequest *proto.UserLoginRequest, userLoginResponse *proto.UserLoginResponse) error {
	isOk, err := u.UserService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}
	userLoginResponse.IsSuccess = isOk
	return nil
}

// GetUserInfo 获取用户信息
func (u *UserController) GetUserInfo(ctx context.Context, userInfoRequest *proto.UserInfoRequest, userInfoResponse *proto.UserInfoResponse) error {
	userModel, err := u.UserService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = userToResponse(userModel)
	return nil
}

// UserForResponse 类型转化
func userToResponse(userModel *model.User) *proto.UserInfoResponse {
	response := &proto.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}

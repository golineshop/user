package service

import (
	"errors"
	"github.com/golineshop/user/domain/model"
	"github.com/golineshop/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	AddUser(user *model.User) (userID int64, err error)
	DeleteUser(userID int64) (err error)
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(name string) (user *model.User, err error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

// NewUserService 创建实例
func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{UserRepository: userRepository}
}

type UserService struct {
	UserRepository repository.IUserRepository
}

// GeneratePassword 加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword 验证用户密码
func ValidatePassword(userPassword string, hashedPassword string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

// AddUser 添加用户
func (u *UserService) AddUser(user *model.User) (userID int64, err error) {
	passwordByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(passwordByte)
	return u.UserRepository.CreateUser(user)
}

// DeleteUser 删除用户
func (u *UserService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

// UpdateUser 更新用户
func (u *UserService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	// 检测是否更新了密码
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

// FindUserByName 根据用户名称查找用信息
func (u *UserService) FindUserByName(name string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(name)
}

// CheckPwd 比对账号密码是否正确
func (u *UserService) CheckPwd(name string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(name)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}

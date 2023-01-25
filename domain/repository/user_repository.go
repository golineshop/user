package repository

import (
	"github.com/golineshop/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	// InitTable 初始化数据表
	InitTable() (err error)
	// FindUserByName 根据用户名称查找用户信息
	FindUserByName(name string) (user *model.User, err error)
	// FindUserByID 根据用户ID查找用户信息
	FindUserByID(userID int64) (user *model.User, err error)
	// CreateUser 创建用户
	CreateUser(user *model.User) (userID int64, err error)
	// DeleteUserByID 根据用户ID删除用户
	DeleteUserByID(userID int64) (err error)
	// UpdateUser 更新用户信息
	UpdateUser(user *model.User) (err error)
	// FindAll 查找所有用
	FindAll() (userAll []model.User, err error)
}

// NewUserRepository 创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 创建表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

// FindUserByName 根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

// FindUserByID 根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_id = ?", userID).Find(user).Error
}

// CreateUser 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

// DeleteUserByID 根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

// UpdateUser 更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&model.User{}).Error
}

// FindAll 查找所有用户
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}

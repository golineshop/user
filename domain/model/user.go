package model

type User struct {
	//主键
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	//用户名称
	UserName string `gorm:"unique_index;not_null"`
	//姓
	FirstName string
	//密码哈希值
	HashPassword string
}

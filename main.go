package main

import (
	"fmt"
	"github.com/golineshop/user/domain/repository"
	"github.com/golineshop/user/domain/service"
	"github.com/golineshop/user/handler"
	"github.com/golineshop/user/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
)

func main() {
	// 服务参数设置
	microService := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	// 初始化服务
	microService.Init()
	// 创建数据库连接
	db, err := gorm.Open("mysql", "root:bowen0216@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	// 用完需要关闭数据库连接
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			println(err)
		}
	}(db)
	// 使用单数表名
	db.SingularTable(true)
	// 只执行一次，数据表初始化
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()
	// 创建服务实例
	userService := service.NewUserService(repository.NewUserRepository(db))
	// 注册Handler
	err = proto.RegisterUserHandler(microService.Server(), &handler.UserController{UserService: userService})
	if err != nil {
		fmt.Println(err)
	}
	// 运行service
	if err := microService.Run(); err != nil {
		fmt.Println(err)
	}
}

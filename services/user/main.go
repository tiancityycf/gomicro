package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"gomicro/basic"
	proto "gomicro/services/user/proto"
	"time"
)

type User struct{}

func (u *User) GetProfileById(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	if req.Id == 0 {
		rsp.Error = &proto.Error{Code: 404, Message: "user not exits"}
		return nil
	}
	profile := proto.User{Id: req.Id, Name: "hello" + req.Name}
	rsp.User = &profile
	return nil
}

func main() {

	basic.Init()
	// 创建新的服务，这里可以传入其它选项。
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("1.0"),
		micro.WrapHandler(logWrapper),
	)

	// 初始化方法会解析命令行标识
	service.Init()

	// 注册处理器
	proto.RegisterUserHandler(service.Server(), new(User))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// 实现server.HandlerWrapper接口
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%v] server request: %s %s \n", time.Now(), req.Endpoint(), req.Body())
		return fn(ctx, req, rsp)
	}
}

//Go Micro包含了函数式编程模型。
//Function是指接收一次请求，执行后便退出的服务
func newFunction() {
	// 创建新函数
	fnc := micro.NewFunction(
		micro.Name("user2"),
	)

	// 初始化命令行
	fnc.Init()

	// 注册handler
	fnc.Handle(new(User))

	// 运行服务
	fnc.Run()
}

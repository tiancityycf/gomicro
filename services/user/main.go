package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "gomicro/services/user/proto"
)

type User struct {
}

func (u *User) Info(ctx context.Context, req *proto.UserRequest, rsp *proto.UserResponse) error {
	rsp.Code = "hello" + req.Name
	return nil
}

func main() {
	// 创建新的服务，这里可以传入其它选项。
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("1.0"),
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

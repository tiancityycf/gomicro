package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "gomicro/services/user/proto"
)

func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(micro.Name("user"))
	service.Init()

	// 创建新的客户端
	user := proto.NewUserService("user", service.Client())

	// 调用user.Info
	rsp, err := user.Info(context.TODO(), &proto.UserRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 打印响应请求
	fmt.Println(rsp.Code)
}
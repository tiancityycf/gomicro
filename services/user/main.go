package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"gomicro/basic"
	"gomicro/basic/conf"
	"gomicro/services/user/handler"
	pb "gomicro/services/user/proto"
	"time"
)

func main() {

	basic.Init()
	// 使用consul注册

	micReg := consul.NewRegistry(registryOptions)

	// 创建新的服务，这里可以传入其它选项。
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("1.0"),
		micro.WrapHandler(logWrapper),
		micro.Registry(micReg),
	)

	// 初始化方法会解析命令行标识
	service.Init()

	// 注册处理器
	pb.RegisterUserHandler(service.Server(), new(handler.Service))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// 实现server.HandlerWrapper接口
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%v] server request: %s %v \n", time.Now(), req.Endpoint(), req.Body())
		return fn(ctx, req, rsp)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := conf.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
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
	fnc.Handle(new(handler.Service))

	// 运行服务
	fnc.Run()
}

#####   consul安装

去官网下载：https://www.consul.io/downloads.html

解压，设置环境变量

启动：


    consul agent -dev


启动成功后可以通过 http://localhost:8500 查看界面，相关服务发现的界面。

#####   设置GOLANG环境变量

* GO111MODULE 最好设置到环境变量中


Powershell下:

* Enable the go modules feature


    $env:GO111MODULE="on"
       
* Set the GOPROXY environment variable


    $env:GOPROXY="https://goproxy.io"


#####   安装Micro


    go get -u github.com/micro/micro
    
    
    windows 下 micro web 界面启动:
    进入GOPATH目录下的bin目录，执行
    micro.exe --server=grpc --client=grpc --transport=grpc web
    查看服务
    micro.exe list services
    获取某个服务
    micro.exe get service go.micro.web

    linux下启动：
    ./micro --server=grpc --client=grpc --transport=grpc web

    
启动成功后可以通过 http://localhost:8082 查看界面，相关服务发现的界面。

#####   Protobuf安装

Protobuf功能是代码生成，可以免去手写一些模板化的代码。

1. 下载 https://github.com/protocolbuffers/protobuf/releases 相应版本

2. 解压后，复制 protoc 至 /usr/local/bin/下

3. 尝试 protoc --version 是否成功


    安装golang的protobuf代码生成器 protoc-gen-go
    go get -u github.com/golang/protobuf/protoc-gen-go
    安装micro的protobuf插件 protoc-gen-micro
    go get -u github.com/micro/protoc-gen-micro


#####  编写服务


编写 user.proto

```
syntax = "proto3";

service User {
	rpc info(UserRequest) returns (UserResponse) {}
}

message UserRequest {
	string name = 1;
}

message UserResponse {
	string code = 1;
}

```

进入该文件所在目录生成代码

protoc --proto_path=. --micro_out=. --go_out=. ./services/user/user.proto


#####   更多参考

https://github.com/micro-in-cn/tutorials


https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro

https://micro.mu/docs/cn/features.html
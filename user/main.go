package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/conf"
	"user/controller"
	"user/service"
)

func main() {
	conf.Init()
	// 注册一个ETCD
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// 注册一个微服务事例
	microS := micro.NewService(micro.Name("userService"), micro.Address("127.0.0.1:9090"), micro.Registry(reg))

	microS.Init() // 初始化

	// 服务注册
	service.RegisterUserServiceHandler(microS.Server(), new(controller.UserService))

	// 服务启动
	microS.Run()
}

package main

import (
	"gateway/lib"
	"gateway/service"
	"gateway/wrapper"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	// 注册一个ETCD
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// 注册一个微服务事例
	umicroS := micro.NewService(micro.Name("userService.client"), micro.WrapClient(wrapper.NewUserWrapper))
	rmicroS := micro.NewService(micro.Name("recordService.client"), micro.WrapClient(wrapper.NewRecordWrapper))
	// 用户服务调用实例
	userS := service.NewUserService("userService", umicroS.Client())
	// 备忘录服务调用实例
	recordS := service.NewRecordService("recordService", rmicroS.Client())
	// 微服务事例，gin暴露端口注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		//gin暴露
		web.Handler(lib.NewRoute(userS, recordS)),
		web.Registry(reg),
		web.RegisterTTL(time.Second*30), //30秒超时
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	server.Init()
	server.Run()
}

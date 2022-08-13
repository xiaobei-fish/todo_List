package wrapper

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type userWrapper struct {
	client.Client
}

func (uw *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	name := req.Service() + "-" + req.Endpoint()
	// 熔断设置
	config := hystrix.CommandConfig{
		Timeout:                3000 * 10, // 超时时长
		RequestVolumeThreshold: 20,        // 熔断阈值，超出开始计算错误百分比
		ErrorPercentThreshold:  50,        // 错误百分比，超出则服务降级
		SleepWindow:            500 * 10,  // 熔断重启检测
	}
	hystrix.ConfigureCommand(name, config)
	// 同步方式
	return hystrix.Do(name, func() error {
		return uw.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

func NewUserWrapper(c client.Client) client.Client {
	return &userWrapper{c}
}

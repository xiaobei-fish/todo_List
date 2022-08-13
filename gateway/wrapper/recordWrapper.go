package wrapper

import (
	"context"
	"gateway/service"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)

func NewRecord(id uint64, name string) *service.RecordModel {
	return &service.RecordModel{
		Id:         id,
		Title:      name,
		Content:    "响应超时",
		StartTime:  1000,
		EndTime:    1000,
		Status:     0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

// 降级函数
func DefaultRecord(resp interface{}) {
	recordModels := make([]*service.RecordModel, 0)
	for i := 0; i < 10; i++ {
		recordModels = append(recordModels, NewRecord(uint64(i), "降级备忘录服务"+strconv.Itoa(20+i)))
	}
	res := resp.(*service.RecordListResponse)
	res.RecordList = recordModels
}

type recordWrapper struct {
	client.Client
}

func (uw *recordWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
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
		return uw.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		return err
	})
}

func NewRecordWrapper(c client.Client) client.Client {
	return &recordWrapper{c}
}

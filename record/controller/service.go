package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"record/model"
	"record/service"
)

// 备忘录操作

// 生成一条记录
func (*RecordService) FormRecord(ctx context.Context, req *service.RecordRequest, resp *service.RecordInfoResponse) error {
	mq, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ.Channel() err | err: " + err.Error())
		return err
	}
	// 开启持久化，非排他队列（非独占），非自动断开，等待服务器确认阻塞
	q, _ := mq.QueueDeclare("record", true, false, false, false, nil)
	body, _ := json.Marshal(req) // 序列化结构体便于操作
	// 使用默认交换机，直连
	err = mq.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,    // 持久化
		ContentType:  "application/json", // 消息类型
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ.Publish() err | err: " + err.Error())
		return err
	}
	return nil
}

// 取一条记录
func (*RecordService) GetRecord(ctx context.Context, req *service.RecordRequest, resp *service.RecordInfoResponse) error {
	record := model.Record{}
	model.DB.First(&record, req.Id) // 返回一条id的记录

	res := SetRecord(record) // 绑定到响应信息内
	resp.RecordInfo = res
	return nil
}

// 取一组记录
func (*RecordService) GetRecordsList(ctx context.Context, req *service.RecordRequest, resp *service.RecordListResponse) error {
	if req.Limit == 0 {
		req.Limit = 5 // 如果未设置取多少条，设为默认取5条数据
	}
	var recordList []model.Record
	count := 0
	// 从数据库中取数据
	// 计数，返回用户总备忘录记录条数
	model.DB.Model(&model.Record{}).Where("uid=?", req.Uid).Count(&count)
	resp.Count = uint32(count)
	err := model.DB.Offset(req.Start).Limit(req.Limit).Where("uid=?", req.Uid).Find(&recordList).Error
	if err != nil {
		err = errors.New("MySql select err | err: " + err.Error())
		return err
	}
	// 将结果封装到切片中
	var res []*service.RecordModel
	for _, v := range recordList {
		res = append(res, SetRecord(v))
	}
	resp.RecordList = res
	return nil
}

// 更新一条记录
func (*RecordService) UpdateRecord(ctx context.Context, req *service.RecordRequest, resp *service.RecordInfoResponse) error {
	record := model.Record{}
	err := model.DB.Model(&model.Record{}).Where("id=? and uid=?", req.Id, req.Uid).Find(&record).Error // 查找出对应的数据
	if err != nil {
		err = errors.New("MySql update err | err: " + err.Error())
		return err
	}
	// 更新数据
	record.Title = req.Title
	record.Content = req.Content
	record.Status = int(req.Status)
	// 更新保存
	model.DB.Save(&record)
	resp.RecordInfo = SetRecord(record)

	return nil
}

// 删除一条记录
func (*RecordService) DeleteRecord(ctx context.Context, req *service.RecordRequest, resp *service.RecordInfoResponse) error {
	err := model.DB.Model(&model.Record{}).Where("id=? and uid=?", req.Id, req.Uid).Delete(&model.Record{}).Error
	if err != nil {
		err = errors.New("MySql delete err | err: " + err.Error())
		return err
	}
	return nil
}

// 绑定记录
func SetRecord(record model.Record) *service.RecordModel {
	RecordModel := service.RecordModel{
		Id:         uint64(record.ID),
		Uid:        uint64(record.Uid),
		Title:      record.Title,
		Content:    record.Content,
		StartTime:  record.StartTime,
		EndTime:    record.EndTime,
		Status:     int64(record.Status),
		CreateTime: record.CreatedAt.Unix(),
		UpdateTime: record.UpdatedAt.Unix(),
	}
	return &RecordModel
}

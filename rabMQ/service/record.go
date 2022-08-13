package service

import (
	"encoding/json"
	"errors"
	"log"
	"rabMQ/model"
)

// 从rabbitMQ读消息写库
func SaveRecord() {
	mq, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ.Channel() err | err: " + err.Error())
		panic(err)
	}
	// 开启持久化，非排他队列（非独占），非自动断开，等待服务器确认阻塞
	q, _ := mq.QueueDeclare("record", true, false, false, false, nil)
	// 消费者默认设置，每次拉取一条消息数据
	err = mq.Qos(1, 0, false)
	// 取消自动确认提交，等待服务器确认阻塞，取消独占，可消费同一链接消息
	msgs, err := mq.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		err = errors.New("rabbitMQ.Consume() err | err: " + err.Error())
		panic(err)
	}
	// 监听状态，需要阻塞主进程
	go func() {
		for msg := range msgs {
			var record model.Record
			err := json.Unmarshal(msg.Body, &record)
			if err != nil {
				err = errors.New("json.Unmarshal() err | err: " + err.Error())
				panic(err)
			}
			model.DB.Create(&record)
			log.Println("Done")
			msg.Ack(false)
		}
	}()
}

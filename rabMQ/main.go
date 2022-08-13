package main

import (
	"rabMQ/conf"
	"rabMQ/service"
)

func main() {
	conf.Init()

	msgchannel := make(chan bool)

	service.SaveRecord()

	<-msgchannel
}

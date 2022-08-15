# todo_List

### 简介：go-micro + MQ + gin + redis 的一个备忘录

### 1、启动

打开etcd开启，redis服务，rabbitMQ服务

之后启动各个文件夹下的main文件即可

### 2、功能

在etcd服务发现，通过rabbitMQ进行备忘录记录的发送和消费，通过reids储存最近处理的20次操作

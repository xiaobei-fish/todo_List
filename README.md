# todo_List

### 简介：go-micro + MQ + gin + redis 的一个备忘录

### 1、启动

打开etcd开启，redis服务，rabbitMQ服务

之后启动各个文件夹下的main文件即可

### 2、功能

在etcd服务发现，通过rabbitMQ进行备忘录记录的发送和消费，通过reids储存最近处理的20次操作



```shell
service:
  AppMode: debug
  HttpPort: 3000

mysql:
  database: mysql
  host: 127.0.0.1
  port: 3306
  user: 用户
  password: 密码
  dbname: todo_List

rabbitmq:
  rabbitMQ: amqp
  MQuser: guest
  MQpassword: guest
  MQhost: 127.0.0.1
  MQport: 5672

redis:
  RDhost: 127.0.0.1:6379
  RDdb: 0
  RDpassword:
```


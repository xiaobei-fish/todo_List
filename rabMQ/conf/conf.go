package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"rabMQ/model"
	"strings"
)

func Init() {
	LoadConfig()
	// MySql配置读取
	dbname := viper.GetString("mysql.dbname")
	sqlhost := viper.GetString("mysql.host")
	sqlport := viper.GetString("mysql.port")
	sqluser := viper.GetString("mysql.user")
	sqlpassword := viper.GetString("mysql.password")

	sqlpath := strings.Join([]string{sqluser, ":", sqlpassword, "@tcp(", sqlhost, ":", sqlport, ")/", dbname, "?charset=utf8&parseTime=true"}, "")
	model.Database(sqlpath)

	// rabbitMQ配置读取
	mqname := viper.GetString("rabbitmq.rabbitMQ")
	mqhost := viper.GetString("rabbitmq.MQhost")
	mqport := viper.GetString("rabbitmq.MQport")
	mquser := viper.GetString("rabbitmq.MQuser")
	mqpassword := viper.GetString("rabbitmq.MQpassword")

	mqpath := strings.Join([]string{mqname, "://", mquser, ":", mqpassword, "@", mqhost, ":", mqport, "/"}, "")
	model.RabbitMQ(mqpath)
}

func LoadConfig() {
	dir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/conf")
	fmt.Println("config path:" + dir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

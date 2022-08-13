package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"user/model"
)

func Init() {
	LoadConfig()
	dbname := viper.GetString("mysql.dbname")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	path := strings.Join([]string{user, ":", password, "@tcp(", host, ":", port, ")/", dbname, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
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

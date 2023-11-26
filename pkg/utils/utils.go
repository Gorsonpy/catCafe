package utils

import (
	"strings"

	"github.com/Gorsonpy/catCafe/config"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetMysqlDSN(c *config.Config) string {
	dsn := strings.Join([]string{c.Mysql.Username, ":", c.Mysql.Passwd, "@tcp(", c.Mysql.Url, ":", c.Mysql.Port, ")/", c.Mysql.Database, "?charset=utf8", "&parseTime=True", "&loc=Local"}, "")
	klog.Info("connect mysql dsn : ", dsn)
	return dsn
}

// func GetMQUrl() string {
// 	if config.RabbitMQ == nil {
// 		klog.Fatal("config not found")
// 	}

// 	url := strings.Join([]string{"amqp://", config.RabbitMQ.Username, ":", config.RabbitMQ.Password, "@", config.RabbitMQ.Addr, "/"}, "")

// 	return url
// }

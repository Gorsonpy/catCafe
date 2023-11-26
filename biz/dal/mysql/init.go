package mysql

import (
	"os"

	"github.com/Gorsonpy/catCafe/config"
	"github.com/Gorsonpy/catCafe/pkg/utils"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dsn string
	DB  *gorm.DB
)

func Init() {
	dataBytes, err := os.ReadFile("../config/config.yaml")
	if err != nil {
		panic(err)
	}
	config := config.Config{}
	yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		panic(err)
	}
	dsn := utils.GetMysqlDSN(&config)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,                                // 禁用默认事务
		Logger:                 logger.Default.LogMode(logger.Info), // 设置日志模式
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		panic(err)
	}
}

package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"server/config"
)

var db *gorm.DB

func init() {
	dbConfig := config.Config().Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	var err error
	db, err = gorm.Open(
		mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},

			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		panic("打开数据库连接失败")
	}

}

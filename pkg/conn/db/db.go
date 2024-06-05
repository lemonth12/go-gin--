package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"log"
	"preject/config"
	"preject/pkg/log"
)

func GormMysql(config config.MysqlConf) (error, *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&collation=utf8mb4_unicode_ci", config.User, config.PassWord, config.Host, config.Port, config.Database)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Logger.Errorf("mysql conn fail %v", err)
		return err, db
	}
	return nil, db
}

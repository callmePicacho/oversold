package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var mysqlConn *gorm.DB
var mysqlOnce sync.Once

func GetMysqlConn() *gorm.DB {
	mysqlOnce.Do(func() {
		dsn := "root:root@tcp(127.0.0.1:3306)/oversold?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("mysql连接失败")
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic("mysql获取连接失败")
		}

		sqlDB.SetMaxIdleConns(100)

		sqlDB.SetMaxOpenConns(100)

		sqlDB.SetConnMaxLifetime(time.Hour)

		err = sqlDB.Ping()
		if err != nil {
			panic("mysql ping失败")
		}

		mysqlConn = db
	})

	return mysqlConn
}

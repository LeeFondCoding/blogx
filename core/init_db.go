package core

import (
	"blogx/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open(global.Conf.DB.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成实体外键约束
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败 %s", err)
	}
	// 初始化数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.DB = db
	logrus.Infoln("数据库连接成功")
}

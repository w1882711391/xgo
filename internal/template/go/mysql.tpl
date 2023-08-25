package dao

import (
    "fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlInit() {
    dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		opt.Cfg.Mysql.Username,
		opt.Cfg.Mysql.Password,
		opt.Cfg.Mysql.Host,
		opt.Cfg.Mysql.Port,
		opt.Cfg.Mysql.DataBase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库启动失败")
	}
	DB = db
}
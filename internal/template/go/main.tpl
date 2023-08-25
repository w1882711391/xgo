package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

var listenHttpErrChan = make(chan error)

func init() {
    // 输出具体目录的某一行log
    logrus.SetFormatter(&logrus.TextFormatter{
    		DisableColors: true,
    		FullTimestamp: true,
    })
    logrus.SetReportCaller(true)
    //解析json绑定数据库信息
    opt.MustInitConfig()

	dao.MysqlInit()
	logrus.Info("mysql数据库已启动")
	dao.RedisInit()
	logrus.Info("redis数据库已启动")
	/*
	model.AutoMigrate()
	logrus.Info("数据库表创建成功")
	*/
}


func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go func() {
		app := route.RouterInit()
		listenHttpErrChan <- app.Listen(":8080")
	}()

	select {
	case err := <- listenHttpErrChan:
		logrus.Errorf("http err: %+v\n", err)
	case <- ctx.Done():
		logrus.Info("Shutting down gracefully...")
	}

}
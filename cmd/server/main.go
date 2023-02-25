package main

import (
	"fmt"
	"github.com/mangohow/cloud-ide-webserver/conf"
	"github.com/mangohow/cloud-ide-webserver/internal/dao/db"
	"github.com/mangohow/cloud-ide-webserver/internal/dao/rdis"
	"github.com/mangohow/cloud-ide-webserver/pkg/httpserver"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/routes"
	"syscall"
)

func main() {
	// 初始化配置
	if err := conf.LoadConf(); err != nil {
		panic(fmt.Errorf("load conf failed, reason:%s", err.Error()))
	}

	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		panic(fmt.Errorf("init logger error, reason:%v", err))
	}

	// 初始化数据库
	if err := db.InitMysql(); err != nil {
		panic(fmt.Errorf("init mysql failed, reason:%s", err.Error()))
	}

	// 初始化redis
	if err := rdis.InitRedis(); err != nil {
		panic(fmt.Errorf("init redis failed, reason:%s", err.Error()))
	}

	// 创建gin路由
	engine := routes.NewGinRouter()
	// 注册路由
	routes.Register(engine)

	// 创建http server
	server := httpserver.NewServer(conf.ServerConfig.Host, conf.ServerConfig.Port, engine)

	// 启动server
	httpserver.ListenAndServe(server)

	fmt.Println("pid:", syscall.Getpid())

	// 等待服务退出
	httpserver.WaitForShutdown(server, func() {
		db.CloseMysql()
		fmt.Println("close mysql connection.")
	})
}

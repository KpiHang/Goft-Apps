package main

import (
	"UseGoft/src/configuration"
	"UseGoft/src/controllers"
	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite(). //初始化脚手架
			Config(configuration.NewDBConfig(), configuration.NewServiceConfig()). // 设置依赖项
			Mount("v1", controllers.NewIndexController(),                          //挂载控制器
				controllers.NewUserController()).
		Launch() //启动 默认8080
}

// http://127.0.0.1:8080/v1/access_token?uname=xh&upwd=123

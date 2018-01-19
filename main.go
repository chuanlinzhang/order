package main

import (
	_ "order/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const APP_VER   = "0.1.1.0227"
func main() {
	//用beego.Info打印日志到文件
	beego.Info(beego.BConfig.AppName,APP_VER)
	//注册模板函数功能
	beego.AddFuncMap("i18n",i18n.Tr)
	beego.Run()
}


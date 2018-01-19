package routers

import (
	"order/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.AppController{})
    beego.Router("/join",&controllers.AppController{},"post:Join")

    beego.Router("/lp",&controllers.LongPollingController{},"get:Join")
    beego.Router("/lp/post",&controllers.LongPollingController{})
    beego.Router("/lp/fetch",&controllers.LongPollingController{},"get:Fetch")
    }

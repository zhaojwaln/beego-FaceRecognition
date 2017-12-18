package routers

import (
	"beego-FaceRecognition/src/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Info("init routers start ...")

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/welcome", &controllers.MainController{}, "*:Welcome")
	beego.Router("/leftMenu", &controllers.MainController{}, "*:LeftMenu")
	beego.Router("/header", &controllers.MainController{}, "*:Header")
	beego.Router("/loadMenu", &controllers.MainController{}, "*:LoadMenu")
	
	//自动绑定映射关系
	beego.AutoRouter(&controllers.FaceController{})

	beego.Info("init routers end.")
}

package routers

import (
	"Log/LogAdmin/controllers"
	"github.com/astaxie/beego"
	"Log/LogAdmin/controllers/logconfig"
)

func init()  {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/logconfig/index",&logconfig.LogConfigController{}, "*:Index")

}

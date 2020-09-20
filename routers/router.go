package routers

import (
	"github.com/astaxie/beego"
	"oa_system/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

package routers

import (
	"github.com/astaxie/beego"
	"oa_system/controllers"
	"oa_system/controllers/login"
)

func init() {
	// 不需要登录
	beego.Router("/", &login.LoginController{})
	beego.Router("/change_captcha/", &login.LoginController{}, "get:ChangeCaptcha")

	// 需要登录
	beego.Router("/main/index/", &controllers.HomeController{})
	beego.Router("/main/welcome/", &controllers.HomeController{}, "get:Welcome")
}

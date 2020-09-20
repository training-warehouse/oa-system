package routers

import (
	"github.com/astaxie/beego"
	"oa_system/controllers/login"
)

func init() {
	beego.Router("/", &login.LoginController{})
	beego.Router("/change_captcha/", &login.LoginController{}, "get:ChangeCaptcha")
}

package routers

import (
	"github.com/astaxie/beego"
	"oa_system/controllers"
	"oa_system/controllers/login"
	"oa_system/controllers/user"
)

func init() {
	// 不需要登录
	beego.Router("/", &login.LoginController{})
	beego.Router("/change_captcha/", &login.LoginController{}, "get:ChangeCaptcha")

	// 需要登录
	beego.Router("/main/index/", &controllers.HomeController{})
	beego.Router("/main/welcome/", &controllers.HomeController{}, "get:Welcome")
	beego.Router("/main/user/list/", &user.UserController{}, "get:List")
	beego.Router("/main/user/add/", &user.UserController{}, "get:Add")
	beego.Router("/main/user/do_add/", &user.UserController{}, "post:DoAdd")
	beego.Router("/main/user/is_active/", &user.UserController{}, "post:IsActive")
}

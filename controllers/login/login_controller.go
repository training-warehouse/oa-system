package login

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"oa_system/models/user"
	"oa_system/utils"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	id, base64, err := utils.GetCaptcha()
	if err != nil {
		return
	}
	l.Data["captcha"] = utils.Captcha{Id: id, BS64: base64}

	l.TplName = "login/login.html"
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")
	captcha := l.GetString("captcha")
	captcha_id := l.GetString("captcha_id")

	is_ok := utils.VerifyCaptcha(captcha_id, captcha)

	password_md5 := utils.GetMd5File(password)

	o := orm.NewOrm()
	is_exist := o.QueryTable("sys_user").Filter(
		"user_name", username).Filter("password", password_md5).Exist()

	ret_map := map[string]interface{}{}
	if !is_exist {
		ret_map["code"] = 10001
		ret_map["msg"] = "用户名或密码错误"
	} else if !is_ok {
		ret_map["code"] = 10001
		ret_map["msg"] = "验证码错误"
	} else {
		user_obj := user.User{}
		o.QueryTable("sys_user").Filter(
			"user_name", username).Filter("password", password_md5).One(&user_obj)
		l.SetSession("id", user_obj.Id)
		ret_map["code"] = 200
		ret_map["msg"] = "登录成功"
	}
	fmt.Println(ret_map)
	l.Data["json"] = ret_map
	l.ServeJSON()
}

func (l *LoginController) ChangeCaptcha() {
	message := map[string]interface{}{}

	id, base64, err := utils.GetCaptcha()
	if err != nil {
		message["msg"] = "生成验证码失败"
		message["Code"] = 404
		l.Data["json"] = message
	} else {
		l.Data["json"] = utils.Captcha{Id: id, BS64: base64, Code: 200}
	}
	// 以json的方式返回
	l.ServeJSON()
}

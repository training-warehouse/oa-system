package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"oa_system/models/user"
	"oa_system/utils"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) List() {
	o := orm.NewOrm()
	var users []user.User

	page_per_num := 2
	current_page, err := u.GetInt("page")
	if err != nil {
		current_page = 1
	}
	offset_num := page_per_num * (current_page - 1)
	qs := o.QueryTable("sys_user")
	qs.Filter("is_del", 0).Limit(page_per_num).Offset(offset_num).All(&users)

	u.Data["users"] = users

	pre_page := 1
	if current_page == 1 {
		pre_page = 1
	} else if current_page > 1 {
		pre_page = current_page - 1
	}
	u.Data["prePage"] = pre_page

	next_page := current_page + 1
	count, _ := qs.Filter("is_del", 0).Count()
	count_page := int(math.Ceil(float64(count) / float64(page_per_num)))
	if next_page >= count_page {
		next_page = count_page
	}
	u.Data["nextPage"] = next_page
	u.Data["currentPage"] = current_page
	u.Data["countPage"] = count_page

	page_map := utils.Paginator(current_page, page_per_num, count)
	u.Data["pageMap"] = page_map

	u.TplName = "user/user_list.html"
}

func (u *UserController) Add() {
	u.TplName = "user/user_add.html"
}

func (u *UserController) DoAdd() {
	username := u.GetString("username")
	password := u.GetString("password")
	age, _ := u.GetInt("age")
	gender, _ := u.GetInt("gender")
	phone := u.GetString("phone")
	addr := u.GetString("addr")
	is_active, _ := u.GetInt("is_active")

	new_password := utils.GetMd5File(password)
	phone_int64, _ := strconv.ParseInt(phone, 10, 64)
	o := orm.NewOrm()
	user_data := user.User{
		UserName: username,
		Password: new_password,
		Age:      age,
		Gender:   gender,
		Phone:    phone_int64,
		Addr:     addr,
		IsActive: is_active,
	}
	_, err := o.Insert(&user_data)

	message_map := map[string]interface{}{}
	if err != nil { //说明插入数据有问题

		message_map["code"] = 10001
		message_map["msg"] = "添加数据出错，请重新添加"
		u.Data["json"] = message_map
	} else {
		message_map["code"] = 200
		message_map["msg"] = "添加成功"
		u.Data["json"] = message_map
	}

	u.ServeJSON()
}

func (u *UserController) IsActive() {
	is_active, _ := u.GetInt("is_active_val")
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user").Filter("Id", id)
	fmt.Println(is_active)
	msg_map := map[string]interface{}{}
	if is_active == 1 {
		qs.Update(orm.Params{
			"is_active": 0,
		})
		msg_map["msg"] = "已停用"
	} else {
		qs.Update(orm.Params{
			"is_active": 1,
		})
		msg_map["msg"] = "已启用"
	}

	u.Data["json"] = msg_map
	u.ServeJSON()
}

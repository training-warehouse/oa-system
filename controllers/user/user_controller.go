package user

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"oa_system/models/user"
	"oa_system/utils"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
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

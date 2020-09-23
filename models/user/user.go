package user

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int       `orm:"pk;auto"`
	UserName   string    `orm:"unique;column(user_name);size(64);description(用户名)"`
	Password   string    `orm:"size(32);description(密码)"`
	Age        int       `orm:"null;description(年龄)"`
	Gender     int       `orm:"null;description(1男，2女，3保密);default(3)"`
	Phone      int64     `orm:"null;description(电话)"`
	Addr       string    `orm:"null;size(255);description(地址)"`
	IsActive   int       `orm:"description(1启用，0停用);default(1)"`
	IsDel      int       `orm:"description(1删除，0不删除);default(0)"`
	CreateTime time.Time `orm:"auto_now;type(datetime);null;description(创建时间)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User))
}

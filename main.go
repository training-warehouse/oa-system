package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "oa_system/models/user"
	_ "oa_system/routers"
)

func init() {

	// 链接数据库
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	db := beego.AppConfig.String("db")
	host := beego.AppConfig.String("host")
	post := beego.AppConfig.String("post")
	dataSource := username + ":" + password + "@tcp(" + host + ":" + post + ")/" + db + "?charset=utf8&loc=Local"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	orm.RunCommand()
	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

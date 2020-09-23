package utils

import "github.com/astaxie/beego/context"

func LoginFilter(ctx *context.Context) {
	// 获取Session
	id := ctx.Input.Session("id")
	if id == nil {
		ctx.Redirect(302, "/")
	}
}

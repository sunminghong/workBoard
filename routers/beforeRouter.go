package routers

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

func FilterUser(ctx *context.Context) {

	user := ctx.Input.Session("user")
	fmt.Println(ctx.Request.RequestURI)

	if ctx.Request.RequestURI == "/api/users/login" {
		return
	}

	if ctx.Request.RequestURI == "/api/users/create" {
		return
	}

	if user == nil {
		ctx.Redirect(302, "/login")
	}
}

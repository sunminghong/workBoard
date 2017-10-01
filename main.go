package main

import (
	"encoding/gob"
	"workBoard/models"
	_ "workBoard/models"
	_ "workBoard/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// beego bug 必须 gob.Register 才能注册 session 相关 struct
	gob.Register(models.Board{})
	gob.Register(models.Users{})
}

func main() {
	beego.SetStaticPath("/static", "frontend/build/static/")
	beego.SetViewsPath("frontend/build")

	beego.Run()
}
